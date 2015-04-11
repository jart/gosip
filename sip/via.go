// SIP Via Address Library

package sip

import (
	"bytes"
	"github.com/jart/gosip/util"
	"strconv"
)

// Example: SIP/2.0/UDP 1.2.3.4:5060;branch=z9hG4bK556f77e6.
type Via struct {
	Protocol  string // should be "SIP"
	Version   string // protocol version e.g. "2.0"
	Transport string // transport type "UDP"
	Host      string // name or ip of egress interface
	Port      uint16 // network port number
	Param     *Param // param like branch, received, rport, etc.
	Next      *Via   // pointer to next via header if any
}

func (via *Via) Append(b *bytes.Buffer) {
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
	via.Param.Append(b)
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
	res.Param = via.Param
	res.Next = via.Next.Copy()
	return res
}

// Branch adds a randomly generated branch ID.
func (via *Via) Branch() *Via {
	via.Param = &Param{"branch", util.GenerateBranch(), via.Param}
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
		if b1 := via.Param.Get("branch"); b1 != nil {
			if b2 := other.Param.Get("branch"); b2 != nil {
				if b1.Value == b2.Value {
					return true
				}
			}
		}
	}
	return false
}
