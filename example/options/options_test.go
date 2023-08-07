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

// Example demonstrating how to ping a server with an OPTIONS message.

package options_test

import (
	"bytes"
	"net"
	"testing"
	"time"

	"github.com/cresta/gosip/sip"
	"github.com/cresta/gosip/util"
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
		Accept:     "application/sdp",
		UserAgent:  "pokémon/1.o",
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
			Param:    &sip.Param{Name: "branch", Value: util.GenerateBranch()},
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
			Param: &sip.Param{Name: "tag", Value: util.GenerateTag()},
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

	msg, err := sip.ParseMsg(memory[0:amt])
	if err != nil {
		t.Fatal(err)
	}

	if !msg.IsResponse() || msg.Status != 200 || msg.Phrase != "OK" {
		t.Error("Not OK :[")
	}
	if options.CallID != msg.CallID {
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
