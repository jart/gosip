package sip_test

import (
	"github.com/jart/gosip/sip"
	"reflect"
	"testing"
)

type addrTest struct {
	s    string
	s2   string
	addr sip.Addr
	err  error
}

var addrTests = []addrTest{

	addrTest{
		s:  "<sip:pokémon.net>",
		s2: "<sip:pok%c3%a9mon.net>",
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   "pokémon.net",
				Port:   5060,
			},
		},
	},

	addrTest{
		s: "<sip:brave@toaster.net;isup-oli=29>;tag=deadbeef",
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				User:   "brave",
				Host:   "toaster.net",
				Port:   5060,
				Params: sip.Params{
					"isup-oli": "29",
				},
			},
			Params: sip.Params{
				"tag": "deadbeef",
			},
		},
	},

	addrTest{
		s: `<sip:pokemon.com>, "Ditto" <sip:ditto@pokemon.com>`,
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   "pokemon.com",
				Port:   5060,
			},
			Next: &sip.Addr{
				Display: "Ditto",
				Uri: &sip.URI{
					Scheme: "sip",
					User:   "ditto",
					Host:   "pokemon.com",
					Port:   5060,
				},
			},
		},
	},

	addrTest{
		s: `<sip:1.2.3.4>, <sip:1.2.3.5>, <sip:[666::dead:beef]>`,
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   "1.2.3.4",
				Port:   5060,
			},
			Next: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "1.2.3.5",
					Port:   5060,
				},
				Next: &sip.Addr{
					Uri: &sip.URI{
						Scheme: "sip",
						Host:   "666::dead:beef",
						Port:   5060,
					},
				},
			},
		},
	},

	addrTest{
		s:  " hello kitty <sip:jtunney@bsdtelecom.net;isup-oli=29>;tag=deadbeef",
		s2: "\"hello kitty\" <sip:jtunney@bsdtelecom.net;isup-oli=29>;tag=deadbeef",
		addr: sip.Addr{
			Display: "hello kitty",
			Uri: &sip.URI{
				Scheme: "sip",
				User:   "jtunney",
				Host:   "bsdtelecom.net",
				Port:   5060,
				Params: sip.Params{
					"isup-oli": "29",
				},
			},
			Params: sip.Params{
				"tag": "deadbeef",
			},
		},
	},

	addrTest{
		s: "\"\\\"\\\"Justine \\\\Tunney \" " +
			"<sip:jtunney@bsdtelecom.net;isup-oli=29>;tag=deadbeef",
		addr: sip.Addr{
			Display: "\"\"Justine \\Tunney ",
			Uri: &sip.URI{
				Scheme: "sip",
				User:   "jtunney",
				Host:   "bsdtelecom.net",
				Port:   5060,
				Params: sip.Params{
					"isup-oli": "29",
				},
			},
			Params: sip.Params{
				"tag": "deadbeef",
			},
		},
	},
}

func TestParseAddr(t *testing.T) {
	for _, test := range addrTests {
		addr, err := sip.ParseAddr(test.s)
		if err != nil {
			if test.err == nil {
				t.Error(err)
				continue
			} else { // Test was supposed to fail.
				panic("TODO(jart): Implement failing support.")
			}
		}
		if !reflect.DeepEqual(&test.addr, addr) {
			t.Errorf("%#v != %#v", &test.addr, addr)
		}
	}
}

func TestAddrString(t *testing.T) {
	for _, test := range addrTests {
		addr := test.addr.String()
		var s string
		if test.s2 != "" {
			s = test.s2
		} else {
			s = test.s
		}
		if s != addr {
			t.Error(s, "!=", addr)
		}
	}
}
