// SIP Dialog Transport.

package sip

import (
	"bytes"
	"errors"
	"flag"
	"github.com/jart/gosip/sdp"
	"github.com/jart/gosip/util"
	"log"
	"net"
	"time"
)

const (
	DialogInit       = 0
	DialogProceeding = 1
	DialogRinging    = 2
	DialogAnswered   = 3
	DialogHangup     = 4
	resendInterval   = 200 * time.Millisecond
	maxResends       = 2
)

var (
	looseSignalling = flag.Bool("looseSignalling", false, "Permit SIP messages from servers other than the next hop.")
)

// Dialog represents an outbound SIP phone call.
type Dialog struct {
	OnErr   <-chan error
	OnState <-chan int
	OnSDP   <-chan *sdp.SDP
	Hangup  chan<- bool
}

type dialogState struct {
	sockMsgs        <-chan *Msg
	sockErrs        <-chan error
	csockMsgs       <-chan *Msg
	csockErrs       <-chan error
	errChan         chan<- error
	sdpChan         chan<- *sdp.SDP
	stateChan       chan<- int
	doHangupChan    <-chan bool
	state           int              // Current state of the dialog.
	dest            string           // Destination hostname (or IP).
	addr            string           // Destination ip:port.
	sock            *net.UDPConn     // Outbound message socket (connected for ICMP)
	csock           *net.UDPConn     // Inbound socket for Contact field.
	routes          *AddressRoute    // List of SRV addresses to attempt contacting.
	invite          *Msg             // Our INVITE that established the dialog.
	remote          *Msg             // Message from remote UA that established dialog.
	request         *Msg             // Current outbound request message.
	requestResends  int              // Number of REsends of message so far.
	requestTimer    <-chan time.Time // Resend timer for message.
	response        *Msg             // Current outbound request message.
	responseResends int              // Number of REsends of message so far.
	responseTimer   <-chan time.Time // Resend timer for message.
	lseq            int              // Local CSeq value.
	rseq            int              // Remote CSeq value.
}

// NewDialog creates a phone call.
func NewDialog(invite *Msg) (dl *Dialog, err error) {
	errChan := make(chan error)
	sdpChan := make(chan *sdp.SDP)
	stateChan := make(chan int)
	doHangupChan := make(chan bool, 4)
	dls := &dialogState{
		errChan:      errChan,
		sdpChan:      sdpChan,
		stateChan:    stateChan,
		doHangupChan: doHangupChan,
		invite:       invite,
	}
	go dls.run()
	return &Dialog{
		OnErr:   errChan,
		OnState: stateChan,
		OnSDP:   sdpChan,
		Hangup:  doHangupChan,
	}, nil
}

