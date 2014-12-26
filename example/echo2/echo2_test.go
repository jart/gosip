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
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	to := &sip.Addr{Uri: &sip.URI{User: "echo", Host: "127.0.0.1", Port: 5060}}
	from := &sip.Addr{Uri: &sip.URI{Host: "127.0.0.1"}}

	// Create an RTP media session.
	rs, err := rtp.NewSession(from.Uri.Host)
	if err != nil {
		t.Fatal("rtp listen:", err)
	}
	defer rs.Sock.Close()
	rtpaddr := rs.Sock.LocalAddr().(*net.UDPAddr)

	// Create an RTP audio sender.
	rtpPeer := make(chan *net.UDPAddr, 1)
	go func() {
		var frame rtp.Frame
		awgn := dsp.NewAWGN(-25.0)
		ticker := time.NewTicker(20 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case rs.Peer = <-rtpPeer:
			case <-ticker.C:
				// Send an audio frame containing comfort noise.
				for n := 0; n < 160; n++ {
					frame[n] = awgn.Get()
				}
				err := rs.Send(&frame)
				if err != nil {
					if !util.IsUseOfClosed(err) {
						t.Error("rtp write", err)
					}
					return
				}
			}
		}
	}()

	// Create the SIP UDP transport layer.
	tp, err := sip.NewTransport(from)
	if err != nil {
		t.Fatal(err)
	}
	defer tp.Sock.Close()

	// Send an INVITE message with an SDP.
	invite := sip.NewRequest(tp, sip.MethodInvite, to, from)
	sip.AttachSDP(invite, sdp.New(rtpaddr, sdp.ULAWCodec, sdp.DTMFCodec))
	err = tp.Send(invite)
	if err != nil {
		t.Fatal(err)
	}

	// Set up some reliability in case a UDP packet drops.
	resend := invite
	resends := 0
	resendInterval := 50 * time.Millisecond
	resendTimer := time.After(resendInterval)

	// Hangup after two seconds.
	deathTimer := time.After(200 * time.Millisecond)

	// Create a SIP dialog handler.
	var answered bool
	var msg *sip.Msg
loop:
	for {
		select {
		case err = <-rs.E:
			t.Fatal("RTP network error", err)
		case err = <-tp.E:
			t.Fatal("SIP network error", err)
		case <-rs.C:
			// Do nothing with inbound audio.
		case msg = <-tp.C:
			if msg.IsResponse {
				if msg.Status >= sip.StatusOK && msg.CSeq == invite.CSeq {
					err = tp.Send(sip.NewAck(invite, msg))
					if err != nil {
						t.Fatal(err)
					}
				}
				if msg.Status < sip.StatusOK {
					if msg.Status == sip.StatusTrying {
						log.Printf("Remote SIP endpoint exists!")
						resendTimer = nil
					} else if msg.Status == sip.StatusRinging {
						log.Printf("Ringing!")
					} else if msg.Status == sip.StatusSessionProgress {
						log.Printf("Probably Ringing!")
					} else {
						log.Printf("Provisional %d %s", msg.Status, msg.Phrase)
					}
				} else if msg.Status == sip.StatusOK {
					if msg.CSeqMethod == sip.MethodInvite {
						log.Printf("Answered!")
						answered = true
					} else if msg.CSeqMethod == sip.MethodBye {
						log.Printf("Hungup!")
						break loop
					} else if msg.CSeqMethod == sip.MethodCancel {
						log.Printf("Cancelled!")
						break loop
					}
				} else if msg.Status > sip.StatusOK {
					t.Errorf("Got %d %s", msg.Status, msg.Phrase)
					return
				}
				if msg.Headers["Content-Type"] == sdp.ContentType {
					log.Printf("Establishing media session")
					ms, err := sdp.Parse(msg.Payload)
					if err != nil {
						t.Fatal("failed to parse sdp", err)
					}
					rtpPeer <- &net.UDPAddr{IP: net.ParseIP(ms.Addr), Port: int(ms.Audio.Port)}
				}
			} else {
				if msg.Method == "BYE" {
					log.Printf("Remote Hangup!")
					err = tp.Send(sip.NewResponse(invite, sip.StatusOK))
					if err != nil {
						t.Fatal(err)
					}
					break loop
				}
			}
		case <-resendTimer:
			if resends == 2 {
				t.Fatal("Failed to send", resend.Method)
			}
			resends++
			err = tp.Send(resend)
			if err != nil {
				t.Fatal(err)
			}
		case <-deathTimer:
			resends = 0
			resendTimer = time.After(resendInterval)
			if answered {
				resend = sip.NewBye(invite, msg)
			} else {
				resend = sip.NewCancel(invite)
			}
			err = tp.Send(resend)
			if err != nil {
				t.Error(err)
			}
		}
	}

	// The dialog has shut down cleanly.
	if !answered {
		t.Error("Call didn't get answered!")
	}
}
