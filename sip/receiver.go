package sip

import (
	"github.com/jart/gosip/util"
	"log"
	"net"
	"strconv"
	"time"
)

func ReceiveMessages(contact *Addr, sock *net.UDPConn, c chan *Msg, e chan error) {
	buf := make([]byte, 2048)
	for {
		amt, addr, err := sock.ReadFromUDP(buf)
		if err != nil {
			if !util.IsUseOfClosed(err) {
				e <- err
			}
			return
		}
		ts := time.Now()
		packet := string(buf[0:amt])
		if *tracing {
			trace("recv", packet, addr, ts)
		}
		msg, err := ParseMsg(packet)
		if err != nil {
			log.Println("Dropping SIP message:", err)
			continue
		}
		addReceived(msg, addr)
		addTimestamp(msg, ts)
		if contact.CompareHostPort(msg.Route) {
			msg.Route = msg.Route.Next
		}
		fixMessagesFromStrictRouters(contact, msg)
		c <- msg
	}
}

func addReceived(msg *Msg, addr *net.UDPAddr) {
	if msg.Via.Host != addr.IP.String() || int(msg.Via.Port) != addr.Port {
		msg.Via.Params["received"] = addr.String()
	}
}

func addTimestamp(msg *Msg, ts time.Time) {
	if *timestampTagging {
		msg.Via.Params["µsi"] = strconv.FormatInt(ts.UnixNano()/int64(time.Microsecond), 10)
	}
}

// RFC3261 16.4 Route Information Preprocessing
// RFC3261 16.12.1.2: Traversing a Strict-Routing Proxy
func fixMessagesFromStrictRouters(contacts *Addr, msg *Msg) {
	if msg.Request != nil && msg.Request.Params.Has("lr") && msg.Route != nil && contacts.Uri.CompareHostPort(msg.Request) {
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
		log.Printf("Fixing request URI after strict router traversal: %s -> %s", oldReq, newReq)
	}
}
