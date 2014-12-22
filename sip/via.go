// SIP Via Address Library

package sip

import (
	"bytes"
	"errors"
	"github.com/jart/gosip/util"
	"strconv"
	"strings"
)

var (
	ViaBadHeader  = errors.New("bad via header")
	ViaProtoBlank = errors.New("via.Proto blank")
)

// Example: SIP/2.0/UDP 1.2.3.4:5060;branch=z9hG4bK556f77e6.
type Via struct {
	Next    *Via   // pointer to next via header if any
	Version string // protocol version e.g. "2.0"
	Proto   string // transport type "UDP"
	Host    string // name or ip of egress interface
	Port    uint16 // network port number
	Params  Params // params like branch, received, rport, etc.
}

// Parses a single SIP Via header, provided the part that comes after "Via: ".
//
// Via headers are goofy; they're like a URI without a scheme, user, and pass.
// They tell about the network interfaces a packet traversed to reach us.
//
// EBNF: http://sofia-sip.sourceforge.net/refdocs/sip/group__sip__via.html
func ParseVia(s string) (via *Via, err error) {
	if s[0:4] == "SIP/" {
		s = s[4:]
		if n1 := strings.Index(s, "/"); n1 > 0 {
			if n2 := strings.Index(s, " "); n2 >= n1+3 {
				via = new(Via)
				via.Version = s[0:n1]
				via.Proto = s[n1+1 : n2]
				s, via.Host, via.Port, err = extractHostPort(s[n2+1:])
				if err != nil {
					return nil, err
				}
				if s != "" && s[0] == ';' {
					via.Params = parseParams(s[1:])
					s = ""
				}
				return via, nil
			}
		}
	}
	return nil, ViaBadHeader
}

func (via *Via) Append(b *bytes.Buffer) error {
	if via.Version == "" {
		via.Version = "2.0"
	}
	if via.Proto == "" {
		via.Proto = "UDP"
	}
	if via.Port == 0 {
		via.Port = 5060
	}
	if via.Host == "" {
		return ViaProtoBlank
	}
	b.WriteString("SIP/")
	b.WriteString(via.Version)
	b.WriteString("/")
	b.WriteString(via.Proto)
	b.WriteString(" ")
	b.WriteString(via.Host)
	if via.Port != 5060 {
		b.WriteString(":")
		b.WriteString(strconv.Itoa(int(via.Port)))
	}
	via.Params.Append(b)
	return nil
}

// deep copies a new Via object
func (via *Via) Copy() *Via {
	if via == nil {
		return nil
	}
	res := new(Via)
	res.Version = via.Version
	res.Proto = via.Proto
	res.Host = via.Host
	res.Port = via.Port
	res.Params = via.Params.Copy()
	res.Next = via.Next.Copy()
	return res
}

// Sets newly generated branch ID and returns self.
func (via *Via) Branch() *Via {
	via = via.Copy()
	via.Params["branch"] = util.GenerateBranch()
	return via
}

// Sets Next field and returns self.
func (via *Via) SetNext(next *Via) *Via {
	via.Next = next
	return via
}

// returns pointer to last via in linked list.
func (via *Via) Last() *Via {
	if via != nil {
		for ; via.Next != nil; via = via.Next {
		}
	}
	return via
}

// returns true if version, proto, ip, port, and branch match
func (via *Via) Compare(other *Via) bool {
	return (via.CompareAddr(other) && via.CompareBranch(other))
}

func (via *Via) CompareAddr(other *Via) bool {
	if via != nil && other != nil {
		if via.Version == other.Version &&
			via.Proto == other.Proto &&
			via.Host == other.Host &&
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
