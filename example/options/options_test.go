// Example demonstrating how to ping a server with an OPTIONS message.

package options_test

import (
	"bytes"
	"github.com/jart/gosip/sip"
	"github.com/jart/gosip/util"
	"net"
	"reflect"
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
		Request: &sip.URI{
			Scheme: "sip",
			User:   "echo",
			Host:   raddr.IP.String(),
			Port:   uint16(raddr.Port),
		},
		Via: &sip.Via{
			Version: "2.0",
			Proto:   "UDP",
			Host:    laddr.IP.String(),
			Port:    uint16(laddr.Port),
			Params:  sip.Params{"branch": util.GenerateBranch()},
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
			Params: sip.Params{"tag": util.GenerateTag()},
		},
		To: &sip.Addr{
			Uri: &sip.URI{
				Host: raddr.IP.String(),
				Port: uint16(raddr.Port),
			},
		},
		Headers: map[string]string{
			"Accept":     "application/sdp",
			"User-Agent": "pokémon/1.o",
		},
	}

	var b bytes.Buffer
	if err := options.Append(&b); err != nil {
		t.Fatal(err)
	}
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

	if !msg.IsResponse || msg.Status != 200 || msg.Phrase != "OK" {
		t.Error("Not OK :[")
	}
	if options.CallID != msg.CallID {
		t.Error("CallID didnt match")
	}
	if options.CSeq != msg.CSeq || options.CSeqMethod != msg.CSeqMethod {
		t.Error("CSeq didnt match")
	}
	if !reflect.DeepEqual(options.From, msg.From) {
		t.Error("From headers didn't match:\n%#v\n%#v", options.From, msg.From)
	}
	if _, ok := msg.To.Params["tag"]; !ok {
		t.Error("Remote UA didnt tag To header")
	}
	msg.To.Params = nil
	if !reflect.DeepEqual(options.To, msg.To) {
		t.Error("To mismatch:\n%#v\n%#v", options.To, msg.To)
	}
}
