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
	"github.com/jart/gosip/util"
)

// Represents a SIP Address Linked List
type Addr struct {
	Uri     *URI   // never nil
	Display string // blank if not specified
	Params  Params // these look like ;key=lol;rport;key=wut
	Next    *Addr  // for comma separated lists of addresses
}

//go:generate ragel -Z -G2 -o addr_parse.go addr_parse.rl

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
		appendQuoted(b, []byte(addr.Display))
		b.WriteByte(' ')
	}
	b.WriteByte('<')
	addr.Uri.Append(b)
	b.WriteByte('>')
	addr.Params.AppendQuoted(b)
	if addr.Next != nil {
		b.WriteByte(',')
		b.WriteByte(' ')
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
