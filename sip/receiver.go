package sip

import (
	"log"
	"net"
	"strconv"
	"time"
)

func ReceiveMessages(sock *net.UDPConn, c chan<- *Msg, e chan<- error) {
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
		msg, err := ParseMsg(packet)
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

func addReceived(msg *Msg, addr *net.UDPAddr) {
	if msg.IsResponse() {
		return
	}
	if int(msg.Via.Port) != addr.Port {
		if msg.Via.Param.Get("rport") == nil {
			msg.Via.Param = &Param{"rport", string(addr.Port), msg.Via.Param}
		}
	}
	if msg.Via.Host != addr.IP.String() {
		if msg.Via.Param.Get("received") == nil {
			msg.Via.Param = &Param{"received", addr.IP.String(), msg.Via.Param}
		}
	}
}

func addTimestamp(msg *Msg, ts time.Time) {
	if *timestampTagging {
		msg.Via.Param = &Param{"usi", strconv.FormatInt(ts.UnixNano()/int64(time.Microsecond), 10), msg.Via.Param}
	}
}

// RFC3261 16.4 Route Information Preprocessing
// RFC3261 16.12.1.2: Traversing a Strict-Routing Proxy
func fixMessagesFromStrictRouters(lhost string, lport uint16, msg *Msg) {
	if msg.Request != nil &&
		msg.Request.Param.Get("lr") != nil &&
		msg.Route != nil &&
		msg.Request.Host == lhost &&
		or5060(msg.Request.Port) == lport {
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
		log.Printf("Fixing request URI after strict router traversal: %s -> %s\r\n", oldReq, newReq)
	}
}
