package sip

import (
	"errors"
	"net"
)

func RouteMessage(via *Via, contact *Addr, old *Msg) (msg *Msg, dest string, err error) {
	var host string
	var port uint16
	msg = new(Msg)
	*msg = *old // Start off with a shallow copy.
	if msg.Contact == nil {
		msg.Contact = contact
	}
	if msg.IsResponse {
		if via.CompareHostPort(msg.Via) {
			msg.Via = msg.Via.Next
		}
		if received, ok := msg.Via.Params["received"]; ok {
			return msg, received, nil
		} else {
			host, port = msg.Via.Host, msg.Via.Port
		}
	} else {
		if contact.CompareHostPort(msg.Route) {
			msg.Route = msg.Route.Next
		}
		if msg.Route != nil {
			if msg.Method == "REGISTER" {
				return nil, "", errors.New("Don't route REGISTER requests")
			}
			if msg.Route.Uri.Params.Has("lr") {
				// RFC3261 16.12.1.1 Basic SIP Trapezoid
				host, port = msg.Route.Uri.Host, msg.Route.Uri.GetPort()
			} else {
				// RFC3261 16.12.1.2: Traversing a Strict-Routing Proxy
				msg.Route = old.Route.Copy()
				msg.Route.Last().Next = &Addr{Uri: msg.Request}
				msg.Request = msg.Route.Uri
				msg.Route = msg.Route.Next
				host, port = msg.Request.Host, msg.Request.GetPort()
			}
		} else {
			host, port = msg.Request.Host, msg.Request.GetPort()
		}
	}
	dest = net.JoinHostPort(host, portstr(port))
	return
}
