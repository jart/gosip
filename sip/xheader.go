// Addr / Via Parameter Library

package sip

import (
	"bytes"
	"strings"
)

// XHeader is a linked list storing an unrecognized SIP headers.
type XHeader struct {
	Name  string // tokenc
	Value []byte // UTF8, never nil
	Next  *XHeader
}

// Get returns an entry in O(n) time.
func (h *XHeader) Get(name string) *XHeader {
	if h == nil {
		return nil
	}
	if strings.EqualFold(h.Name, name) {
		return h
	}
	return h.Next.Get(name)
}

// String turns XHeader into a string.
func (h *XHeader) String() string {
	var b bytes.Buffer
	h.Append(&b)
	return b.String()
}

// Append serializes headers in insertion order.
func (h *XHeader) Append(b *bytes.Buffer) {
	if h == nil {
		return
	}
	h.Next.Append(b)
	appendSanitized(b, []byte(h.Name), tokenc)
	b.WriteString(": ")
	b.Write(h.Value)
	b.WriteString("\r\n")
}
