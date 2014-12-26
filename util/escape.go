// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// this code adapted from go/src/pkg/http/url.go for gosip because the
// darn public interface won't let me turn off doPlus :'(

package util

import (
	"strconv"
)

type URLEscapeError string

func (e URLEscapeError) Error() string {
	return "invalid URL escape " + strconv.Quote(string(e))
}

// Escape quoted stuff in SIP addresses.
func EscapeDisplay(s string) (res string) {
	for _, c := range s {
		switch c {
		case '"':
			res += "\\\""
		case '\\':
			res += "\\\\"
		default:
			res += string(c)
		}
	}
	return
}

// Return true if the specified character should be escaped when
// appearing in a URL string, according to RFC 2396.
func shouldEscape(c byte) bool {
	if c <= ' ' || c >= 0x7F {
		return true
	}
	switch c {
	case '<', '>', '#', '%', '"', // RFC 2396 delims
		'{', '}', '|', '\\', '^', '[', ']', '`', // RFC2396 unwise
		'?', '&', '=': // RFC 2396 reserved in path
		return true
	}
	return false
}

// urlUnescape is like URLUnescape but can be told not to
// convert + into space.  URLUnescape implements what is
// called "URL encoding" but that only applies to query strings.
// Elsewhere in the URL, + does not mean space.
func URLUnescape(s string, doPlus bool) (string, error) {
	// Count %, check that they're well-formed.
	n := 0
	hasPlus := false
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			n++
			if i+2 >= len(s) || !ishex(s[i+1]) || !ishex(s[i+2]) {
				s = s[i:]
				if len(s) > 3 {
					s = s[0:3]
				}
				return "", URLEscapeError(s)
			}
			i += 3
		case '+':
			hasPlus = doPlus
			i++
		default:
			i++
		}
	}

	if n == 0 && !hasPlus {
		return s, nil
	}

	t := make([]byte, len(s)-2*n)
	j := 0
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			t[j] = unhex(s[i+1])<<4 | unhex(s[i+2])
			j++
			i += 3
		case '+':
			if doPlus {
				t[j] = ' '
			} else {
				t[j] = '+'
			}
			j++
			i++
		default:
			t[j] = s[i]
			j++
			i++
		}
	}
	return string(t), nil
}

func URLEscape(s string, doPlus bool) string {
	spaceCount, hexCount := 0, 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c) {
			if c == ' ' && doPlus {
				spaceCount++
			} else {
				hexCount++
			}
		}
	}

	if spaceCount == 0 && hexCount == 0 {
		return s
	}

	t := make([]byte, len(s)+2*hexCount)
	j := 0
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case c == ' ' && doPlus:
			t[j] = '+'
			j++
		case shouldEscape(c):
			t[j] = '%'
			t[j+1] = "0123456789abcdef"[c>>4]
			t[j+2] = "0123456789abcdef"[c&15]
			j += 3
		default:
			t[j] = s[i]
			j++
		}
	}
	return string(t)
}

func ishex(c byte) bool {
	switch {
	case '0' <= c && c <= '9':
		return true
	case 'a' <= c && c <= 'f':
		return true
	case 'A' <= c && c <= 'F':
		return true
	}
	return false
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
