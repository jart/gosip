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

package sip_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/cresta/gosip/sip"
)

type addrTest struct {
	name        string
	s           string
	s_canonical string
	addr        sip.Addr
	err         error
}

var addrTests = []addrTest{

	{
		name: "Basic address",
		s:    "<sip:pokemon.net>",
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   "pokemon.net",
			},
		},
	},

	{
		name: "Address parameter",
		s:    "<sip:pokemon.net>;tag=deadbeef",
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   "pokemon.net",
			},
			Param: &sip.Param{Name: "tag", Value: "deadbeef"},
		},
	},

	{
		name:        "Address parameter spacing",
		s:           "<sip:pokemon.net>\t ;\t tag\t = \tdeadbeef",
		s_canonical: "<sip:pokemon.net>;tag=deadbeef",
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   "pokemon.net",
			},
			Param: &sip.Param{Name: "tag", Value: "deadbeef"},
		},
	},

	{
		name:        "Address parameter quoted",
		s:           "<sip:pokemon.net>;tag=\"deadbeef\"",
		s_canonical: "<sip:pokemon.net>;tag=deadbeef",
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   "pokemon.net",
			},
			Param: &sip.Param{Name: "tag", Value: "deadbeef"},
		},
	},

	{
		name:        "Address parameter quoted spacing",
		s:           "<sip:pokemon.net>\t ;\t tag\t = \t\"deadbeef\"",
		s_canonical: "<sip:pokemon.net>;tag=deadbeef",
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   "pokemon.net",
			},
			Param: &sip.Param{Name: "tag", Value: "deadbeef"},
		},
	},

	{
		name: "Address parameter quoted escaped",
		s:    "<sip:pokemon.net>;tag=\"\\\"deadbeef\\\"\"",
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   "pokemon.net",
			},
			Param: &sip.Param{Name: "tag", Value: "\"deadbeef\""},
		},
	},

	{
		name: "URI parameter",
		s:    "<sip:brave@toaster.net;isup-oli=29>",
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				User:   "brave",
				Host:   "toaster.net",
				Param:  &sip.URIParam{Name: "isup-oli", Value: "29"},
			},
		},
	},

	{
		name: "Address + URI parameter",
		s:    "<sip:brave@toaster.net;isup-oli=29>;tag=deadbeef",
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				User:   "brave",
				Host:   "toaster.net",
				Param:  &sip.URIParam{Name: "isup-oli", Value: "29"},
			},
			Param: &sip.Param{Name: "tag", Value: "deadbeef"},
		},
	},

	{
		s: `<sip:pokemon.com>, Ditto <sip:ditto@pokemon.com>`,
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   "pokemon.com",
			},
			Next: &sip.Addr{
				Display: "Ditto",
				Uri: &sip.URI{
					Scheme: "sip",
					User:   "ditto",
					Host:   "pokemon.com",
				},
			},
		},
	},

	{
		s: `<sip:1.2.3.4>, <sip:1.2.3.5>, <sip:[666::dead:beef]>`,
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   "1.2.3.4",
			},
			Next: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "1.2.3.5",
				},
				Next: &sip.Addr{
					Uri: &sip.URI{
						Scheme: "sip",
						Host:   "666::dead:beef",
					},
				},
			},
		},
	},

	{
		s: "\"\\\"\\\"Justine \\\\Tunney \" " +
			"<sip:jart@google.com;isup-oli=29>;tag=deadbeef",
		addr: sip.Addr{
			Display: "\"\"Justine \\Tunney ",
			Uri: &sip.URI{
				Scheme: "sip",
				User:   "jart",
				Host:   "google.com",
				Param:  &sip.URIParam{Name: "isup-oli", Value: "29"},
			},
			Param: &sip.Param{Name: "tag", Value: "deadbeef"},
		},
	},
}

func ParseAddrBytes(s []byte) (addr *sip.Addr, err error) {
	var b bytes.Buffer
	b.WriteString("SIP/2.0 900 ParseAddrBytes()\r\nContact:")
	b.Write(s)
	b.WriteString("\r\n\r\n")
	msg, err := sip.ParseMsg(b.Bytes())
	if err != nil {
		return nil, err
	}
	return msg.Contact, nil
}

func TestParseAddrBytes(t *testing.T) {
	for _, test := range addrTests {
		addr, err := ParseAddrBytes([]byte(test.s))
		if err != nil {
			if test.err == nil {
				t.Error(err)
				continue
			} else { // Test was supposed to fail.
				t.Fatal("TODO(jart): Implement failing support.")
			}
		}
		if !reflect.DeepEqual(&test.addr, addr) {
			t.Errorf("%s:\n%#v !=\n%#v\n%#v !=\n%#v", test.name, &test.addr, addr, test.addr.Uri, addr.Uri)
		}
	}
}

func TestAddrString(t *testing.T) {
	for _, test := range addrTests {
		addr := test.addr.String()
		if test.s_canonical != "" {
			if test.s_canonical != addr {
				t.Errorf("%s (Canonicalization)\nWant: %s\nGot:  %s", test.name, test.s_canonical, addr)
			}
		} else {
			if test.s != addr {
				t.Errorf("%s\nWant: %s\nGot:  %s", test.name, test.s, addr)
			}
		}
	}
}

func TestReversed(t *testing.T) {
	a := &sip.Addr{
		Uri: &sip.URI{
			Scheme: "sip",
			Host:   "1.2.3.4",
			Port:   5060,
		},
		Next: &sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   "2.3.4.5",
				Port:   5060,
			},
		},
	}
	b := a.Reversed()
	if b.Uri.Host != "2.3.4.5" {
		t.Error("first bad", b.Uri.Host)
	}
	if b.Next.Uri.Host != "1.2.3.4" {
		t.Error("second bad", b.Next.Uri.Host)
	}
	if b.Next.Next != nil {
		t.Error("wtf", b.Next.Next)
	}
}

func addrError(t *testing.T, name string, want, got *sip.Addr) {
	if want != nil && got != nil {
		t.Errorf("%s:\n%#v\n%#v\n!=\n%#v\n%#v", name, want, want.Uri, got, got.Uri)
	} else {
		t.Errorf("%s:\n%#v\n!=\n%#v", name, want, got)
	}
}
