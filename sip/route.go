package sip

import (
	"errors"
	"log"
	"net"
)

type AddressRoute struct {
	Address string
	Next    *AddressRoute
}

func RouteMessage(via *Via, contact *Addr, old *Msg) (msg *Msg, host string, port uint16, err error) {
	msg = new(Msg)
	*msg = *old // Start off with a shallow copy.
	if msg.Contact == nil {
		msg.Contact = contact
	}
	if msg.IsResponse {
		if via.CompareHostPort(msg.Via) {
			msg.Via = msg.Via.Next
		}
		host, port = msg.Via.Host, msg.Via.Port
		if received, ok := msg.Via.Params["received"]; ok {
			host = received
		}
	} else {
		if contact.CompareHostPort(msg.Route) {
			msg.Route = msg.Route.Next
		}
		if msg.Route != nil {
			if msg.Method == "REGISTER" {
				return nil, "", 0, errors.New("Don't route REGISTER requests")
			}
			if msg.Route.Uri.Params.Has("lr") {
				// RFC3261 16.12.1.1 Basic SIP Trapezoid
				host, port = msg.Route.Uri.Host, msg.Route.Uri.Port
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
	return
}

func RouteAddress(host string, port uint16) (routes *AddressRoute, err error) {
	if net.ParseIP(host) != nil {
		if port == 0 {
			port = 5060
		}
		return &AddressRoute{Address: net.JoinHostPort(host, portstr(port))}, nil
	}
	if port == 0 {
		_, srvs, err := net.LookupSRV("sip", "udp", host)
		if err == nil && len(srvs) > 0 {
			for i := len(srvs) - 1; i >= 0; i-- {
				routes = &AddressRoute{
					Address: net.JoinHostPort(srvs[i].Target, portstr(srvs[i].Port)),
					Next:    routes,
				}
			}
			return routes, nil
		}
		log.Println("net.LookupSRV(sip, udp, %s) failed: %s", err)
		port = 5060
	}
	addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, portstr(port)))
	if err != nil {
		return nil, err
	}
	return &AddressRoute{Address: addr.String()}, nil
}
