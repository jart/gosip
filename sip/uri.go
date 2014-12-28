// SIP URI Library
//
// We can't use net.URL because it doesn't support SIP URIs. This is because:
// a) it doesn't support semicolon parameters; b) it doesn't extract the user
// and host information when the "//" isn't present.
//
// For example:
//
//   jtunney@bsdtelecom.net;isup-oli=29
//
// Roughly equates to:
//
//   {Scheme: "sip",
//    User: "jtunney",
//    Pass: "",
//    Host: "bsdtelecom.net",
//    Port: "",
//    Params: {"isup-oli": "29"}}
//

package sip

import (
	"bytes"
	"errors"
	"github.com/jart/gosip/util"
	"strings"
)

const (
	delims = ":/@;?#<>"
)

var (
	URIEmpty          = errors.New("empty uri")
	URISchemeNotFound = errors.New("scheme not found")
	URIMissingHost    = errors.New("host missing")
	URIBadPort        = errors.New("invalid port number")
)

type Params map[string]string

type URI struct {
	Scheme string // sip, tel, etc.
	User   string // sip:USER@host
	Pass   string // sip:user:PASS@host
	Host   string // example.com, 1.2.3.4, etc.
	Port   uint16 // 5060, 80, etc.
	Params Params // semicolon delimited params after uris and addrs
}

// Parses a SIP URI.
func ParseURI(s string) (uri *URI, err error) {
	uri = new(URI)
	if s == "" {
		return nil, URIEmpty
	}

	// Extract scheme.
	n := strings.IndexAny(s, delims)
	if n < 0 || s[n] != ':' {
		return nil, URISchemeNotFound
	}
	uri.Scheme, s = s[0:n], s[n+1:]

	// Extract user/pass.
	n = strings.IndexAny(s, delims)
	if n > 0 && s[n] == ':' { // sip:user:pass@host
		// if next token isn't '@' then assume 'sip:host:port'
		s2 := s[n+1:]
		n2 := strings.IndexAny(s2, delims)
		if n2 > 0 && s2[n2] == '@' {
			uri.User = s[0:n]
			s, n = s2, n2
			if n < 0 || s[n] != '@' {
				return nil, URIMissingHost
			}
			uri.Pass, s = s[0:n], s[n+1:]
		}
	} else if n > 0 && s[n] == '@' { // user@host
		uri.User, s = s[0:n], s[n+1:]
	}

	// Extract host/port.
	s, uri.Host, uri.Port, err = extractHostPort(s)
	if err != nil {
		return nil, err
	}

	// Extract semicolon delimited params.
	if s != "" && s[0] == ';' {
		uri.Params = parseParams(s[1:])
		s = ""
	}

	// if s != "" {
	// 	fmt.Fprintf(os.Stderr, "leftover data: %v\n", s)
	// }

	uri.User, err = util.URLUnescape(uri.User, false)
	if err != nil {
		return nil, err
	}
	uri.Pass, err = util.URLUnescape(uri.Pass, false)
	if err != nil {
		return nil, err
	}
	uri.Host, err = util.URLUnescape(uri.Host, false)
	if err != nil {
		return nil, err
	}

	return uri, nil
}

// Deep copies a URI object.
func (uri *URI) Copy() *URI {
	if uri == nil {
		return nil
	}
	res := new(URI)
	*res = *uri
	res.Params = uri.Params.Copy()
	return res
}

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
		uri.Scheme = "sip"
	}
	b.WriteString(uri.Scheme)
	b.WriteString(":")
	if uri.User != "" {
		if uri.Pass != "" {
			b.WriteString(util.URLEscape(uri.User, false))
			b.WriteString(":")
			b.WriteString(util.URLEscape(uri.Pass, false))
		} else {
			b.WriteString(util.URLEscape(uri.User, false))
		}
		b.WriteString("@")
	}
	if util.IsIPv6(uri.Host) {
		b.WriteString("[" + util.URLEscape(uri.Host, false) + "]")
	} else {
		b.WriteString(util.URLEscape(uri.Host, false))
	}
	if uri.Port > 0 {
		b.WriteString(":" + portstr((uri.Port)))
	}
	uri.Params.Append(b)
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

func (params Params) Copy() Params {
	res := make(Params, len(params))
	for k, v := range params {
		res[k] = v
	}
	return res
}

func (params Params) Append(b *bytes.Buffer) {
	if params != nil && len(params) > 0 {
		for k, v := range params {
			b.WriteString(";")
			b.WriteString(util.URLEscape(k, false))
			if v != "" {
				b.WriteString("=")
				b.WriteString(util.URLEscape(v, false))
			}
		}
	}
}

func (params Params) Has(k string) bool {
	_, ok := params["lr"]
	return ok
}
