// SIP Transport Layer.  Responsible for serializing messages to/from
// your network.

package sip

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/jart/gosip/util"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

var (
	tracing          = flag.Bool("tracing", false, "Enable SIP message tracing")
	timestampTagging = flag.Bool("timestampTagging", false, "Add microsecond timestamps to Via tags")
)

// Transport defines any object capable of sending and receiving SIP messages.
// Such objects are responsible for their own reliability. This means checking
// network errors, raising alarms, rebinding sockets, etc.
type Transport interface {
	// Sends a SIP message. Will not modify msg.
	Send(msg *Msg) error

	// Receives a SIP message. Must only be called by one goroutine. The received
	// address is injected into the first Via header as the "received" param.
	Recv() (msg *Msg, err error)

	// Closes underlying resources. Please make sure all calls using this
	// transport complete first.
	Close() error

	// When you send an outbound request (not a response) you have to set the via
	// tag: ``msg.Via = tport.Via().SetBranch().SetNext(msg.Via)``. The details
	// of the branch parameter... are tricky.
	Via() *Via

	// Returns a linked list of all canonical names and/or IP addresses that may
	// be used to contact *this specific* transport.
	Contact() *Addr
}

// Transport implementation that serializes messages to/from a UDP socket.
type udpTransport struct {
	sock    *net.UDPConn // thing returned by ListenUDP
	addr    *net.UDPAddr // handy for getting ip (contact might be host)
	buf     []byte       // reusable memory for serialization
	via     *Via         // who are we?
	contact *Addr        // uri that points to this specific transport
}

// Creates a new stateless network mechanism for transmitting and receiving SIP
// signalling messages.
//
// 'contact' is a SIP address, e.g. "<sip:1.2.3.4>", that tells how to bind
// sockets. This value is also used for contact headers which tell other
// user-agents where to send responses and hence should only contain an IP or
// canonical address.
func NewUDPTransport(contact *Addr) (tp Transport, err error) {
	saddr := util.HostPortToString(contact.Uri.Host, contact.Uri.Port)
	sock, err := net.ListenPacket("udp", saddr)
	if err != nil {
		return nil, err
	}
	addr := sock.LocalAddr().(*net.UDPAddr)
	contact = contact.Copy()
	contact.Uri.Port = uint16(addr.Port)
	contact.Uri.Params["transport"] = addr.Network()
	return &udpTransport{
		sock:    sock.(*net.UDPConn),
		addr:    addr,
		buf:     make([]byte, 2048),
		contact: contact,
		via: &Via{
			Version: "2.0",
			Proto:   strings.ToUpper(addr.Network()),
			Host:    contact.Uri.Host,
			Port:    contact.Uri.Port,
		},
	}, nil
}

func (tp *udpTransport) Send(msg *Msg) error {
	msg, saddr, err := tp.route(msg)
	if err != nil {
		return err
	}
	addr, err := net.ResolveUDPAddr("ip", saddr)
	if err != nil {
		return errors.New(fmt.Sprintf(
			"udpTransport(%s) failed to resolve %s: %s", tp.addr, saddr, err))
	}
	msg.MaxForwards--
	ts := time.Now()
	addTimestamp(msg, ts)
	var b bytes.Buffer
	msg.Append(&b)
	if *tracing {
		tp.trace("send", b.String(), addr, ts)
	}
	_, err = tp.sock.WriteTo(b.Bytes(), addr)
	if err != nil {
		return errors.New(fmt.Sprintf(
			"udpTransport(%s) write failed: %s", tp.addr, err))
	}
	return nil
}

func (tp *udpTransport) Recv() (msg *Msg, err error) {
	for {
		amt, addr, err := tp.sock.ReadFromUDP(tp.buf)
		if err != nil {
			return nil, errors.New(fmt.Sprintf(
				"udpTransport(%s) read failed: %s", tp.addr, err))
		}
		ts := time.Now()
		packet := string(tp.buf[0:amt])
		if *tracing {
			tp.trace("recv", packet, addr, ts)
		}
		// Validation: http://tools.ietf.org/html/rfc3261#section-16.3
		msg, err = ParseMsg(packet)
		if err != nil {
			log.Printf("udpTransport(%s) got bad message from %s: %s\n%s", tp.addr, addr, err, packet)
			continue
		}
		if msg.Via.Host != addr.IP.String() || int(msg.Via.Port) != addr.Port {
			msg.Via.Params["received"] = addr.String()
		}
		addTimestamp(msg, ts)
		if !tp.sanityCheck(msg) {
			continue
		}
		tp.preprocess(msg)
		break
	}
	return msg, nil
}

func (tp *udpTransport) Via() *Via {
	return tp.via
}

