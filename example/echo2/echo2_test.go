// Echo test that uses slightly higher level APIs.

package echo2_test

import (
	"github.com/jart/gosip/dsp"
	"github.com/jart/gosip/rtp"
	"github.com/jart/gosip/sdp"
	"github.com/jart/gosip/sip"
	"log"
	"net"
	"testing"
	"time"
)

func TestCallToEchoApp(t *testing.T) {
	to := &sip.Addr{Uri: &sip.URI{User: "echo", Host: "127.0.0.1", Port: 5060}}
	from := &sip.Addr{Uri: &sip.URI{Host: "127.0.0.1"}}

	// Create an RTP media session.
	rs, err := rtp.NewSession(from.Uri.Host)
	if err != nil {
		t.Fatal("RTP listen failed:", err)
	}
	defer rs.Sock.Close()
	rtpaddr := rs.Sock.LocalAddr().(*net.UDPAddr)

	// Create the SIP UDP transport layer.
	tp, err := sip.NewTransport(from)
	if err != nil {
		t.Fatal(err)
	}
	defer tp.Sock.Close()

	// Send an INVITE message with an SDP media session description.
	invite := sip.NewRequest(tp, sip.MethodInvite, to, from)
	sip.AttachSDP(invite, sdp.New(rtpaddr, sdp.ULAWCodec, sdp.DTMFCodec))
	err = tp.Send(invite)
	if err != nil {
		t.Fatal(err)
	}

	// We're going to send white noise every 20ms.
	var frame rtp.Frame
	awgn := dsp.NewAWGN(-45.0)
	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()

	// Hangup after 200ms.
	deathTimer := time.After(200 * time.Millisecond)

	// Set up some reliability in case a UDP packet drops.
	resend := invite
	resends := 0
	resendInterval := 50 * time.Millisecond
	resendTimer := time.After(resendInterval)

	// Create a SIP dialog handler.
	var answered bool
	var msg *sip.Msg
loop:
	for {
		select {
		case err = <-rs.E:
			t.Fatal("RTP recv failed:", err)
		case err = <-tp.E:
			t.Fatal("SIP recv failed:", err)
		case <-rs.C:
			// Do nothing with received audio.
		case <-ticker.C:
			for n := 0; n < 160; n++ {
				frame[n] = awgn.Get()
			}
			if err := rs.Send(&frame); err != nil {
				t.Fatal("RTP send failed:", err)
			}
		case msg = <-tp.C:
			if msg.IsResponse() {
				if msg.Status >= sip.StatusOK && msg.CSeq == invite.CSeq {
					err = tp.Send(sip.NewAck(msg, invite))
					if err != nil {
						t.Fatal("SIP send failed:", err)
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
				if msg.ContentType == sdp.ContentType {
					log.Printf("Establishing media session")
					ms, err := sdp.Parse(msg.Payload)
					if err != nil {
						t.Fatal("Failed to parse SDP", err)
					}
					rs.Peer = &net.UDPAddr{IP: net.ParseIP(ms.Addr), Port: int(ms.Audio.Port)}
				}
			} else {
				if msg.Method == "BYE" {
					log.Printf("Remote Hangup!")
					err = tp.Send(sip.NewResponse(invite, sip.StatusOK))
					if err != nil {
						t.Fatal("SIP send failed:", err)
					}
					break loop
				}
			}
		case <-resendTimer:
			if resends == 2 {
				t.Fatal("Failed to send", resend.Method)
			}
			err = tp.Send(resend)
			if err != nil {
				t.Fatal("SIP send failed:", err)
			}
			resends++
			resendTimer = time.After(resendInterval)
		case <-deathTimer:
			if answered {
				resend = sip.NewBye(invite, msg, nil)
			} else {
				resend = sip.NewCancel(invite)
			}
			err = tp.Send(resend)
			if err != nil {
				t.Error("SIP send failed:", err)
			}
			resends = 0
			resendTimer = time.After(resendInterval)
		}
	}

	// The dialog has shut down cleanly. Was it answered?
	if !answered {
		t.Error("Call didn't get answered!")
	}
}
