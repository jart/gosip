// SIP Address Library
//
// For example:
//
//   "J.A. Roberts Tunney" <sip:jtunney@bsdtelecom.net;isup-oli=29>;tag=deadbeef
//
// Roughly equates to:
//
//   {Display: "J.A. Roberts Tunney",
//    Params: {"tag": "deadbeef"},
//    Uri: {Scheme: "sip",
//          User: "jtunney",
//          Pass: "",
//          Host: "bsdtelecom.net",
//          Port: "",
//          Params: {"isup-oli": "29"}}}
//

package sip

import (
	"bytes"
	"errors"
	"github.com/jart/gosip/util"
	"strings"
)

// Represents a SIP Address Linked List
type Addr struct {
	Uri     *URI   // never nil
	Display string // blank if not specified
	Params  Params // these look like ;key=lol;rport;key=wut
	Next    *Addr  // for comma separated lists of addresses
}

// Parses a SIP address.
func ParseAddr(s string) (addr *Addr, err error) {
	addr = new(Addr)
	l := len(s)
	if l == 0 {
		return nil, errors.New("empty addr")
	}

	// Extract display.
	switch n := strings.IndexAny(s, "\"<"); {
	case n < 0:
		return nil, errors.New("Invalid address")
	case s[n] == '<': // Display is not quoted.
		addr.Display, s = strings.Trim(s[0:n], " "), s[n+1:]
	case s[n] == '"': // We found an opening quote.
		s = s[n+1:]
	LOL:
		for s != "" {
			switch s[0] {
			case '"': // Closing quote.
				s = s[1:]
				break LOL
			case '\\': // Escape sequence.
				if len(s) < 2 {
					return nil, errors.New("Evil quote escape")
				}
				switch s[1] {
				case '"':
					addr.Display += "\""
				case '\\':
					addr.Display += "\\"
				}
				s = s[2:]
			default: // Generic character.
				addr.Display += string(s[0])
				s = s[1:]
			}
		}
		if s == "" {
			return nil, errors.New("No closing quote in display")
		}
		for s != "" {
			c := s[0]
			s = s[1:]
			if c == '<' {
				break
			}
		}
	}

	if n := strings.Index(s, ">"); n > 0 {
		addr.Uri, err = ParseURI(s[0:n])
		if err != nil {
			return nil, err
		}
		s = s[n+1:]
	} else {
		addr.Uri, err = ParseURI(s)
		if err != nil {
			return nil, err
		}
		s = ""
	}

	// Extract semicolon delimited params.
	if s != "" && s[0] == ';' {
		addr.Params = parseParams(s[1:])
		s = ""
	}

	// Is there another address?
	s = strings.TrimLeft(s, " \t")
	if s != "" && s[0] == ',' {
		s = strings.TrimLeft(s[1:], " \t")
		if s != "" {
			addr.Next, err = ParseAddr(s)
			if err != nil {
				return nil, err
			}
		}
	}

	return addr, nil
}

func (addr *Addr) String() string {
	if addr == nil {
		return "<nil>"
	}
	var b bytes.Buffer
	addr.Append(&b)
	return b.String()
}

// Returns self if non-nil, otherwise other.
func (addr *Addr) Or(other *Addr) *Addr {
	if addr == nil {
		return other
	}
	return addr
}

// Sets newly generated tag ID and returns self.
func (addr *Addr) Tag() *Addr {
	addr.Params["tag"] = util.GenerateTag()
	return addr
}

// Reassembles a SIP address into a buffer.
func (addr *Addr) Append(b *bytes.Buffer) error {
	if addr.Display != "" {
		b.WriteString("\"")
		b.WriteString(util.EscapeDisplay(addr.Display))
		b.WriteString("\" ")
	}
	b.WriteString("<")
	addr.Uri.Append(b)
	b.WriteString(">")
	addr.Params.Append(b)
	if addr.Next != nil {
		b.WriteString(",")
		addr.Next.Append(b)
	}
	return nil
}

// Deep copies a new Addr object.
func (addr *Addr) Copy() *Addr {
	if addr == nil {
		return nil
	}
	res := new(Addr)
	res.Uri = addr.Uri.Copy()
	res.Params = addr.Params.Copy()
	res.Next = addr.Next.Copy()
	return res
}

func (addr *Addr) CompareHostPort(other *Addr) bool {
	if addr != nil && other != nil {
		if addr.Uri.CompareHostPort(other.Uri) {
			return true
		}
	}
	return false
}

// Returns pointer to last addr in linked list.
func (addr *Addr) Last() *Addr {
	if addr != nil {
		for ; addr.Next != nil; addr = addr.Next {
		}
	}
	return addr
}

// Returns number of items in the linked list.
func (addr *Addr) Len() int {
	count := 0
	for ; addr != nil; addr = addr.Next {
		count++
	}
	return count
}

// Returns self with linked list reversed.
func (addr *Addr) Reversed() *Addr {
	var res *Addr
	for ; addr != nil; addr = addr.Next {
		a := new(Addr)
		*a = *addr
		a.Next = res
		res = a
	}
	return res
}
