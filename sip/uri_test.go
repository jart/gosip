package sip_test

import (
	"github.com/jart/gosip/sip"
	"reflect"
	"testing"
)

type uriTest struct {
	s   string  // user input we want to convert
	s2  string  // non-blank if 's' changes after we parse/format it
	uri sip.URI // what 's' should become after parsing
	err error   // if we expect parsing to fail
}

var uriTests = []uriTest{

	uriTest{
		s: "sip:bsdtelecom.net",
		uri: sip.URI{
			Scheme: "sip",
			Host:   "bsdtelecom.net",
			Port:   5060,
		},
	},

	uriTest{
		s: "sip:bsdtelecom.net:666",
		uri: sip.URI{
			Scheme: "sip",
			Host:   "bsdtelecom.net",
			Port:   666,
		},
	},

	uriTest{
		s: "sip:jtunney@bsdtelecom.net",
		uri: sip.URI{
			Scheme: "sip",
			User:   "jtunney",
			Host:   "bsdtelecom.net",
			Port:   5060,
		},
	},

	uriTest{
		s:  "sip:jtunney@bsdtelecom.net:5060",
		s2: "sip:jtunney@bsdtelecom.net",
		uri: sip.URI{
			Scheme: "sip",
			User:   "jtunney",
			Host:   "bsdtelecom.net",
			Port:   5060,
		},
	},

	uriTest{
		s:  "sip:jtunney:lawl@bsdtelecom.net:5060",
		s2: "sip:jtunney:lawl@bsdtelecom.net",
		uri: sip.URI{
			Scheme: "sip",
			User:   "jtunney",
			Pass:   "lawl",
			Host:   "bsdtelecom.net",
			Port:   5060,
		},
	},

	uriTest{
		s:  "sip:jtunney:lawl@bsdtelecom.net:5060;isup-oli=00;omg;lol=cat",
		s2: "sip:jtunney:lawl@bsdtelecom.net;isup-oli=00;omg;lol=cat",
		uri: sip.URI{
			Scheme: "sip",
			User:   "jtunney",
			Pass:   "lawl",
			Host:   "bsdtelecom.net",
			Port:   5060,
			Params: sip.Params{
				"isup-oli": "00",
				"omg":      "",
				"lol":      "cat",
			},
		},
	},

	uriTest{
		s: "sip:jtunney@bsdtelecom.net;isup-oli=00;omg;lol=cat",
		uri: sip.URI{
			Scheme: "sip",
			User:   "jtunney",
			Host:   "bsdtelecom.net",
			Port:   5060,
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
			Port:   5060,
		},
	},

	uriTest{
		s:  "sip:[dead:beef::666]:5060",
		s2: "sip:[dead:beef::666]",
		uri: sip.URI{
			Scheme: "sip",
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
		s: "sip:jtunney%3e:la%3ewl@bsdtelecom%3e.net:65535" +
			";isup%3e-oli=00%3e;%3eomg;omg;lol=cat",
		uri: sip.URI{
			Scheme: "sip",
			User:   "jtunney>",
			Pass:   "la>wl",
			Host:   "bsdtelecom>.net",
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
		s := test.s
		if test.s2 != "" {
			s = test.s2
		}
		if s != uri {
			t.Error(test.s, "!=", uri)
		}
	}
}
