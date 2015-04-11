// URI Header Library

package sip

import (
	"bytes"
	"strings"
)

// URIHeader is a linked list of ?key=values for URI headers.
type URIHeader struct {
	Name  string
	Value string
	Next  *URIHeader
}

// Get returns an entry in O(n) time.
func (p *URIHeader) Get(name string) *URIHeader {
	if p == nil {
		return nil
	}
	if strings.EqualFold(p.Name, name) {
		return p
	}
	return p.Next.Get(name)
}

// Append serializes URI headers in insertion order.
func (p *URIHeader) Append(b *bytes.Buffer) {
	if p == nil {
		return
	}
	if p.Next != nil {
		p.Next.Append(b)
		b.WriteByte('&')
	} else {
		b.WriteByte('?')
	}
	appendEscaped(b, []byte(p.Name), paramc)
	if p.Value != "" {
		b.WriteByte('=')
		appendEscaped(b, []byte(p.Value), paramc)
	}
}
