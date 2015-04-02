package sip

import (
	"bytes"
)

// tokencify removes all characters that aren't tokenc.
func tokencify(s []byte) []byte {
	t := make([]byte, len(s))
	j := 0
	l := 0
	for i := 0; i < len(s); i++ {
		if tokenc(s[i]) {
			t[j] = s[i]
			j++
			l++
		}
	}
	return t[:l]
}

// quote formats an address parameter value or display name.
//
// Quotation marks are only added if necessary. This implementation will
// truncate on input error.
func quote(s []byte) []byte {
	for i := 0; i < len(s); i++ {
		if !tokenc(s[i]) && s[i] != ' ' {
			return quoteQuoted(s)
		}
	}
	return s
}

// quoteQuoted formats an address parameter value or display name with quotes.
//
// This implementation will truncate on input error.
func quoteQuoted(s []byte) []byte {
	var b bytes.Buffer
	b.WriteByte('"')
	for i := 0; i < len(s); i++ {
		if qdtextc(s[i]) {
			if s[i] == '\r' {
				if i+2 >= len(s) || s[i+1] != '\n' ||
					!(s[i+2] == ' ' || s[i+2] == '\t') {
					break
				}
			}
		} else {
			if !qdtextesc(s[i]) {
				break
			}
			b.WriteByte('\\')
		}
		b.WriteByte(s[i])
	}
	b.WriteByte('"')
	return b.Bytes()
}
