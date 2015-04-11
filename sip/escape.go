package sip

import (
	"bytes"
)

const (
	hexChars = "0123456789abcdef"
)

// appendEscaped appends while URL encoding bytes that don't match the predicate..
func appendEscaped(b *bytes.Buffer, s []byte, p func(byte) bool) {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if p(c) {
			b.WriteByte(c)
		} else {
			b.WriteByte('%')
			b.WriteByte(hexChars[c>>4])
			b.WriteByte(hexChars[c%16])
		}
	}
}
