package sip

import (
	"errors"
	"github.com/jart/gosip/util"
	"log"
	"net"
)

type AddressRoute struct {
	Address string
	Next    *AddressRoute
}

func PopulateMessage(via *Via, contact *Addr, msg *Msg) {
	if !msg.IsResponse() {
		if msg.Via == nil {
			msg.Via = via
		}
		if msg.Contact == nil {
			msg.Contact = contact
		}
		if msg.To == nil {
			msg.To = &Addr{Uri: msg.Request}
		}
		if msg.From == nil {
			msg.From = msg.Contact.Copy()
			msg.From.Uri.Params = nil
		}
		if msg.CallID == "" {
			msg.CallID = util.GenerateCallID()
		}
		if msg.CSeq == 0 {
			msg.CSeq = util.GenerateCSeq()
		}
		if msg.CSeqMethod == "" {
			msg.CSeqMethod = msg.Method
		}
		if msg.MaxForwards == 0 {
			msg.MaxForwards = 70
		}
		if msg.UserAgent == "" {
			msg.UserAgent = GosipUA
		}
		if _, ok := msg.Via.Params["branch"]; !ok {
			msg.Via = msg.Via.Copy()
			msg.Via.Params["branch"] = util.GenerateBranch()
		}
		if _, ok := msg.From.Params["tag"]; !ok {
			msg.From = msg.From.Copy()
			msg.From.Params["tag"] = util.GenerateTag()
		}
	}
}

func RouteMessage(via *Via, contact *Addr, msg *Msg) (host string, port uint16, err error) {
	if msg.IsResponse() {
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
				return "", 0, errors.New("Don't route REGISTER requests")
			}
			if msg.Route.Uri.Params.Has("lr") {
				// RFC3261 16.12.1.1 Basic SIP Trapezoid
				host, port = msg.Route.Uri.Host, msg.Route.Uri.Port
			} else {
				// RFC3261 16.12.1.2: Traversing a Strict-Routing Proxy
				msg.Route = msg.Route.Copy()
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

func RouteAddress(host string, port uint16, wantSRV bool) (routes *AddressRoute, err error) {
	if net.ParseIP(host) != nil {
		if port == 0 {
			port = 5060
		}
		return &AddressRoute{Address: net.JoinHostPort(host, portstr(port))}, nil
	}
	if port == 0 {
		if wantSRV {
			_, srvs, err := net.LookupSRV("sip", "udp", host)
			if err == nil && len(srvs) > 0 {
				s := ""
				for i := len(srvs) - 1; i >= 0; i-- {
					routes = &AddressRoute{
						Address: net.JoinHostPort(srvs[i].Target, portstr(srvs[i].Port)),
						Next:    routes,
					}
					s = " " + routes.Address + s
				}
				log.Printf("%s routes to: %s", host, s)
				return routes, nil
			}
			log.Println("net.LookupSRV(sip, udp, %s) failed: %s", err)
		}
		port = 5060
	}
	addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, portstr(port)))
	if err != nil {
		return nil, err
	}
	return &AddressRoute{Address: addr.String()}, nil
}
