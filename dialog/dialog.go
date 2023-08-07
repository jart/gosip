// Copyright 2020 Justine Alexandra Roberts Tunney
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// SIP Dialog Transport.

package dialog

import (
	"bytes"
	"errors"
	"flag"
	"log"
	"net"
	"time"

	"github.com/cresta/gosip/rtp"
	"github.com/cresta/gosip/sdp"
	"github.com/cresta/gosip/sip"
	"github.com/cresta/gosip/util"
)

const (
	Proceeding = iota
	Ringing
	Answered
	Hangup
)

var (
	looseSignalling = flag.Bool("looseSignalling", false, "Permit SIP messages from servers other than the next hop.")
	resendInterval  = flag.Int("resendInterval", 400, "Milliseconds between SIP resends.")
	maxResends      = flag.Int("maxResends", 2, "Max SIP message retransmits.")
)

// Dialog represents an outbound SIP phone call.
type Dialog struct {
	OnErr   <-chan error
	OnState <-chan int
	OnPeer  <-chan *net.UDPAddr
	Hangup  chan<- bool
}

type dialogState struct {
	sockMsgs        <-chan *sip.Msg
	sockErrs        <-chan error
	csockMsgs       <-chan *sip.Msg
	csockErrs       <-chan error
	errChan         chan<- error
	stateChan       chan<- int
	peerChan        chan<- *net.UDPAddr
	sendHangupChan  <-chan bool
	state           int              // Current state of the dialog.
	dest            string           // Destination hostname (or IP).
	addr            string           // Destination ip:port.
	sock            *net.UDPConn     // Outbound message socket (connected for ICMP)
	csock           *net.UDPConn     // Inbound socket for Contact field.
	routes          *AddressRoute    // List of SRV addresses to attempt contacting.
	invite          *sip.Msg         // Our INVITE that established the dialog.
	remote          *sip.Msg         // Message from remote UA that established dialog.
	request         *sip.Msg         // Current outbound request message.
	requestResends  int              // Number of REsends of message so far.
	requestTimer    <-chan time.Time // Resend timer for message.
	response        *sip.Msg         // Current outbound request message.
	responseResends int              // Number of REsends of message so far.
	responseTimer   <-chan time.Time // Resend timer for message.
	lseq            int              // Local CSeq value.
	rseq            int              // Remote CSeq value.
	b               bytes.Buffer     // Outbound message buffer.
}

// NewDialog creates a phone call.
func NewDialog(invite *sip.Msg) (dl *Dialog, err error) {
	errChan := make(chan error)
	stateChan := make(chan int)
	peerChan := make(chan *net.UDPAddr)
	sendHangupChan := make(chan bool, 4)
	dls := &dialogState{
		errChan:        errChan,
		stateChan:      stateChan,
		peerChan:       peerChan,
		sendHangupChan: sendHangupChan,
		invite:         invite,
	}
	go dls.run()
	return &Dialog{
		OnErr:   errChan,
		OnState: stateChan,
		OnPeer:  peerChan,
		Hangup:  sendHangupChan,
	}, nil
}

func (dls *dialogState) run() {
	defer dls.cleanup()
	if !dls.sendRequest(dls.invite) {
		return
	}
	for {
		select {
		case err := <-dls.sockErrs:
			dls.sock.Close()
			dls.sock = nil
			if util.IsRefused(err) {
				log.Printf("ICMP refusal: %s (%s)\r\n", dls.sock.RemoteAddr(), dls.dest)
				if !dls.popRoute() {
					return
				}
			} else {
				dls.errChan <- err
				return
			}
		case err := <-dls.csockErrs:
			dls.csock.Close()
			dls.csock = nil
			dls.errChan <- err
			return
		case <-dls.requestTimer:
			if !dls.resendRequest() {
				return
			}
		case <-dls.responseTimer:
			if !dls.resendResponse() {
				return
			}
		case msg := <-dls.sockMsgs:
			if !dls.handleMessage(msg) {
				return
			}
		case msg := <-dls.csockMsgs:
			if !dls.handleMessage(msg) {
				return
			}
		case <-dls.sendHangupChan:
			if !dls.sendHangup() {
				return
			}
		}
	}
}

