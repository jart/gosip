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

// Example demonstrating how to make a phone call (SIP/SDP/RTP/DTMF) to an Echo
// application without any higher level SIP abstractions.
//
// Example Trace:
//
// send 765 bytes to udp/[10.11.34.37]:5060 at 15:46:48.569658:
//    ------------------------------------------------------------------------
//    INVITE sip:10.11.34.37 SIP/2.0
//    Via: SIP/2.0/UDP 10.11.34.37:59516;rport;branch=z9hG4bKS308QB9UUpNyD
//    Max-Forwards: 70
//    From: <sip:10.11.34.37:59516>;tag=S1jX7UtK5Zerg
//    To: <sip:10.11.34.37>
//    Call-ID: 87704115-03b8-122e-08b5-001bfcce6bdf
//    CSeq: 133097268 INVITE
//    Contact: <sip:10.11.34.37:59516>
//    User-Agent: tube/0.1
//    Allow: INVITE, ACK, BYE, CANCEL, OPTIONS, PRACK, MESSAGE, SUBSCRIBE, NOTIFY, REFER, UPDATE, INFO
//    Supported: timer, 100rel
//    Allow-Events: talk
//    Content-Type: application/sdp
//    Content-Disposition: session
//    Content-Length: 218
//
//    v=0
//    o=- 2862054018559638081 6057228511765453924 IN IP4 10.11.34.37
//    s=-
//    c=IN IP4 10.11.34.37
//    t=0 0
//    m=audio 23448 RTP/AVP 0 101
//    a=rtpmap:0 PCMU/8000
//    a=rtpmap:101 telephone-event/8000
//    a=fmtp:101 0-16
//    a=ptime:20
//    ------------------------------------------------------------------------
// recv 282 bytes from udp/[10.11.34.37]:5060 at 15:46:48.570890:
//    ------------------------------------------------------------------------
//    SIP/2.0 100 Trying
//    Via: SIP/2.0/UDP 10.11.34.37:59516;rport=59516;branch=z9hG4bKS308QB9UUpNyD
//    From: <sip:10.11.34.37:59516>;tag=S1jX7UtK5Zerg
//    To: <sip:10.11.34.37>
//    Call-ID: 87704115-03b8-122e-08b5-001bfcce6bdf
//    CSeq: 133097268 INVITE
//    User-Agent: tube/0.1
//    Content-Length: 0
//
//    ------------------------------------------------------------------------
// recv 668 bytes from udp/[10.11.34.37]:5060 at 15:46:48.571844:
//    ------------------------------------------------------------------------
//    SIP/2.0 200 OK
//    Via: SIP/2.0/UDP 10.11.34.37:59516;rport=59516;branch=z9hG4bKS308QB9UUpNyD
//    From: <sip:10.11.34.37:59516>;tag=S1jX7UtK5Zerg
//    To: <sip:10.11.34.37>;tag=a1vFUD7vvK4ZN
//    Call-ID: 87704115-03b8-122e-08b5-001bfcce6bdf
//    CSeq: 133097268 INVITE
//    Contact: <sip:10.11.34.37>
//    User-Agent: tube/0.1
//    Accept: application/sdp
//    Allow: INVITE, ACK, BYE, CANCEL, OPTIONS
//    Supported: timer
//    Content-Type: application/sdp
//    Content-Disposition: session
//    Content-Length: 196
//
//    v=0
//    o=- 403551387931241779 5960509760717556241 IN IP4 10.11.34.37
//    s=-
//    c=IN IP4 10.11.34.37
//    t=0 0
//    m=audio 19858 RTP/AVP 0
//    a=rtpmap:0 PCMU/8000
//    a=rtmap:97 speex/16000
//    a=rtmap:98 speex/8000
//    ------------------------------------------------------------------------
// send 296 bytes to udp/[10.11.34.37]:5060 at 15:46:48.572320:
//    ------------------------------------------------------------------------
//    ACK sip:10.11.34.37 SIP/2.0
//    Via: SIP/2.0/UDP 10.11.34.37:59516;rport;branch=z9hG4bKtct1S6SZrZBHS
//    Max-Forwards: 70
//    From: <sip:10.11.34.37:59516>;tag=S1jX7UtK5Zerg
//    To: <sip:10.11.34.37>;tag=a1vFUD7vvK4ZN
//    Call-ID: 87704115-03b8-122e-08b5-001bfcce6bdf
//    CSeq: 133097268 ACK
//    Content-Length: 0
//
//    ------------------------------------------------------------------------
// send 1017 bytes to udp/[10.11.34.37]:5060 at 15:46:53.048617:
//    ------------------------------------------------------------------------
//    BYE sip:10.11.34.37 SIP/2.0
//    Via: SIP/2.0/UDP 10.11.34.37:59516;rport;branch=z9hG4bKUNKtU1a3N813m
//    Max-Forwards: 70
//    From: <sip:10.11.34.37:59516>;tag=S1jX7UtK5Zerg
//    To: <sip:10.11.34.37>;tag=a1vFUD7vvK4ZN
//    Call-ID: 87704115-03b8-122e-08b5-001bfcce6bdf
//    CSeq: 133097269 BYE
//    User-Agent: tube/0.1
//    Content-Type: text/plain
//    Content-Length: 671
//
//              .--.                              .--.
//            .'(`                               /  ..\
//         __.>\ '.  _.---,._,'             ____.'  _o/
//       /.--.  : |/' _.--.<                '--.     |.__
//    _..-'    `\     /'     `'             _.-'     /--'
//     >_.-``-. `Y  /' _.---._____     _.--'        /
//    '` .-''. \|:  \.'   ___, .-'`   ~'--....___.-'
//     .'--._ `-:  \/   /'    \\
//         /.'`\ :;    /'       `-.
//        -`    |     |
//              :.; : |           thank you for flying tube
//              |:    |                version o.1
//              |     |
//              :. :  |        besiyata dishmaya
//            .jgs    ;             telecommunications inc.
//            /:::.    `\
//    ------------------------------------------------------------------------
// recv 353 bytes from udp/[10.11.34.37]:5060 at 15:46:53.049140:
//    ------------------------------------------------------------------------
//    SIP/2.0 200 OK
//    Via: SIP/2.0/UDP 10.11.34.37:59516;rport=59516;branch=z9hG4bKUNKtU1a3N813m
//    From: <sip:10.11.34.37:59516>;tag=S1jX7UtK5Zerg
//    To: <sip:10.11.34.37>;tag=a1vFUD7vvK4ZN
//    Call-ID: 87704115-03b8-122e-08b5-001bfcce6bdf
//    CSeq: 133097269 BYE
//    User-Agent: tube/0.1
//    Allow: INVITE, ACK, BYE, CANCEL, OPTIONS
//    Supported: timer
//    Content-Length: 0
//
//    ------------------------------------------------------------------------
//

