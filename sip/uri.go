// SIP URI Library
//
// We can't use net.URL because it doesn't support SIP URIs. This is because:
// a) it doesn't support semicolon parameters; b) it doesn't extract the user
// and host information when the "//" isn't present.
//
// For example:
//
//   jart@example.test;isup-oli=29
//
// Roughly equates to:
//
//   {Scheme: "sip",
//    User: "jart",
//    Pass: "",
//    Host: "example.test",
//    Port: "",
//    Params: {"isup-oli": "29"}}
//

package sip

import (
	"bytes"
	"errors"
	"github.com/jart/gosip/util"
	"sort"
)

const (
	delims = ":/@;?#<>"
)

var (
	URISchemeNotFound = errors.New("scheme not found")
	URIMissingHost    = errors.New("host missing")
	URIBadPort        = errors.New("invalid port number")
)

type URIHeaders map[string]string

type URI struct {
	Scheme  string     // e.g. sip, sips, tel, etc.
	User    string     // e.g. sip:USER@host
	Pass    string     // e.g. sip:user:PASS@host
	Host    string     // e.g. example.com, 1.2.3.4, etc.
	Port    uint16     // e.g. 5060, 80, etc.
	Params  Params     // e.g. ;isup-oli=00;day=tuesday
	Headers URIHeaders // e.g. ?subject=project%20x&lol=cat
}

//go:generate ragel -Z -G2 -o uri_parse.go uri_parse.rl

// Deep copies a URI object.
func (uri *URI) Copy() *URI {
	if uri == nil {
		return nil
	}
	res := new(URI)
	*res = *uri
	res.Params = uri.Params.Copy()
	res.Headers = uri.Headers.Copy()
	return res
}

// TODO(jart): URI Comparison https://tools.ietf.org/html/rfc3261#section-19.1.4

func (uri *URI) String() string {
	if uri == nil {
		return "<nil>"
	}
	var b bytes.Buffer
	uri.Append(&b)
	return b.String()
}

func (uri *URI) Append(b *bytes.Buffer) {
	if uri.Scheme == "" {
		b.WriteString("sip:")
	} else {
		b.WriteString(uri.Scheme)
		b.WriteByte(':')
	}
	if uri.User != "" {
		if uri.Pass != "" {
			appendEscaped(b, []byte(uri.User), userc)
			b.WriteByte(':')
			appendEscaped(b, []byte(uri.Pass), passc)
		} else {
			appendEscaped(b, []byte(uri.User), userc)
		}
		b.WriteByte('@')
	}
	if util.IsIPv6(uri.Host) {
		b.WriteByte('[')
		b.WriteString(uri.Host)
		b.WriteByte(']')
	} else {
		b.WriteString(uri.Host)
	}
	if uri.Port > 0 {
		b.WriteByte(':')
		b.WriteString(portstr(uri.Port))
	}
	uri.Params.Append(b)
	uri.Headers.Append(b)
}

func (uri *URI) CompareHostPort(other *URI) bool {
	if uri != nil && other != nil {
		if uri.Host == other.Host && uri.GetPort() == other.GetPort() {
			return true
		}
	}
	return false
}

func (uri *URI) GetPort() uint16 {
	if uri.Port == 0 {
		if uri.Scheme == "sips" {
			return 5061
		} else {
			return 5060
		}
	} else {
		return uri.Port
	}
}

func (headers URIHeaders) Copy() URIHeaders {
	res := make(URIHeaders, len(headers))
	for k, v := range headers {
		res[k] = v
	}
	return res
}

func (headers URIHeaders) Append(b *bytes.Buffer) {
	if headers != nil && len(headers) > 0 {
		keys := make([]string, len(headers))
		i := 0
		for k, _ := range headers {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		first := true
		for _, k := range keys {
			if first {
				b.WriteByte('?')
				first = false
			} else {
				b.WriteByte('&')
			}
			appendEscaped(b, []byte(k), headerc)
			v := headers[k]
			if v != "" {
				b.WriteByte('=')
				appendEscaped(b, []byte(v), headerc)
			}
		}
	}
}