func (tp *udpTransport) Contact() *Addr {
	return tp.contact
}

func (tp *udpTransport) Close() error {
	return tp.sock.Close()
}

func (tp *udpTransport) trace(dir, pkt string, addr net.Addr, t time.Time) {
	size := len(pkt)
	bar := strings.Repeat("-", 72)
	suffix := "\n  "
	if pkt[len(pkt)-1] == '\n' {
		suffix = ""
	}
	log.Printf(
		"%s %d bytes to %s/%s at %s\n"+
			"%s\n"+
			"%s%s"+
			"%s\n",
		dir, size, addr.Network(), addr.String(),
		t.Format(time.RFC3339Nano),
		bar,
		strings.Replace(pkt, "\n", "\n  ", -1), suffix,
		bar)
}

// Test if this message is acceptable.
func (tp *udpTransport) sanityCheck(msg *Msg) bool {
	if msg.MaxForwards <= 0 {
		log.Printf("udpTransport(%s) froot loop detected\n%s", tp.addr, msg)
		go tp.Send(NewResponse(msg, 483))
		return false
	}
	if msg.IsResponse {
		if msg.Status >= 700 {
			log.Printf("udpTransport(%s) msg has crazy status number\n%s", tp.addr, msg)
			go tp.Send(NewResponse(msg, 400))
			return false
		}
	} else {
		if msg.CSeqMethod == "" || msg.CSeqMethod != msg.Method {
			log.Printf("udpTransport(%s) bad cseq number\n%s", tp.addr, msg)
			go tp.Send(NewResponse(msg, 400))
			return false
		}
	}
	return true
}

// Perform some ingress message mangling.
func (tp *udpTransport) preprocess(msg *Msg) {
	if tp.contact.Compare(msg.Route) {
		log.Printf("udpTransport(%s) removing our route header: %s", tp.addr, msg.Route)
		msg.Route = msg.Route.Next
	}
	if _, ok := msg.Request.Params["lr"]; ok && msg.Route != nil && tp.contact.Uri.Compare(msg.Request) {
		// RFC3261 16.4 Route Information Preprocessing
		// RFC3261 16.12.1.2: Traversing a Strict-Routing Proxy
		var oldReq, newReq *URI
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
		log.Printf("udpTransport(%s) fixing request uri after strict router traversal: %s -> %s",
			tp.addr, oldReq, newReq)
	}
}

func (tp *udpTransport) route(old *Msg) (msg *Msg, saddr string, err error) {
	var host string
	var port uint16
	msg = new(Msg)
	*msg = *old // Shallow copy is sufficient.
	if msg.IsResponse {
		msg.Via = old.Via.Copy()
		if msg.Via.CompareAddr(tp.via) {
			// In proxy scenarios we have to remove our own Via.
			msg.Via = msg.Via.Next
		}
		if msg.Via == nil {
			return nil, "", errors.New("Ran out of Via headers when forwarding Response!")
		}
		if msg.Via != nil {
			if received, ok := msg.Via.Params["received"]; ok {
				return msg, received, nil
			} else {
				host, port = msg.Via.Host, msg.Via.Port
			}
		} else {
			return nil, "", errors.New("Message missing Via header")
		}
	} else {
		if msg.Request == nil {
			return nil, "", errors.New("Missing request URI")
		}
		if !msg.Via.CompareAddr(tp.via) {
			return nil, "", errors.New("You forgot to say: msg.Via = tp.Via(msg.Via)")
		}
		if msg.Route != nil {
			if msg.Method == "REGISTER" {
				return nil, "", errors.New("Don't loose route register requests")
			}
			if _, ok := msg.Route.Uri.Params["lr"]; ok {
				// RFC3261 16.12.1.1 Basic SIP Trapezoid
				route := msg.Route
				msg.Route = msg.Route.Next
				host, port = route.Uri.Host, route.Uri.Port
			} else {
				// RFC3261 16.12.1.2: Traversing a Strict-Routing Proxy
				msg.Route = old.Route.Copy()
				msg.Route.Last().Next = &Addr{Uri: msg.Request}
				msg.Request = msg.Route.Uri
				msg.Route = msg.Route.Next
				host, port = msg.Request.Host, msg.Request.Port
			}
		} else {
			host, port = msg.Request.Host, msg.Request.Port
		}
	}
	if msg.OutboundProxy != "" {
		saddr = msg.OutboundProxy
	} else {
		saddr = util.HostPortToString(host, port)
	}
	return
}

func addTimestamp(msg *Msg, ts time.Time) {
	if *timestampTagging {
		msg.Via.Params["µsi"] = strconv.FormatInt(ts.UnixNano()/int64(time.Microsecond), 10)
	}
}