package echo_test

import (
	"bytes"
	"log"
	"math/rand"
	"net"
	"testing"
	"time"

	"github.com/cresta/gosip/rtp"
	"github.com/cresta/gosip/sdp"
	"github.com/cresta/gosip/sip"
	"github.com/cresta/gosip/util"
)

func TestCallToEchoApp(t *testing.T) {
	// Connect to the remote SIP UDP endpoint.
	conn, err := net.Dial("udp", "127.0.0.1:5060")
	if err != nil {
		t.Error("sip dial:", err)
		return
	}
	defer conn.Close()
	raddr := conn.RemoteAddr().(*net.UDPAddr)
	laddr := conn.LocalAddr().(*net.UDPAddr)

	// Create an RTP socket.
	rtpsock, err := net.ListenPacket("udp", "108.61.60.146:0")
	if err != nil {
		t.Error("rtp listen:", err)
		return
	}
	defer rtpsock.Close()
	rtpaddr := rtpsock.LocalAddr().(*net.UDPAddr)

	// Create an invite message and attach the SDP.
	invite := &sip.Msg{
		CallID:     util.GenerateCallID(),
		CSeq:       util.GenerateCSeq(),
		Method:     "INVITE",
		CSeqMethod: "INVITE",
		Request: &sip.URI{
			Scheme: "sip",
			User:   "echo",
			Host:   raddr.IP.String(),
			Port:   uint16(raddr.Port),
		},
		Via: &sip.Via{
			Host:  laddr.IP.String(),
			Port:  uint16(laddr.Port),
			Param: &sip.Param{Name: "branch", Value: util.GenerateBranch()},
		},
		From: &sip.Addr{
			Display: "Echo Test",
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   laddr.IP.String(),
				Port:   uint16(laddr.Port),
			},
			Param: &sip.Param{Name: "tag", Value: util.GenerateTag()},
		},
		To: &sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   raddr.IP.String(),
				Port:   uint16(raddr.Port),
			},
		},
		Contact: &sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   laddr.IP.String(),
				Port:   uint16(laddr.Port),
			},
		},
		UserAgent: "gosip/1.o",
		Payload:   sdp.New(rtpaddr, sdp.ULAWCodec, sdp.DTMFCodec),
	}

	// Turn invite message into a packet and send via UDP socket.
	var b bytes.Buffer
	invite.Append(&b)
	log.Printf(">>> %s\n%s\n", raddr, b.String())
	if amt, err := conn.Write(b.Bytes()); err != nil || amt != b.Len() {
		t.Fatal(err)
	}

	// Receive provisional 100 Trying.
	conn.SetDeadline(time.Now().Add(time.Second))
	memory := make([]byte, 2048)
	amt, err := conn.Read(memory)
	if err != nil {
		t.Fatal("read 100 trying:", err)
	}
	log.Printf("<<< %s\n%s\n", raddr, string(memory[0:amt]))
	msg, err := sip.ParseMsg(memory[0:amt])
	if err != nil {
		t.Fatal("parse 100 trying", err)
	}
	if !msg.IsResponse() || msg.Status != 100 || msg.Phrase != "Trying" {
		t.Fatal("didn't get 100 trying :[")
	}

	// Receive 200 OK.
	conn.SetDeadline(time.Now().Add(5 * time.Second))
	amt, err = conn.Read(memory)
	if err != nil {
		t.Fatal("read 200 ok:", err)
	}
	log.Printf("<<< %s\n%s\n", raddr, string(memory[0:amt]))
	msg, err = sip.ParseMsg(memory[0:amt])
	if err != nil {
		t.Fatal("parse 200 ok:", err)
	}
	if !msg.IsResponse() || msg.Status != 200 || msg.Phrase != "OK" {
		t.Fatal("wanted 200 ok but got:", msg.Status, msg.Phrase)
	}

	// Figure out where they want us to send RTP.
	var rrtpaddr *net.UDPAddr
	if ms, ok := msg.Payload.(*sdp.SDP); ok {
		rrtpaddr = &net.UDPAddr{IP: net.ParseIP(ms.Addr), Port: int(ms.Audio.Port)}
	} else {
		t.Fatal("200 ok didn't have sdp payload")
	}

	// Acknowledge the 200 OK to answer the call.
	var ack sip.Msg
	ack.Request = invite.Request
	ack.From = msg.From
	ack.To = msg.To
	ack.CallID = msg.CallID
	ack.Method = "ACK"
	ack.CSeq = msg.CSeq
	ack.CSeqMethod = "ACK"
	ack.Via = msg.Via
	b.Reset()
	ack.Append(&b)
	if amt, err := conn.Write(b.Bytes()); err != nil || amt != b.Len() {
		t.Fatal(err)
	}

	// Send RTP packets containing junk until we get an echo response.
	quit := make(chan bool)
	go func() {
		frameout := make([]byte, rtp.HeaderSize+160)
		rtpHeader := rtp.Header{
			PT:   sdp.ULAWCodec.PT,
			Seq:  666,
			TS:   0,
			Ssrc: rand.Uint32(),
		}
		for n := 0; n < 160; n++ {
			frameout[rtp.HeaderSize+n] = byte(n)
		}
		ticker := time.NewTicker(20 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				rtpHeader.Write(frameout)
				rtpHeader.TS += 160
				rtpHeader.Seq++
				amt, err = rtpsock.WriteTo(frameout, rrtpaddr)
				if err != nil {
					t.Fatal("rtp write", err)
				}
			case <-quit:
				return
			}
		}
	}()
	defer func() { quit <- true }()

	// We're talking to an echo application so they should send us back exactly
	// the same audio.
	rtpsock.SetDeadline(time.Now().Add(5 * time.Second))
	amt, _, err = rtpsock.ReadFrom(memory)
	if err != nil {
		t.Fatal("rtp read", err)
	}
	if amt != rtp.HeaderSize+160 {
		t.Fatal("rtp recv amt != 12+160")
	}
	var rtpHeader rtp.Header
	err = rtpHeader.Read(memory)
	if err != nil {
		t.Fatal(err)
	}
	for n := 0; n < 160; n++ {
		if memory[rtp.HeaderSize+n] != byte(n) {
			t.Fatal("rtp response audio didnt match")
		}
	}

	// Hangup (we'll be lazy and just change up the ack Msg)
	ack.Method = "BYE"
	ack.CSeqMethod = "BYE"
	ack.CSeq++
	b.Reset()
	ack.Append(&b)
	amt, err = conn.Write(b.Bytes())
	if err != nil || amt != b.Len() {
		t.Fatal(err)
	}

	// Wait for acknowledgment of hangup.
	conn.SetDeadline(time.Now().Add(time.Second))
	amt, err = conn.Read(memory)
	if err != nil {
		t.Fatal(err)
	}
	msg, err = sip.ParseMsg(memory[0:amt])
	if err != nil {
		t.Fatal(err)
	}
	if !msg.IsResponse() || msg.Status != 200 || msg.Phrase != "OK" {
		t.Fatal("wanted bye response 200 ok but got:", msg.Status, msg.Phrase)
	}
}
