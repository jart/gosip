// Echo test that uses slightly higher level APIs.

package echo2_test

import (
	"github.com/jart/gosip/dsp"
	"github.com/jart/gosip/rtp"
	"github.com/jart/gosip/sip"
	"net"
	"testing"
	"time"
)

func TestCallToEchoApp(t *testing.T) {
	invite := &sip.Msg{
		Method:  sip.MethodInvite,
		Request: &sip.URI{User: "echo", Host: "127.0.0.1", Port: 5060},
	}

	// Create RTP audio session.
	rs, err := rtp.NewSession("")
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Close()
	rtpPort := uint16(rs.Sock.LocalAddr().(*net.UDPAddr).Port)

	// Create a SIP phone call.
	dl, err := sip.NewDialog(invite, rtpPort)
	if err != nil {
		t.Fatal(err)
	}

	// We're going to send white noise every 20ms.
	var frame rtp.Frame
	awgn := dsp.NewAWGN(-45.0)
	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()

	// Hangup after 200ms.
	death := time.After(200 * time.Millisecond)

	// Let's GO!
	var answered bool
	for {
		select {
		case <-ticker.C:
			for n := 0; n < 160; n++ {
				frame[n] = awgn.Get()
			}
			if err := rs.Send(&frame); err != nil {
				t.Fatal("RTP send failed:", err)
			}
		case err := <-dl.OnErr:
			t.Error(err)
			return
		case state := <-dl.OnState:
			switch state {
			case sip.DialogAnswered:
				answered = true
			case sip.DialogHangup:
				if !answered {
					t.Error("Call didn't get answered!")
				}
				return
			}
		case rs.Peer = <-dl.OnPeer:
		case frame := <-rs.C:
			rs.R <- frame
		case err := <-rs.E:
			t.Error("RTP recv failed:", err)
			rs.CloseAfterError()
			dl.Hangup <- true
		case <-death:
			dl.Hangup <- true
		}
	}
}
