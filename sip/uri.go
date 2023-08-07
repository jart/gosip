// Copyright 2020 Justine Alexandra Roberts Tunney
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

	"github.com/cresta/gosip/util"
)

const (
	delims = ":/@;?#<>"
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
	var b bytes.Buffer
	uri.Append(&b)
	return b.String()
}

func (uri *URI) Append(b *bytes.Buffer) {
	if uri == nil {
		return
	}
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
		b.WriteString(util.Portstr(uri.Port))
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