func (dls *dialogState) sendRequest(request *sip.Msg) bool {
	host, port, err := RouteMessage(nil, nil, request)
	if err != nil {
		dls.errChan <- err
		return false
	}
	wantSRV := dls.state < Answered
	routes, err := RouteAddress(host, port, wantSRV)
	if err != nil {
		dls.errChan <- err
		return false
	}
	dls.request = request
	dls.routes = routes
	dls.dest = host
	return dls.popRoute()
}

func (dls *dialogState) popRoute() bool {
	if dls.routes == nil {
		dls.errChan <- errors.New("Failed to contact: " + dls.dest)
		return false
	}
	dls.addr = dls.routes.Address
	dls.routes = dls.routes.Next
	if !dls.connect() {
		return dls.popRoute()
	}
	dls.populate(dls.request)
	if dls.state < Answered {
		dls.rseq = 0
		dls.remote = nil
		dls.lseq = dls.request.CSeq
	}
	dls.requestResends = 0
	dls.requestTimer = time.After(duration(resendInterval))
	return dls.send(dls.request)
}

func (dls *dialogState) connect() bool {
	if dls.sock == nil || dls.addr != dls.sock.RemoteAddr().String() {
		// Create socket through which we send messages. This socket is connected
		// to the remote address so we can receive ICMP unavailable errors. It also
		// allows us to discover the appropriate IP address for the local machine.
		dls.cleanupSock()
		conn, err := net.Dial("udp", dls.addr)
		if err != nil {
			log.Printf("net.Dial(udp, %s) failed: %s\r\n", dls.addr, err)
			return false
		}
		dls.sock = conn.(*net.UDPConn)
		sockMsgs := make(chan *sip.Msg)
		sockErrs := make(chan error)
		dls.sockMsgs = sockMsgs
		dls.sockErrs = sockErrs
		go ReceiveMessages(dls.sock, sockMsgs, sockErrs)

		// But a connected UDP socket can only receive packets from a single host.
		// SIP signalling paths can change depending on the environment, so we need
		// to be able to accept packets from anyone.
		if dls.csock == nil && *looseSignalling {
			cconn, err := rtp.Listen("")
			if err != nil {
				log.Printf("Loose signalling not possible: %s\r\n", err)
				return false
			}
			dls.csock = cconn.(*net.UDPConn)
			csockMsgs := make(chan *sip.Msg)
			csockErrs := make(chan error)
			dls.csockMsgs = csockMsgs
			dls.csockErrs = csockErrs
			go ReceiveMessages(dls.csock, csockMsgs, csockErrs)
		}
	}
	return true
}

func (dls *dialogState) populate(msg *sip.Msg) {
	laddr := dls.sock.LocalAddr().(*net.UDPAddr)
	lhost := laddr.IP.String()
	lport := uint16(laddr.Port)

	if msg.Via == nil {
		msg.Via = &sip.Via{Host: lhost}
	}
	msg.Via.Port = lport
	branch := msg.Via.Param.Get("branch")
	if branch != nil {
		branch.Value = util.GenerateBranch()
	} else {
		msg.Via.Param = &sip.Param{
			Name:  "branch",
			Value: util.GenerateBranch(),
			Next:  msg.Via.Param,
		}
	}

	if msg.Contact == nil {
		msg.Contact = &sip.Addr{Uri: &sip.URI{Scheme: "sip", Host: lhost}}
	}
	if dls.csock != nil {
		msg.Contact.Uri.Port = uint16(dls.csock.LocalAddr().(*net.UDPAddr).Port)
	} else {
		msg.Contact.Uri.Port = lport
	}
	if msg.Contact.Uri.Param.Get("transport") == nil {
		msg.Contact.Uri.Param = &sip.URIParam{
			Name:  "transport",
			Value: "udp",
			Next:  msg.Contact.Uri.Param,
		}
	}

	if msg.Method == sip.MethodInvite {
		if ms, ok := msg.Payload.(*sdp.SDP); ok {
			if ms.Addr == "" {
				ms.Addr = lhost
			}
			if ms.Origin.Addr == "" {
				ms.Origin.Addr = lhost
			}
			if ms.Origin.ID == "" {
				ms.Origin.ID = util.GenerateOriginID()
			}
		}
	}
	PopulateMessage(nil, nil, msg)
}

