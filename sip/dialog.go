// SIP Transport Layer.  Responsible for serializing messages to/from
// your network.

package sip

import (
	"bytes"
	"errors"
	"github.com/jart/gosip/sdp"
	"github.com/jart/gosip/util"
	"log"
	"net"
	"time"
)

const (
	DialogConnected = 1
	DialogRinging   = 2
	DialogAnswered  = 3
	DialogHangup    = 4
	resendInterval  = 200 * time.Millisecond
	maxResends      = 2
)

// Dialog represents an outbound phone call.
type Dialog struct {
	OnErr   <-chan error
	OnState <-chan int
	OnSDP   <-chan *sdp.SDP
	Hangup  chan<- bool
}

type dialogState struct {
	sock         *net.UDPConn
	sockMsgs     <-chan *Msg
	sockErrs     <-chan error
	errChan      chan<- error
	sdpChan      chan<- *sdp.SDP
	stateChan    chan<- int
	doHangupChan <-chan bool
	routes       *AddressRoute
	invite       *Msg
	response     *Msg
	resend       *Msg
	resends      int
	timer        <-chan time.Time
}

// NewDialog creates a phone call.
func NewDialog(invite *Msg) (dl *Dialog, err error) {
	invite, host, port, err := RouteMessage(nil, nil, invite)
	if err != nil {
		return nil, err
	}
	routes, err := RouteAddress(host, port)
	if err != nil {
		return nil, err
	}
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
		routes:       routes,
	}
	go dls.run()
	return &Dialog{
		OnErr:   errChan,
		OnState: stateChan,
		OnSDP:   sdpChan,
		Hangup:  doHangupChan,
	}, nil
}

func (dls *dialogState) popRoute() bool {
	if dls.routes == nil {
		dls.errChan <- errors.New("failed to contact host")
		return false
	}
	dls.cleanup()
	conn, err := net.Dial("udp", dls.routes.Address)
	dls.routes = dls.routes.Next
	if err != nil {
		log.Println("net.Dial() failed:", err)
		return dls.popRoute()
	}
	dls.sock = conn.(*net.UDPConn)
	laddr := conn.LocalAddr().(*net.UDPAddr)
	lhost := laddr.IP.String()
	lport := uint16(laddr.Port)
	dls.invite.Via = &Via{
		Host:   lhost,
		Port:   lport,
		Params: Params{"branch": util.GenerateBranch()},
	}
	dls.invite.Contact = &Addr{
		Uri: &URI{
			Scheme: "sip",
			Host:   lhost,
			Port:   lport,
			Params: Params{"transport": "udp"},
		},
	}
	PopulateMessage(nil, nil, dls.invite)
	dls.resend = dls.invite
	dls.timer = time.After(resendInterval)
	dls.resends = 0
	sockMsgs := make(chan *Msg)
	sockErrs := make(chan error)
	dls.sockMsgs = sockMsgs
	dls.sockErrs = sockErrs
	go ReceiveMessages(dls.invite.Contact, dls.sock, sockMsgs, sockErrs)
	return dls.send(dls.resend)
}

func (dls *dialogState) run() {
	defer dls.sabotage()
	defer dls.cleanup()
	if !dls.popRoute() {
		return
	}
	for {
		select {
		case err := <-dls.sockErrs:
			if util.IsRefused(err) {
				if !dls.popRoute() {
					return
				}
			} else {
				dls.errChan <- err
				return
			}
		case <-dls.timer:
			if dls.resends < maxResends {
				if !dls.send(dls.resend) {
					return
				}
				dls.resends++
				dls.timer = time.After(resendInterval)
			} else {
				if !dls.popRoute() {
					return
				}
			}
		case <-dls.doHangupChan:
			if !dls.hangup() {
				return
			}
		case msg := <-dls.sockMsgs:
			if msg.CallID != dls.invite.CallID {
				continue
			}
			if msg.IsResponse {
				if msg.Status >= StatusOK && msg.CSeq == dls.invite.CSeq {
					if msg.Contact != nil {
						if !dls.send(NewAck(dls.invite, msg)) {
							return
						}
					}
					if msg.Status > StatusOK {
						dls.errChan <- errors.New(msg.Phrase)
						return
					}
				}
				switch msg.Status {
				case StatusTrying:
					dls.routes = nil
					dls.timer = nil
					dls.stateChan <- DialogConnected
				case StatusRinging, StatusSessionProgress:
					dls.stateChan <- DialogRinging
				case StatusOK:
					switch msg.CSeqMethod {
					case dls.invite.Method:
						if dls.response == nil {
							dls.stateChan <- DialogAnswered
						}
						dls.response = msg
					case MethodBye, MethodCancel:
						dls.stateChan <- DialogHangup
						return
					default:
						dls.errChan <- errors.New("Bad CSeq Method")
						return
					}
				}
				if msg.Headers["Content-Type"] == sdp.ContentType {
					ms, err := sdp.Parse(msg.Payload)
					if err != nil {
						log.Println("Bad SDP payload:", err)
					} else {
						dls.sdpChan <- ms
					}
				}
			} else {
				if msg.MaxForwards <= 0 {
					if !dls.send(NewResponse(msg, StatusTooManyHops)) {
						return
					}
					dls.errChan <- errors.New("Froot loop detected")
					return
				}
				switch msg.Method {
				case MethodBye:
					if !dls.send(NewResponse(msg, StatusOK)) {
						return
					}
					dls.stateChan <- DialogHangup
					return
				}
			}
		}
	}
}

func (dls *dialogState) send(msg *Msg) bool {
	// TODO(jart): Double-check route matches socket binding.
	if msg.MaxForwards > 0 {
		msg.MaxForwards--
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

func (dls *dialogState) hangup() bool {
	if dls.response != nil {
		dls.resend = NewBye(dls.invite, dls.response)
	} else {
		dls.resend = NewCancel(dls.invite)
	}
	if !dls.send(dls.resend) {
		return false
	}
	dls.resends = 0
	dls.timer = time.After(resendInterval)
	return true
}

func (dls *dialogState) cleanup() {
	if dls.sock != nil {
		dls.sock.Close()
		dls.sock = nil
	}
}

func (dls *dialogState) sabotage() {
	close(dls.errChan)
	close(dls.sdpChan)
	close(dls.stateChan)
}
