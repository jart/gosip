// URI Parameter Library

package sip

import (
	"bytes"
	"strings"
)

// URIParam is a linked list of ;key=values for URI parameters.
type URIParam struct {
	Name  string
	Value string
	Next  *URIParam
}

// Get returns an entry in O(n) time.
func (p *URIParam) Get(name string) *URIParam {
	if p == nil {
		return nil
	}
	if strings.EqualFold(p.Name, name) {
		return p
	}
	return p.Next.Get(name)
}

// Append serializes URI parameters in insertion order.
func (p *URIParam) Append(b *bytes.Buffer) {
	if p == nil {
		return
	}
	p.Next.Append(b)
	b.WriteByte(';')
	appendEscaped(b, []byte(p.Name), paramc)
	if p.Value != "" {
		b.WriteByte('=')
		appendEscaped(b, []byte(p.Value), paramc)
	}
}
