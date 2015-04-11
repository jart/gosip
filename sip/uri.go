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
)

const (
	delims = ":/@;?#<>"
)

var (
	URISchemeNotFound = errors.New("scheme not found")
	URIMissingHost    = errors.New("host missing")
	URIBadPort        = errors.New("invalid port number")
)

type URI struct {
	Scheme string     // e.g. sip, sips, tel, etc.
	User   string     // e.g. sip:USER@host
	Pass   string     // e.g. sip:user:PASS@host
	Host   string     // e.g. example.com, 1.2.3.4, etc.
	Port   uint16     // e.g. 5060, 80, etc.
	Param  *URIParam  // e.g. ;isup-oli=00;day=tuesday
	Header *URIHeader // e.g. ?subject=project%20x&lol=cat
}

//go:generate ragel -Z -G2 -o uri_parse.go uri_parse.rl

// Deep copies a URI object.
func (uri *URI) Copy() *URI {
	if uri == nil {
		return nil
	}
	res := new(URI)
	*res = *uri
	res.Param = uri.Param
	res.Header = uri.Header
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
	uri.Param.Append(b)
	uri.Header.Append(b)
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
