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
	to := &sip.Addr{Uri: &sip.URI{User: "echo", Host: "127.0.0.1", Port: 5060}}
	from := &sip.Addr{Uri: &sip.URI{Host: "127.0.0.1"}}

	// Create the SIP UDP transport layer.
	tp, err := sip.NewTransport(from)
	if err != nil {
		t.Fatal(err)
	}
	defer tp.Sock.Close()

	// Create an RTP session.
	rtpsock, err := net.ListenPacket("udp", "108.61.60.146:0")
	if err != nil {
		t.Fatal("rtp listen:", err)
	}
	defer rtpsock.Close()
	rtpaddr := rtpsock.LocalAddr().(*net.UDPAddr)
	rrtpaddrChan := make(chan *net.UDPAddr)
	go func() {
		var rrtpaddr *net.UDPAddr
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
				if rrtpaddr != nil {
					_, err := rtpsock.WriteTo(frameout, rrtpaddr)
					if err != nil {
						t.Fatal("rtp write", err)
					}
				}
			case rrtpaddr = <-rrtpaddrChan:
				if rrtpaddr == nil {
					return
				}
			}
		}
	}()
	defer func() { rrtpaddrChan <- nil }()

	// Send an INVITE message with an SDP.
	invite := sip.NewRequest(tp, "INVITE", to, from)
	sip.AttachSDP(invite, sdp.New(rtpaddr, sdp.ULAWCodec, sdp.DTMFCodec))
	err = tp.Send(invite)
	if err != nil {
		t.Fatal(err)
	}

	// Consume provisional messages until we receive answer.
	var msg *sip.Msg
	for {
		tp.Sock.SetDeadline(time.Now().Add(time.Second))
		msg = tp.Recv()
		if msg.Error != nil {
			t.Fatal(msg.Error)
		}
		if msg.Status < sip.StatusOK {
			log.Printf("Provisional %d %s", msg.Status, msg.Phrase)
		} else if msg.Status == sip.StatusOK {
			log.Printf("Answered!")
		}
		if msg.Status >= sip.StatusOK {
			err = tp.Send(sip.NewAck(invite, msg))
			if err != nil {
				t.Fatal(err)
			}
		}
		if msg.Headers["Content-Type"] == "application/sdp" {
			log.Printf("Establishing media session")
			rsdp, err := sdp.Parse(msg.Payload)
			if err != nil {
				t.Fatal("failed to parse sdp", err)
			}
			rrtpaddrChan <- &net.UDPAddr{IP: net.ParseIP(rsdp.Addr), Port: int(rsdp.Audio.Port)}
		}
		if msg.Status == sip.StatusOK {
			break
		} else if msg.Status > sip.StatusOK {
			t.Fatalf("Got %d %s", msg.Status, msg.Phrase)
		}
	}

	// We're talking to an echo application so they should send us back exactly
	// the same audio.
	memory := make([]byte, 2048)
	rtpsock.SetDeadline(time.Now().Add(5 * time.Second))
	amt, _, err := rtpsock.ReadFrom(memory)
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
	err = tp.Send(sip.NewBye(invite, msg))
	if err != nil {
		t.Fatal(err)
	}

	// Wait for acknowledgment of hangup.
	tp.Sock.SetDeadline(time.Now().Add(time.Second))
	msg = tp.Recv()
	if msg.Error != nil {
		t.Fatal(msg.Error)
	}
	if msg.Status != 200 || msg.CSeqMethod != "BYE" {
		t.Fatal("Wanted BYE 200:", msg)
	}
}
