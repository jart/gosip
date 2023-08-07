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

// SIP Transport Layer.  Responsible for serializing messages to/from
// your network.

package dialog

import (
	"bytes"
	"net"
	"time"

	"github.com/cresta/gosip/sip"
	"github.com/cresta/gosip/util"
)

// Transport sends and receives SIP messages over UDP with stateless routing.
type Transport struct {
	C       <-chan *sip.Msg
	E       <-chan error
	Sock    *net.UDPConn
	Via     *sip.Via
	Contact *sip.Addr
}

// Creates a new stateless network mechanism for transmitting and receiving SIP
// signalling messages.
//
// contact is a SIP address, e.g. "<sip:1.2.3.4>", that tells how to bind
// sockets. If contact.Uri.Port is 0, it'll be mutated with a randomly selected
// port. This value is also used for contact headers which tell other
// user-agents where to send responses and hence should only contain an IP or
// canonical address.
func NewTransport(contact *sip.Addr) (tp *Transport, err error) {
	saddr := net.JoinHostPort(contact.Uri.Host, util.Portstr(contact.Uri.Port))
	conn, err := net.ListenPacket("udp", saddr)
	if err != nil {
		return nil, err
	}
	sock := conn.(*net.UDPConn)
	addr := conn.LocalAddr().(*net.UDPAddr)
	contact = contact.Copy()
	contact.Next = nil
	contact.Uri.Port = uint16(addr.Port)
	contact.Uri.Param = &sip.URIParam{
		Name:  "transport",
		Value: "udp",
		Next:  contact.Uri.Param,
	}
	c := make(chan *sip.Msg)
	e := make(chan error)
	tp = &Transport{
		C:       c,
		E:       e,
		Sock:    sock,
		Contact: contact,
		Via: &sip.Via{
			Host: addr.IP.String(),
			Port: uint16(addr.Port),
		},
	}
	go ReceiveMessages(sock, c, e)
	return
}

// Sends a SIP message.
func (tp *Transport) Send(msg *sip.Msg) error {
	PopulateMessage(tp.Via, tp.Contact, msg)
	host, port, err := RouteMessage(tp.Via, tp.Contact, msg)
	if err != nil {
		return err
	}
	addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, util.Portstr(port)))
	if err != nil {
		return err
	}
	if msg.MaxForwards > 0 {
		msg.MaxForwards--
	}
	ts := time.Now()
	addTimestamp(msg, ts)
	var b bytes.Buffer
	msg.Append(&b)
	if *tracing {
		trace("send", b.Bytes(), addr)
	}
	_, err = tp.Sock.WriteTo(b.Bytes(), addr)
	if err != nil {
		return err
	}
	return nil
}
