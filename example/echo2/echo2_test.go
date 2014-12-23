// Echo test that uses slightly higher level APIs.

package echo2_test

import (
	"github.com/jart/gosip/rtp"
	"github.com/jart/gosip/sdp"
	"github.com/jart/gosip/sip"
	"log"
	"math/rand"
	"net"
	"testing"
	"time"
)

func TestCallToEchoApp(t *testing.T) {
	to := &sip.Addr{Uri: &sip.URI{Host: "127.0.0.1", Port: 5060}}
	from := &sip.Addr{Uri: &sip.URI{Host: "127.0.0.1"}}

	// Create the SIP UDP transport layer.
	tp, err := sip.NewTransport(from)
	if err != nil {
		t.Fatal(err)
	}
	defer tp.Sock.Close()

	// Create an RTP socket.
	rtpsock, err := net.ListenPacket("udp", "108.61.60.146:0")
	if err != nil {
		t.Fatal("rtp listen:", err)
	}
	defer rtpsock.Close()
	rtpaddr := rtpsock.LocalAddr().(*net.UDPAddr)

	// Send an INVITE message with an SDP.
	invite := sip.NewRequest(tp, "INVITE", to, from)
	sip.AttachSDP(invite, sdp.New(rtpaddr, sdp.ULAWCodec, sdp.DTMFCodec))
	err = tp.Send(invite)
	if err != nil {
		t.Fatal(err)
	}

	// Consume provisional messages until we receive answer.
	var rrtpaddr *net.UDPAddr
	for {
		tp.Sock.SetDeadline(time.Now().Add(time.Second))
		msg := tp.Recv()
		if msg.Error != nil {
			t.Fatal(msg.Error)
		}
		if msg.Status < sip.StatusOK {
			log.Printf("Got provisional %d %s", msg.Status, msg.Phrase)
		}
		if msg.Headers["Content-Type"] == "application/sdp" {
			log.Printf("Establishing media session")
			rsdp, err := sdp.Parse(msg.Payload)
			if err != nil {
				t.Fatal("failed to parse sdp", err)
			}
			rrtpaddr = &net.UDPAddr{IP: net.ParseIP(rsdp.Addr), Port: int(rsdp.Audio.Port)}
		}
		if msg.Status == sip.StatusOK {
			log.Printf("Answered!")
			break
		} else if msg.Status > sip.StatusOK {
			t.Fatalf("Got %d %s", msg.Status, msg.Phrase)
			NewAck(invite)
		}
	}

	// Figure out where they want us to send RTP.

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
	msg, err = sip.ParseMsg(string(memory[0:amt]))
	if err != nil {
		t.Fatal(err)
	}
	if !msg.IsResponse || msg.Status != 200 || msg.Phrase != "OK" {
		t.Fatal("wanted bye response 200 ok but got:", msg.Status, msg.Phrase)
	}
}
