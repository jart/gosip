// SIP Via Address Library

package sip

import (
	"bytes"
	"errors"
	"github.com/jart/gosip/util"
	"strconv"
)

var (
	ViaBadHeader  = errors.New("Bad Via header")
	ViaProtoBlank = errors.New("Via Transport blank")
)

// Example: SIP/2.0/UDP 1.2.3.4:5060;branch=z9hG4bK556f77e6.
type Via struct {
	Protocol  string // should be "SIP"
	Version   string // protocol version e.g. "2.0"
	Transport string // transport type "UDP"
	Host      string // name or ip of egress interface
	Port      uint16 // network port number
	Params    Params // params like branch, received, rport, etc.
	Next      *Via   // pointer to next via header if any
}

func (via *Via) Append(b *bytes.Buffer) error {
	if via.Host == "" {
		return ViaProtoBlank
	}
	if via.Protocol == "" {
		b.WriteString("SIP/")
	} else {
		b.WriteString(via.Protocol)
		b.WriteString("/")
	}
	if via.Version == "" {
		b.WriteString("2.0/")
	} else {
		b.WriteString(via.Version)
		b.WriteString("/")
	}
	if via.Transport == "" {
		b.WriteString("UDP ")
	} else {
		b.WriteString(via.Transport)
		b.WriteString(" ")
	}
	b.WriteString(via.Host)
	if via.Port != 5060 {
		b.WriteString(":")
		b.WriteString(strconv.Itoa(int(via.Port)))
	}
	via.Params.Append(b)
	return nil
}

// Copy returns a deep copy of via.
func (via *Via) Copy() *Via {
	if via == nil {
		return nil
	}
	res := new(Via)
	res.Protocol = via.Protocol
	res.Version = via.Version
	res.Transport = via.Transport
	res.Host = via.Host
	res.Port = via.Port
	res.Params = via.Params.Copy()
	res.Next = via.Next.Copy()
	return res
}

// Branch mutates via with a newly generated branch ID.
func (via *Via) Branch() *Via {
	via.Params["branch"] = util.GenerateBranch()
	return via
}

// Detach returns a shallow copy of via with Next set to nil.
func (via *Via) Detach() *Via {
	res := new(Via)
	*res = *via
	res.Next = nil
	return res
}

// Last returns pointer to last via in linked list.
func (via *Via) Last() *Via {
	if via != nil {
		for ; via.Next != nil; via = via.Next {
		}
	}
	return via
}

func (via *Via) CompareHostPort(other *Via) bool {
	if via != nil && other != nil {
		if via.Host == other.Host &&
			via.Port == other.Port {
			return true
		}
	}
	return false
}

func (via *Via) CompareBranch(other *Via) bool {
	if via != nil && other != nil {
		if b1, ok := via.Params["branch"]; ok {
			if b2, ok := other.Params["branch"]; ok {
				if b1 == b2 {
					return true
				}
			}
		}
	}
	return false
}
