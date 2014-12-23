// Echo test that uses slightly higher level APIs.

package echo2_test

import (
	"github.com/jart/gosip/dsp"
	"github.com/jart/gosip/rtp"
	"github.com/jart/gosip/sdp"
	"github.com/jart/gosip/sip"
	"github.com/jart/gosip/util"
	"log"
	"net"
	"testing"
	"time"
)

func TestCallToEchoApp(t *testing.T) {
	to := &sip.Addr{Uri: &sip.URI{User: "echo", Host: "127.0.0.1", Port: 5060}}
	from := &sip.Addr{Uri: &sip.URI{Host: "127.0.0.1"}}

	// Create the SIP UDP transport layer.
	tp, err := sip.NewTransport(from)
	if err != nil {
		t.Fatal(err)
	}
	defer tp.Sock.Close()

	// Used to notify main thread when subthreads die.
	rtpDeath := make(chan bool, 2)

	// Create an RTP session.
	session, err := rtp.NewSession(from.Uri.Host)
	if err != nil {
		t.Fatal("rtp listen:", err)
	}
	defer session.Sock.Close()
	rtpaddr := session.Sock.LocalAddr().(*net.UDPAddr)
	rtppeerChan := make(chan *net.UDPAddr, 1)
	go func() {
		var frame rtp.Frame
		awgn := dsp.NewAWGN(-25.0)
		ticker := time.NewTicker(20 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				for n := 0; n < 160; n++ {
					frame[n] = awgn.Get()
				}
				err := session.Send(frame)
				if err != nil {
					if !util.IsUseOfClosed(err) {
						t.Error("rtp write", err)
					}
					rtpDeath <- true
					return
				}
			case session.Peer = <-rtppeerChan:
				if session.Peer == nil {
					return
				}
			}
		}
	}()
	defer func() { rtppeerChan <- nil }()

	// Create an RTP message consumer.
	go func() {
		var frame rtp.Frame
		for {
			err := session.Recv(frame)
			if err != nil {
				if !util.IsUseOfClosed(err) {
					t.Errorf("rtp read: %s %#v", err, err)
				}
				rtpDeath <- true
				return
			}
		}
	}()

	// Create a SIP message consumer.
	sipChan := make(chan *sip.Msg, 32)
	go func() {
		for {
			msg := tp.Recv()
			if msg.Error != nil && util.IsUseOfClosed(msg.Error) {
				return
			}
			sipChan <- msg
		}
	}()

	// Send an INVITE message with an SDP.
	invite := sip.NewRequest(tp, "INVITE", to, from)
	sip.AttachSDP(invite, sdp.New(rtpaddr, sdp.ULAWCodec, sdp.DTMFCodec))
	err = tp.Send(invite)
	if err != nil {
		t.Fatal(err)
	}

	// Create a SIP dialog handler.
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	var answered bool
	var msg *sip.Msg
	for {
		select {
		case <-ticker.C:
			if answered {
				err = tp.Send(sip.NewBye(invite, msg))
			} else {
				err = tp.Send(sip.NewCancel(invite))
			}
			if err != nil {
				t.Error(err)
			}
		case <-rtpDeath:
			if answered {
				err = tp.Send(sip.NewBye(invite, msg))
			} else {
				err = tp.Send(sip.NewCancel(invite))
			}
			if err != nil {
				t.Error(err)
			}
		case msg = <-sipChan:
			if msg.Error != nil {
				t.Fatal(msg.Error)
			}
			if msg.IsResponse {
				if msg.Status >= sip.StatusOK && msg.CSeq == invite.CSeq {
					err = tp.Send(sip.NewAck(invite, msg))
					if err != nil {
						t.Fatal(err)
					}
				}
				if msg.Status < sip.StatusOK {
					log.Printf("Provisional %d %s", msg.Status, msg.Phrase)
				} else if msg.Status == sip.StatusOK {
					if msg.CSeqMethod == "INVITE" {
						log.Printf("Answered!")
						answered = true
					} else if msg.CSeqMethod == "BYE" {
						log.Printf("Hungup!")
						return
					} else if msg.CSeqMethod == "CANCEL" {
						log.Printf("Cancelled!")
						return
					}
				} else if msg.Status > sip.StatusOK {
					t.Errorf("Got %d %s", msg.Status, msg.Phrase)
					return
				}
				if msg.Headers["Content-Type"] == "application/sdp" {
					log.Printf("Establishing media session")
					rsdp, err := sdp.Parse(msg.Payload)
					if err != nil {
						t.Fatal("failed to parse sdp", err)
					}
					rtppeerChan <- &net.UDPAddr{IP: net.ParseIP(rsdp.Addr), Port: int(rsdp.Audio.Port)}
				}
			} else {
				if msg.Method == "BYE" {
					log.Printf("Remote Hangup!")
					err = tp.Send(sip.NewResponse(invite, sip.StatusOK))
					if err != nil {
						t.Fatal(err)
					}
					return
				}
			}
		}
	}
}