func (dls *dialogState) handleMessage(msg *sip.Msg) bool {
	if msg.VersionMajor != 2 || msg.VersionMinor != 0 {
		if !dls.send(NewResponse(msg, sip.StatusVersionNotSupported)) {
			return false
		}
		dls.errChan <- errors.New("Remote UA is using a strange SIP version")
		return false
	}
	if msg.CallID != dls.request.CallID {
		log.Printf("Received message doesn't match dialog\r\n")
		return dls.send(NewResponse(msg, sip.StatusCallTransactionDoesNotExist))
	}
	if msg.IsResponse() {
		return dls.handleResponse(msg)
	} else {
		return dls.handleRequest(msg)
	}
}

func (dls *dialogState) handleResponse(msg *sip.Msg) bool {
	if !ResponseMatch(dls.request, msg) {
		log.Printf("Received response doesn't match transaction\r\n")
		return true
	}
	if msg.Status >= sip.StatusOK && dls.request.Method == sip.MethodInvite {
		if msg.Contact == nil {
			dls.errChan <- errors.New("Remote UA sent >=200 response w/o Contact")
			return false
		}
		if !dls.send(NewAck(msg, dls.request)) {
			return false
		}
	}
	dls.routes = nil
	dls.requestTimer = nil
	if msg.Status <= sip.StatusOK {
		dls.checkSDP(msg)
	}
	switch msg.Status {
	case sip.StatusTrying:
		dls.transition(Proceeding)
	case sip.StatusRinging, sip.StatusSessionProgress:
		dls.transition(Ringing)
	case sip.StatusOK:
		switch msg.CSeqMethod {
		case sip.MethodInvite:
			if dls.remote == nil {
				dls.transition(Answered)
			}
			dls.remote = msg
		case sip.MethodBye, sip.MethodCancel:
			dls.transition(Hangup)
			return false
		}
	case sip.StatusServiceUnavailable:
		if dls.request == dls.invite {
			log.Printf("Service unavailable: %s (%s)\r\n", dls.sock.RemoteAddr(), dls.dest)
			return dls.popRoute()
		} else {
			dls.errChan <- &sip.ResponseError{Msg: msg}
			return false
		}
	case sip.StatusMovedPermanently, sip.StatusMovedTemporarily:
		dls.invite.Request = msg.Contact.Uri
		dls.invite.Route = nil
		return dls.sendRequest(dls.invite)
	default:
		if msg.Status > sip.StatusOK {
			dls.errChan <- &sip.ResponseError{Msg: msg}
			return false
		}
	}
	return true
}

func (dls *dialogState) handleRequest(msg *sip.Msg) bool {
	if msg.MaxForwards <= 0 {
		if !dls.send(NewResponse(msg, sip.StatusTooManyHops)) {
			return false
		}
		dls.errChan <- errors.New("Remote froot loop detected")
		return false
	}
	if dls.rseq == 0 {
		dls.rseq = msg.CSeq
	} else {
		if msg.CSeq < dls.rseq {
			// RFC 3261 mandates a 500 response for out of order requests.
			return dls.send(NewResponse(msg, sip.StatusInternalServerError))
		}
		dls.rseq = msg.CSeq
	}
	switch msg.Method {
	case sip.MethodBye:
		if !dls.send(NewResponse(msg, sip.StatusOK)) {
			return false
		}
		dls.transition(Hangup)
		return false
	case sip.MethodOptions: // Probably a keep-alive ping.
		return dls.send(NewResponse(msg, sip.StatusOK))
	case sip.MethodInvite: // Re-INVITEs are used to change the RTP or signalling path.
		dls.remote = msg
		dls.checkSDP(msg)
		return dls.sendResponse(NewResponse(msg, sip.StatusOK))
	case sip.MethodAck: // Re-INVITE response has been ACK'd.
		dls.response = nil
		dls.responseTimer = nil
		return true
	default:
		return dls.send(NewResponse(msg, sip.StatusMethodNotAllowed))
	}
}

