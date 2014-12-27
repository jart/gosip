package sip_test

import (
	"github.com/jart/gosip/sip"
	"reflect"
	"testing"
)

type uriTest struct {
	s   string  // user input we want to convert
	uri sip.URI // what 's' should become after parsing
	err error   // if we expect parsing to fail
}

var uriTests = []uriTest{

	uriTest{
		s: "sip:google.com",
		uri: sip.URI{
			Scheme: "sip",
			Host:   "google.com",
		},
	},

	uriTest{
		s: "sip:jart@google.com",
		uri: sip.URI{
			Scheme: "sip",
			User:   "jart",
			Host:   "google.com",
		},
	},

	uriTest{
		s: "sip:jart@google.com:5060",
		uri: sip.URI{
			Scheme: "sip",
			User:   "jart",
			Host:   "google.com",
			Port:   5060,
		},
	},

	uriTest{
		s: "sip:google.com:666",
		uri: sip.URI{
			Scheme: "sip",
			Host:   "google.com",
			Port:   666,
		},
	},

	uriTest{
		s: "sip:+12125650666@cat.lol",
		uri: sip.URI{
			Scheme: "sip",
			User:   "+12125650666",
			Host:   "cat.lol",
		},
	},

	uriTest{
		s: "sip:jart:lawl@google.com",
		uri: sip.URI{
			Scheme: "sip",
			User:   "jart",
			Pass:   "lawl",
			Host:   "google.com",
		},
	},

	uriTest{
		s: "sip:jart:lawl@google.com;isup-oli=00;omg;lol=cat",
		uri: sip.URI{
			Scheme: "sip",
			User:   "jart",
			Pass:   "lawl",
			Host:   "google.com",
			Params: sip.Params{
				"isup-oli": "00",
				"omg":      "",
				"lol":      "cat",
			},
		},
	},

	uriTest{
		s: "sip:jart@google.com;isup-oli=00;omg;lol=cat",
		uri: sip.URI{
			Scheme: "sip",
			User:   "jart",
			Host:   "google.com",
			Params: sip.Params{
				"isup-oli": "00",
				"omg":      "",
				"lol":      "cat",
			},
		},
	},

	uriTest{
		s: "sip:[dead:beef::666]",
		uri: sip.URI{
			Scheme: "sip",
			Host:   "dead:beef::666",
		},
	},

	uriTest{
		s: "sips:[dead:beef::666]:5060",
		uri: sip.URI{
			Scheme: "sips",
			Host:   "dead:beef::666",
			Port:   5060,
		},
	},

	uriTest{
		s: "sip:lol:cat@[dead:beef::666]:65535",
		uri: sip.URI{
			Scheme: "sip",
			User:   "lol",
			Pass:   "cat",
			Host:   "dead:beef::666",
			Port:   65535,
		},
	},

	uriTest{
		s: "sip:lol:cat@[dead:beef::666]:65535;oh;my;goth",
		uri: sip.URI{
			Scheme: "sip",
			User:   "lol",
			Pass:   "cat",
			Host:   "dead:beef::666",
			Port:   65535,
			Params: sip.Params{
				"oh":   "",
				"my":   "",
				"goth": "",
			},
		},
	},

	uriTest{
		s: "sip:jart%3e:la%3ewl@google%3e.net:65535" +
			";isup%3e-oli=00%3e;%3eomg;omg;lol=cat",
		uri: sip.URI{
			Scheme: "sip",
			User:   "jart>",
			Pass:   "la>wl",
			Host:   "google>.net",
			Port:   65535,
			Params: sip.Params{
				"isup>-oli": "00>",
				">omg":      "",
				"omg":       "",
				"lol":       "cat",
			},
		},
	},
}

func TestParse(t *testing.T) {
	for _, test := range uriTests {
		uri, err := sip.ParseURI(test.s)
		if err != nil {
			if test.err == nil {
				t.Errorf("%v", err)
				continue
			} else { // test was supposed to fail
				panic("TODO(jart): Implement failing support.")
			}
		}
		if !reflect.DeepEqual(&test.uri, uri) {
			t.Errorf("%#v != %#v", &test.uri, uri)
		}
	}
}

func TestFormat(t *testing.T) {
	for _, test := range uriTests {
		uri := test.uri.String()
		if test.s != uri {
			t.Error(test.s, "!=", uri)
		}
	}
}
