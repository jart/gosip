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

	// Channel to which received SIP messages and errors are published.
	C <-chan *Msg
	E <-chan error

	// Underlying UDP socket.
	Sock *net.UDPConn

	// When you send an outbound request (not a response) you have to set the via
	// tag: ``msg.Via = tp.Via.Copy().Branch().SetNext(msg.Via)``. The details of
	// the branch parameter... are tricky.
	Via *Via

	// Contact that gets put in outbound SIP messages.
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
	c := make(chan *Msg, 32)
	e := make(chan error, 1)
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
	msg, host, port, err := RouteMessage(tp.Via, tp.Contact, msg)
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