func (dls *dialogState) checkSDP(msg *sip.Msg) {
	if ms, ok := msg.Payload.(*sdp.SDP); ok {
		dls.peerChan <- &net.UDPAddr{IP: net.ParseIP(ms.Addr), Port: int(ms.Audio.Port)}
	}
}

func (dls *dialogState) send(msg *sip.Msg) bool {
	if msg.MaxForwards > 0 {
		msg.MaxForwards--
		if msg.MaxForwards == 0 {
			dls.errChan <- errors.New("Local froot loop detected")
			return false
		}
	}
	ts := time.Now()
	addTimestamp(msg, ts)
	dls.b.Reset()
	msg.Append(&dls.b)
	if *tracing {
		trace("send", dls.b.Bytes(), dls.sock.RemoteAddr())
	}
	_, err := dls.sock.Write(dls.b.Bytes())
	if err != nil {
		dls.errChan <- err
		return false
	}
	return true
}

func (dls *dialogState) resendRequest() bool {
	if dls.request == nil {
		return true
	}
	if dls.requestResends < *maxResends {
		if !dls.send(dls.request) {
			return false
		}
		dls.requestResends++
		dls.requestTimer = time.After(duration(resendInterval))
	} else {
		log.Printf("Timeout: %s (%s)\r\n", dls.sock.RemoteAddr(), dls.dest)
		if !dls.popRoute() {
			return false
		}
	}
	return true
}

// sendResponse is used to reliably send a response to an INVITE only.
func (dls *dialogState) sendResponse(msg *sip.Msg) bool {
	dls.response = msg
	dls.responseResends = 0
	dls.responseTimer = time.After(duration(resendInterval))
	return dls.send(dls.response)
}

func (dls *dialogState) resendResponse() bool {
	if dls.response == nil {
		return true
	}
	if dls.responseResends < *maxResends {
		if !dls.send(dls.response) {
			return false
		}
		dls.responseResends++
		dls.responseTimer = time.After(duration(resendInterval))
	} else {
		// TODO(jart): If resending INVITE 200 OK, start sending BYE.
		log.Printf("Timeout sending response: %s (%s)\r\n", dls.sock.RemoteAddr(), dls.dest)
		if !dls.popRoute() {
			return false
		}
	}
	return true
}

func (dls *dialogState) sendHangup() bool {
	switch dls.state {
	case Proceeding, Ringing:
		return dls.send(NewCancel(dls.invite))
	case Answered:
		return dls.sendRequest(NewBye(dls.invite, dls.remote, &dls.lseq))
	case Hangup:
		panic("Why didn't the event loop break?")
	default:
		//  o  A UA or proxy cannot send CANCEL for a transaction until it gets a
		//     provisional response for the request.  This was allowed in RFC 2543
		//     but leads to potential race conditions.
		dls.transition(Hangup)
		return false
	}
}

func (dls *dialogState) transition(state int) {
	dls.state = state
	dls.stateChan <- state
}

func (dls *dialogState) cleanup() {
	dls.cleanupSock()
	dls.cleanupCSock()
	close(dls.errChan)
	close(dls.stateChan)
	close(dls.peerChan)
}

func (dls *dialogState) cleanupSock() {
	if dls.sock != nil {
		dls.sock.Close()
		dls.sock = nil
		<-dls.sockMsgs
		<-dls.sockErrs
	}
}

func (dls *dialogState) cleanupCSock() {
	if dls.csock != nil {
		dls.csock.Close()
		dls.csock = nil
		<-dls.csockMsgs
		<-dls.csockErrs
	}
}