func (dls *dialogState) run() {
	defer dls.sabotage()
	defer dls.cleanup()
	if !dls.sendRequest(dls.invite, true) {
		return
	}
	for {
		select {
		case err := <-dls.sockErrs:
			if util.IsRefused(err) {
				log.Printf("ICMP refusal: %s (%s)", dls.sock.RemoteAddr(), dls.dest)
				if !dls.popRoute() {
					return
				}
			} else {
				dls.errChan <- err
				return
			}
		case err := <-dls.csockErrs:
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
		case <-dls.doHangupChan:
			if !dls.sendHangup() {
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
		}
	}
}

func (dls *dialogState) sendRequest(request *Msg, wantSRV bool) bool {
	host, port, err := RouteMessage(nil, nil, request)
	if err != nil {
		dls.errChan <- err
		return false
	}
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
	PopulateMessage(nil, nil, dls.request)
	dls.lseq = dls.request.CSeq
	dls.requestResends = 0
	dls.requestTimer = time.After(resendInterval)
	return dls.send(dls.request)
}

func (dls *dialogState) connect() bool {
	if dls.sock != nil && dls.sock.RemoteAddr().String() == dls.addr {
		return true
	}

	// Create socket through which we send messages. This socket is connected to
	// the remote address so we can receive ICMP unavailable errors. It also
	// allows us to discover the appropriate IP address for this machine.
	dls.cleanupSock()
	conn, err := net.Dial("udp", dls.addr)
	if err != nil {
		log.Printf("net.Dial(udp, %s) failed: %s", dls.addr, err)
		return false
	}
	dls.sock = conn.(*net.UDPConn)
	dls.rseq = 0
	dls.remote = nil
	laddr := conn.LocalAddr().(*net.UDPAddr)
	lhost := laddr.IP.String()
	lport := uint16(laddr.Port)
	dls.request.Via = &Via{
		Host:   lhost,
		Port:   lport,
		Params: Params{"branch": util.GenerateBranch()},
	}
	sockMsgs := make(chan *Msg)
	sockErrs := make(chan error)
	dls.sockMsgs = sockMsgs
	dls.sockErrs = sockErrs
	go ReceiveMessages(dls.request.Contact, dls.sock, sockMsgs, sockErrs)

	// But a connected UDP socket can only receive packets from a single host.
	// SIP signalling paths can change depending on the environment, so we need
	// to be able to accept packets from anyone.
	if *looseSignalling {
		if dls.csock == nil {
			cconn, err := net.ListenPacket("udp", ":0")
			if err != nil {
				log.Printf("net.ListenPacket(udp, :0) failed: %s", err)
				return false
			}
			dls.csock = cconn.(*net.UDPConn)
			dls.request.Contact = &Addr{
				Uri: &URI{
					Scheme: "sip",
					Host:   lhost,
					Port:   uint16(dls.csock.LocalAddr().(*net.UDPAddr).Port),
					Params: Params{"transport": "udp"},
				},
			}
		} else {
			dls.request.Contact.Uri.Host = lhost
		}
		csockMsgs := make(chan *Msg)
		csockErrs := make(chan error)
		dls.csockMsgs = csockMsgs
		dls.csockErrs = csockErrs
		go ReceiveMessages(dls.request.Contact, dls.csock, csockMsgs, csockErrs)
	} else {
		dls.request.Contact = &Addr{
			Uri: &URI{
				Scheme: "sip",
				Host:   lhost,
				Port:   lport,
				Params: Params{"transport": "udp"},
			},
		}
	}

	return true
}

func (dls *dialogState) handleMessage(msg *Msg) bool {
	if msg.VersionMajor != 2 || msg.VersionMinor != 0 {
		if !dls.send(NewResponse(msg, StatusVersionNotSupported)) {
			return false
		}
		dls.errChan <- errors.New("Remote UA is using a strange SIP version")
		return false
	}
	if msg.CallID != dls.request.CallID {
		log.Printf("Received message doesn't match dialog")
		return dls.send(NewResponse(msg, StatusCallTransactionDoesNotExist))
	}
	if msg.IsResponse() {
		return dls.handleResponse(msg)
	} else {
		return dls.handleRequest(msg)
	}
}

func (dls *dialogState) handleResponse(msg *Msg) bool {
	if !ResponseMatch(dls.request, msg) {
		log.Println("Received response doesn't match transaction")
		return true
	}
	if msg.Status >= StatusOK && dls.request.Method == MethodInvite {
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
	if msg.Status <= StatusOK {
		dls.checkSDP(msg)
	}
	switch msg.Status {
	case StatusTrying:
		dls.transition(DialogProceeding)
	case StatusRinging, StatusSessionProgress:
		dls.transition(DialogRinging)
	case StatusOK:
		switch msg.CSeqMethod {
		case MethodInvite:
			if dls.remote == nil {
				dls.transition(DialogAnswered)
			}
			dls.remote = msg
		case MethodBye, MethodCancel:
			dls.transition(DialogHangup)
			return false
		}
	case StatusServiceUnavailable:
		if dls.request == dls.invite {
			log.Printf("Service unavailable: %s (%s)", dls.sock.RemoteAddr(), dls.dest)
			return dls.popRoute()
		} else {
			dls.errChan <- &ResponseError{Msg: msg}
			return false
		}
	case StatusMovedPermanently, StatusMovedTemporarily:
		dls.invite.Request = msg.Contact.Uri
		dls.invite.Route = nil
		return dls.sendRequest(dls.invite, true)
	default:
		if msg.Status > StatusOK {
			dls.errChan <- &ResponseError{Msg: msg}
			return false
		}
	}
	return true
}

func (dls *dialogState) handleRequest(msg *Msg) bool {
	if msg.MaxForwards <= 0 {
		if !dls.send(NewResponse(msg, StatusTooManyHops)) {
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
			return dls.send(NewResponse(msg, StatusInternalServerError))
		}
		dls.rseq = msg.CSeq
	}
	switch msg.Method {
	case MethodBye:
		if !dls.send(NewResponse(msg, StatusOK)) {
			return false
		}
		dls.transition(DialogHangup)
		return false
	case MethodOptions: // Probably a keep-alive ping.
		return dls.send(NewResponse(msg, StatusOK))
	case MethodInvite: // Re-INVITEs are used to change the RTP or signalling path.
		dls.remote = msg
		dls.checkSDP(msg)
		return dls.sendResponse(NewResponse(msg, StatusOK))
	case MethodAck: // Re-INVITE response has been ACK'd.
		dls.response = nil
		dls.responseTimer = nil
		return true
	default:
		return dls.send(NewResponse(msg, StatusMethodNotAllowed))
	}
}

func (dls *dialogState) checkSDP(msg *Msg) {
	if msg.Headers["Content-Type"] == sdp.ContentType {
		ms, err := sdp.Parse(msg.Payload)
		if err != nil {
			log.Println("Bad SDP payload:", err)
		} else {
			dls.sdpChan <- ms
		}
	}
}

func (dls *dialogState) send(msg *Msg) bool {
	if msg.MaxForwards > 0 {
		msg.MaxForwards--
		if msg.MaxForwards == 0 {
			dls.errChan <- errors.New("Local froot loop detected")
			return false
		}
	}
	ts := time.Now()
	addTimestamp(msg, ts)
	var b bytes.Buffer
	msg.Append(&b)
	if *tracing {
		trace("send", b.String(), dls.sock.RemoteAddr(), ts)
	}
	_, err := dls.sock.Write(b.Bytes())
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
	if dls.requestResends < maxResends {
		if !dls.send(dls.request) {
			return false
		}
		dls.requestResends++
		dls.requestTimer = time.After(resendInterval)
	} else {
		log.Printf("Timeout: %s (%s)", dls.sock.RemoteAddr(), dls.dest)
		if !dls.popRoute() {
			return false
		}
	}
	return true
}

// sendResponse is used to reliably send a response to an INVITE only.
func (dls *dialogState) sendResponse(msg *Msg) bool {
	dls.response = msg
	dls.responseResends = 0
	dls.responseTimer = time.After(resendInterval)
	return dls.send(dls.response)
}

func (dls *dialogState) resendResponse() bool {
	if dls.response == nil {
		return true
	}
	if dls.responseResends < maxResends {
		if !dls.send(dls.response) {
			return false
		}
		dls.responseResends++
		dls.responseTimer = time.After(resendInterval)
	} else {
		// TODO(jart): If resending INVITE 200 OK, start sending BYE.
		log.Printf("Timeout sending response: %s (%s)", dls.sock.RemoteAddr(), dls.dest)
		if !dls.popRoute() {
			return false
		}
	}
	return true
}

func (dls *dialogState) sendHangup() bool {
	switch dls.state {
	case DialogProceeding, DialogRinging:
		return dls.send(NewCancel(dls.invite))
	case DialogAnswered:
		return dls.sendRequest(NewBye(dls.invite, dls.remote, &dls.lseq), false)
	case DialogHangup:
		panic("Why didn't the event loop break?")
	default:
		//  o  A UA or proxy cannot send CANCEL for a transaction until it gets a
		//     provisional response for the request.  This was allowed in RFC 2543
		//     but leads to potential race conditions.
		dls.transition(DialogHangup)
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
}

func (dls *dialogState) cleanupSock() {
	if dls.sock != nil {
		dls.sock.Close()
		dls.sock = nil
		_, _ = <-dls.sockMsgs
		<-dls.sockErrs
	}
}

func (dls *dialogState) cleanupCSock() {
	if dls.csock != nil {
		dls.csock.Close()
		dls.csock = nil
		_, _ = <-dls.csockMsgs
		<-dls.csockErrs
	}
}

func (dls *dialogState) sabotage() {
	close(dls.errChan)
	close(dls.sdpChan)
	close(dls.stateChan)
}
