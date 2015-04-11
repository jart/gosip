// Addr / Via Parameter Library

package sip

import (
	"bytes"
	"strings"
)

// Param is a linked list of ;key="values" for Addr/Via parameters.
type Param struct {
	Name  string
	Value string
	Next  *Param
}

// Get returns an entry in O(n) time.
func (p *Param) Get(name string) *Param {
	if p == nil {
		return nil
	}
	if strings.EqualFold(p.Name, name) {
		return p
	}
	return p.Next.Get(name)
}

// Append serializes parameters in insertion order.
func (p *Param) Append(b *bytes.Buffer) {
	if p == nil {
		return
	}
	p.Next.Append(b)
	b.WriteByte(';')
	appendSanitized(b, []byte(p.Name), tokenc)
	if p.Value != "" {
		b.WriteByte('=')
		appendQuoted(b, []byte(p.Value))
	}
}
