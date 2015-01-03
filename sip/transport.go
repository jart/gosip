// SIP Transport Layer.  Responsible for serializing messages to/from
// your network.

package sip

import (
	"bytes"
	"net"
	"time"
)

// Transport sends and receives SIP messages over UDP with stateless routing.
type Transport struct {
	C       <-chan *Msg
	E       <-chan error
	Sock    *net.UDPConn
	Via     *Via
	Contact *Addr
}

// Creates a new stateless network mechanism for transmitting and receiving SIP
// signalling messages.
//
// contact is a SIP address, e.g. "<sip:1.2.3.4>", that tells how to bind
// sockets. If contact.Uri.Port is 0, it'll be mutated with a randomly selected
// port. This value is also used for contact headers which tell other
// user-agents where to send responses and hence should only contain an IP or
// canonical address.
func NewTransport(contact *Addr) (tp *Transport, err error) {
	saddr := net.JoinHostPort(contact.Uri.Host, portstr(contact.Uri.Port))
	conn, err := net.ListenPacket("udp", saddr)
	if err != nil {
		return nil, err
	}
	sock := conn.(*net.UDPConn)
	addr := conn.LocalAddr().(*net.UDPAddr)
	contact = contact.Copy()
	contact.Next = nil
	contact.Uri.Port = uint16(addr.Port)
	contact.Uri.Params["transport"] = "udp"
	c := make(chan *Msg)
	e := make(chan error)
	tp = &Transport{
		C:       c,
		E:       e,
		Sock:    sock,
		Contact: contact,
		Via: &Via{
			Host: addr.IP.String(),
			Port: uint16(addr.Port),
		},
	}
	go ReceiveMessages(contact, sock, c, e)
	return
}

// Sends a SIP message.
func (tp *Transport) Send(msg *Msg) error {
	PopulateMessage(tp.Via, tp.Contact, msg)
	host, port, err := RouteMessage(tp.Via, tp.Contact, msg)
	if err != nil {
		return err
	}
	addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, portstr(port)))
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
		trace("send", b.String(), addr, ts)
	}
	_, err = tp.Sock.WriteTo(b.Bytes(), addr)
	if err != nil {
		return err
	}
	return nil
}
