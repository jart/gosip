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

// This code demonstrates how SIP works without the zillion layers of
// sugar coating.

package rawsip_test

import (
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/cresta/gosip/util"
)

// An 'OPTIONS' message is used to:
//
// - Ping a server to see if it's alive.
// - Keep a connection alive in nat situations.
// - Ask a user agent what features they support.
//
func TestRawSIPOptions(t *testing.T) {
	// create a new udp socket bound to a random port and "connect"
	// the socket to the remote address
	raddr := "127.0.0.1:5060"
	conn, err := net.Dial("udp", raddr)
	if err != nil {
		t.Error(err)
		return
	}

	// What local ip/port binding did the kernel choose for us?
	laddr := conn.LocalAddr().String()

	// Construct a SIP message.
	cseq := util.GenerateCSeq()
	fromtag := util.GenerateTag()
	callid := util.GenerateCallID()
	packet := "" +
		"OPTIONS sip:echo@" + laddr + " SIP/2.0\r\n" +
		"Via: SIP/2.0/UDP " + laddr + "\r\n" +
		"Max-Forwards: 70\r\n" +
		"To: <sip:" + raddr + ">\r\n" +
		"From: <sip:" + laddr + ">;tag=" + fromtag + "\r\n" +
		"Call-ID: " + callid + "\r\n" +
		"CSeq: " + strconv.Itoa(cseq) + " OPTIONS\r\n" +
		"Contact: <sip:" + laddr + ">\r\n" +
		"User-Agent: pokémon/1.o\r\n" +
		"Accept: application/sdp\r\n" +
		"Content-Length: 0\r\n" +
		"\r\n"

	// Transmit our message.
	bpacket := []uint8(packet)
	amt, err := conn.Write(bpacket)
	if err != nil || amt != len(bpacket) {
		t.Error(err)
		return
	}

	// Wait no longer than a second for them to get back to us.
	err = conn.SetDeadline(time.Now().Add(time.Second))
	if err != nil {
		t.Error(err)
		return
	}

	// Receive response.
	buf := make([]byte, 2048)
	amt, err = conn.Read(buf)
	if err != nil {
		t.Error(err)
		return
	}
	response := buf[0:amt]

	// Read response.
	msg := string(response)
	lines := strings.Split(msg, "\r\n")
	if lines[0] != "SIP/2.0 200 OK" {
		t.Errorf("not ok :[\n%s", msg)
		return
	}
}
