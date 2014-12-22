package sip

import (
	"github.com/jart/gosip/util"
	"strconv"
	"strings"
)

func extractHostPort(s string) (s2, host string, port uint16, err error) {
	port = 5060
	if s == "" {
		err = URIMissingHost
	} else {
		if s[0] == '[' { // quoted/ipv6: sip:[dead:beef::666]:5060
			n := strings.Index(s, "]")
			if n < 0 {
				err = URIMissingHost
			}
			host, s = s[1:n], s[n+1:]
			if s != "" && s[0] == ':' { // we has a port too
				s = s[1:]
				s, port, err = extractPort(s)
			}
		} else { // non-quoted host: sip:1.2.3.4:5060
			switch n := strings.IndexAny(s, delims); {
			case n < 0:
				host, s = s, ""
			case s[n] == ':': // host:port
				host, s = s[0:n], s[n+1:]
				s, port, err = extractPort(s)
			default:
				host, s = s[0:n], s[n:]
			}
		}
	}
	return s, host, port, err
}

func parseParams(s string) (res Params) {
	items := strings.Split(s, ";")
	if items == nil || len(items) == 0 || items[0] == "" {
		return
	}
	res = make(Params, len(items))
	for _, item := range items {
		if item == "" {
			continue
		}
		n := strings.Index(item, "=")
		var k, v string
		if n > 0 {
			k, v = item[0:n], item[n+1:]
		} else {
			k, v = item, ""
		}
		k, kerr := util.URLUnescape(k, false)
		v, verr := util.URLUnescape(v, false)
		if kerr != nil || verr != nil {
			continue
		}
		res[k] = v
	}
	return res
}

func extractPort(s string) (s2 string, port uint16, err error) {
	if n := strings.IndexAny(s, delims); n > 0 {
		port, err = parsePort(s[0:n])
		s = s[n:]
	} else {
		port, err = parsePort(s)
		s = ""
	}
	return s, port, err
}

func parsePort(s string) (port uint16, err error) {
	i, err := strconv.ParseUint(s, 10, 16)
	port = uint16(i)
	return
}
