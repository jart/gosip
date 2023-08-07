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

package dialog

import (
	"log"
	"net"
	"strconv"
	"time"

	"github.com/cresta/gosip/sip"
)

func ReceiveMessages(sock *net.UDPConn, c chan<- *sip.Msg, e chan<- error) {
	buf := make([]byte, 2048)
	laddr := sock.LocalAddr().(*net.UDPAddr)
	lhost := laddr.IP.String()
	lport := uint16(laddr.Port)
	for {
		amt, addr, err := sock.ReadFromUDP(buf)
		if err != nil {
			e <- err
			break
		}
		ts := time.Now()
		packet := buf[0:amt]
		if *tracing {
			trace("recv", packet, addr)
		}
		msg, err := sip.ParseMsg(packet)
		if err != nil {
			log.Printf("Dropping SIP message: %s\r\n", err)
			continue
		}
		addReceived(msg, addr)
		addTimestamp(msg, ts)
		if msg.Route != nil && msg.Route.Uri.Host == lhost && or5060(msg.Route.Uri.Port) == lport {
			msg.Route = msg.Route.Next
		}
		fixMessagesFromStrictRouters(lhost, lport, msg)
		c <- msg
	}
	close(c)
	close(e) // Must be unbuffered!
}

func addReceived(msg *sip.Msg, addr *net.UDPAddr) {
	if msg.IsResponse() {
		return
	}
	if int(msg.Via.Port) != addr.Port {

		rport := msg.Via.Param.Get("rport")
		port := strconv.Itoa(addr.Port)

		if rport == nil {
			msg.Via.Param = &sip.Param{
				Name:  "rport",
				Value: port,
				Next:  msg.Via.Param,
			}
		} else {

			// implied rport is 5060, but some NAT will use another port,we use real port instead
			if len(rport.Value) == 0 {
				rport.Value = port
			}
		}
	}
	if msg.Via.Host != addr.IP.String() {
		if msg.Via.Param.Get("received") == nil {
			msg.Via.Param = &sip.Param{
				Name:  "received",
				Value: addr.IP.String(),
				Next:  msg.Via.Param,
			}
		}
	}
}

func addTimestamp(msg *sip.Msg, ts time.Time) {
	if *timestampTagging {
		msg.Via.Param = &sip.Param{
			Name:  "usi",
			Value: strconv.FormatInt(ts.UnixNano()/int64(time.Microsecond), 10),
			Next:  msg.Via.Param,
		}
	}
}

// RFC3261 16.4 Route Information Preprocessing
// RFC3261 16.12.1.2: Traversing a Strict-Routing Proxy
func fixMessagesFromStrictRouters(lhost string, lport uint16, msg *sip.Msg) {
	if msg.Request != nil &&
		msg.Request.Param.Get("lr") != nil &&
		msg.Route != nil &&
		msg.Request.Host == lhost &&
		or5060(msg.Request.Port) == lport {
		var oldReq, newReq *sip.URI
		if msg.Route.Next == nil {
			oldReq, newReq = msg.Request, msg.Route.Uri
			msg.Request = msg.Route.Uri
			msg.Route = nil
		} else {
			seclast := msg.Route
			for ; seclast.Next.Next != nil; seclast = seclast.Next {
			}
			oldReq, newReq = msg.Request, seclast.Next.Uri
			msg.Request = seclast.Next.Uri
			seclast.Next = nil
			msg.Route.Last()
		}
		log.Printf("Fixing request URI after strict router traversal: %s -> %s\r\n", oldReq, newReq)
	}
}
