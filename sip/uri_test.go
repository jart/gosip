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
	"errors"
	"reflect"
	"testing"

	"github.com/cresta/gosip/sip"
)

type uriTest struct {
	s          string
	e          error
	uri        *sip.URI
	skipFormat bool
}

var uriTests = []uriTest{

	{
		s: "",
		e: errors.New("Incomplete URI: "),
	},

	{
		s: "sip:",
		e: errors.New("Incomplete URI: sip:"),
	},

	{
		s: "sip:example.com:LOL",
		e: errors.New("Error in URI at pos 16: sip:example.com:LOL"),
	},

	{
		s: "sip:example.com",
		uri: &sip.URI{
			Scheme: "sip",
			Host:   "example.com",
		},
	},

	{
		s: "sip:example.com:5060",
		uri: &sip.URI{
			Scheme: "sip",
			Host:   "example.com",
			Port:   5060,
		},
	},

	{
		s: "sips:jart@google.com",
		uri: &sip.URI{
			Scheme: "sips",
			User:   "jart",
			Host:   "google.com",
		},
	},

	{
		s: "sips:jart@google.com:5060",
		uri: &sip.URI{
			Scheme: "sips",
			User:   "jart",
			Host:   "google.com",
			Port:   5060,
		},
	},

	{
		s: "sips:jart:letmein@google.com",
		uri: &sip.URI{
			Scheme: "sips",
			User:   "jart",
			Pass:   "letmein",
			Host:   "google.com",
		},
	},

	{
		s: "sips:jart:LetMeIn@google.com:5060",
		uri: &sip.URI{
			Scheme: "sips",
			User:   "jart",
			Pass:   "LetMeIn",
			Host:   "google.com",
			Port:   5060,
		},
	},

	{
		s: "sips:GOOGLE.com",
		uri: &sip.URI{
			Scheme: "sips",
			Host:   "google.com",
		},
		skipFormat: true,
	},

	{
		s: "sip:[dead:beef::666]:5060",
		uri: &sip.URI{
			Scheme: "sip",
			Host:   "dead:beef::666",
			Port:   5060,
		},
	},

	{
		s: "sip:dead:beef::666:5060",
		e: errors.New("Error in URI at pos 9: sip:dead:beef::666:5060"),
	},

	{
		s: "tel:+12126660420",
		uri: &sip.URI{
			Scheme: "tel",
			Host:   "+12126660420",
		},
	},

	{
		s: "sip:bob%20barker:priceisright@[dead:beef::666]:5060;isup-oli=00",
		uri: &sip.URI{
			Scheme: "sip",
			User:   "bob barker",
			Pass:   "priceisright",
			Host:   "dead:beef::666",
			Port:   5060,
			Param:  &sip.URIParam{Name: "isup-oli", Value: "00"},
		},
	},

	{
		s: "sips:google.com ;lol ;h=omg",
		e: errors.New("Error in URI at pos 15: sips:google.com ;lol ;h=omg"),
	},

	{
		s: "SIP:example.com",
		uri: &sip.URI{
			Scheme: "sip",
			Host:   "example.com",
		},
		skipFormat: true,
	},

	{
		s: "sips:alice@atlanta.com?priority=urgent&subject=project%20x",
		uri: &sip.URI{
			Scheme: "sips",
			User:   "alice",
			Host:   "atlanta.com",
			Header: &sip.URIHeader{
				Name:  "subject",
				Value: "project x",
				Next: &sip.URIHeader{
					Name:  "priority",
					Value: "urgent",
				},
			},
		},
	},

	{
		s: "sip:+1-212-555-1212:1234@gateway.com;user=phone",
		uri: &sip.URI{
			Scheme: "sip",
			User:   "+1-212-555-1212",
			Pass:   "1234",
			Host:   "gateway.com",
			Param:  &sip.URIParam{Name: "user", Value: "phone"},
		},
	},

	{
		s: "sip:atlanta.com;method=register?to=alice%40atlanta.com",
		uri: &sip.URI{
			Scheme: "sip",
			Host:   "atlanta.com",
			Param:  &sip.URIParam{Name: "method", Value: "register"},
			Header: &sip.URIHeader{Name: "to", Value: "alice@atlanta.com"},
		},
	},

	{
		s: "sip:alice;day=tuesday@atlanta.com",
		uri: &sip.URI{
			Scheme: "sip",
			User:   "alice;day=tuesday",
			Host:   "atlanta.com",
		},
		skipFormat: true, // TODO(jart): Fix this.
	},
}

func TestParseURI(t *testing.T) {
	for _, test := range uriTests {
		uri, err := sip.ParseURI([]byte(test.s))
		if err != nil {
			if !reflect.DeepEqual(test.e, err) {
				t.Errorf("%s\nWant: %#v\nGot:  %#v", test.s, test.e, err)
			}
		} else {
			if !reflect.DeepEqual(test.uri, uri) {
				t.Errorf("%s\nWant: %#v\nGot:  %#v", test.s, test.uri, uri)
			}
		}
	}
}

func TestFormatURI(t *testing.T) {
	for _, test := range uriTests {
		if test.skipFormat || test.e != nil {
			continue
		}
		uri := test.uri.String()
		if test.s != uri {
			t.Errorf("\n%s !=\n%s", test.s, uri)
		}
	}
}
