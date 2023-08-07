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

// Echo test that uses slightly higher level APIs.

package echo3_test

import (
	"net"
	"testing"
	"time"

	"github.com/cresta/gosip/dialog"
	"github.com/cresta/gosip/dsp"
	"github.com/cresta/gosip/rtp"
	"github.com/cresta/gosip/sdp"
	"github.com/cresta/gosip/sip"
	"github.com/cresta/gosip/util"
)

func TestCallToEchoApp(t *testing.T) {
	// Create RTP audio session.
	rs, err := rtp.NewSession("")
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Close()
	rtpPort := uint16(rs.Sock.LocalAddr().(*net.UDPAddr).Port)

	invite := &sip.Msg{
		Method:  sip.MethodInvite,
		Request: &sip.URI{User: "echo", Host: "127.0.0.1", Port: 5060},
		Payload: &sdp.SDP{
			Origin: sdp.Origin{ID: util.GenerateOriginID()},
			Audio: &sdp.Media{
				Port:   rtpPort,
				Codecs: []sdp.Codec{sdp.ULAWCodec, sdp.DTMFCodec},
			},
		},
	}

	// Create a SIP phone call.
	dl, err := dialog.NewDialog(invite)
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
			case dialog.Answered:
				answered = true
			case dialog.Hangup:
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
