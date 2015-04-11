// Example demonstrating how to ping a server with an OPTIONS message.

package options_test

import (
	"bytes"
	"github.com/jart/gosip/sip"
	"github.com/jart/gosip/util"
	"net"
	"testing"
	"time"
)

func TestOptions(t *testing.T) {
	sock, err := net.Dial("udp", "127.0.0.1:5060")
	if err != nil {
		t.Error(err)
		return
	}
	defer sock.Close()
	raddr := sock.RemoteAddr().(*net.UDPAddr)
	laddr := sock.LocalAddr().(*net.UDPAddr)

	options := sip.Msg{
		CSeq:       util.GenerateCSeq(),
		CallID:     util.GenerateCallID(),
		Method:     "OPTIONS",
		CSeqMethod: "OPTIONS",
		Accept:     []byte("application/sdp"),
		UserAgent:  []byte("pokémon/1.o"),
		Request: &sip.URI{
			Scheme: "sip",
			User:   "echo",
			Host:   raddr.IP.String(),
			Port:   uint16(raddr.Port),
		},
		Via: &sip.Via{
			Version:  "2.0",
			Protocol: "UDP",
			Host:     laddr.IP.String(),
			Port:     uint16(laddr.Port),
			Param:    &sip.Param{"branch", util.GenerateBranch(), nil},
		},
		Contact: &sip.Addr{
			Uri: &sip.URI{
				Host: laddr.IP.String(),
				Port: uint16(laddr.Port),
			},
		},
		From: &sip.Addr{
			Uri: &sip.URI{
				User: "gosip",
				Host: "justinetunney.com",
				Port: 5060,
			},
			Param: &sip.Param{"tag", util.GenerateTag(), nil},
		},
		To: &sip.Addr{
			Uri: &sip.URI{
				Host: raddr.IP.String(),
				Port: uint16(raddr.Port),
			},
		},
	}

	var b bytes.Buffer
	options.Append(&b)
	if amt, err := sock.Write(b.Bytes()); err != nil || amt != b.Len() {
		t.Fatal(err)
	}

	memory := make([]byte, 2048)
	sock.SetDeadline(time.Now().Add(time.Second))
	amt, err := sock.Read(memory)
	if err != nil {
		t.Fatal(err)
	}

	msg, err := sip.ParseMsg(string(memory[0:amt]))
	if err != nil {
		t.Fatal(err)
	}

	if !msg.IsResponse() || msg.Status != 200 || msg.Phrase != "OK" {
		t.Error("Not OK :[")
	}
	if !bytes.Equal(options.CallID, msg.CallID) {
		t.Error("CallID didnt match")
	}
	if options.CSeq != msg.CSeq || options.CSeqMethod != msg.CSeqMethod {
		t.Error("CSeq didnt match")
	}
	if options.From.String() != msg.From.String() {
		t.Errorf("From headers didn't match:\n%s\n%s\n\n%s", options.From, msg.From, memory[0:amt])
	}
	if msg.To.Param.Get("tag") == nil {
		t.Errorf("Remote UA didnt tag To header:\n%s\n\n%s", msg.To, memory[0:amt])
	}
	msg.To.Param = nil
	if options.To.String() != msg.To.String() {
		t.Errorf("To headers didn't match:\n%s\n%s\n\n%s", options.To, msg.To, memory[0:amt])
	}
}
