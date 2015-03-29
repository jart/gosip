package sip_test

import (
	"github.com/jart/gosip/sip"
	"reflect"
	"testing"
)

type addrTest struct {
	s    string
	addr sip.Addr
	err  error
}

var addrTests = []addrTest{

	addrTest{
		s: "<sip:pokemon.net>",
		addr: sip.Addr{
			Uri: &sip.URI{
				Scheme: "sip",
				Host:   "pokemon.net",
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
		s: `<sip:pokemon.com>,"Ditto" <sip:ditto@pokemon.com>`,
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

	addrTest{
		s: `<sip:1.2.3.4>,<sip:1.2.3.5>,<sip:[666::dead:beef]>`,
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

	addrTest{
		s: "\"\\\"\\\"Justine \\\\Tunney \" " +
			"<sip:jart@google.com;isup-oli=29>;tag=deadbeef",
		addr: sip.Addr{
			Display: "\"\"Justine \\Tunney ",
			Uri: &sip.URI{
				Scheme: "sip",
				User:   "jart",
				Host:   "google.com",
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
				t.Fatal("TODO(jart): Implement failing support.")
			}
		}
		if !reflect.DeepEqual(&test.addr, addr) {
			t.Errorf("%#v != %#v\n%#v != %#v", &test.addr, addr, test.addr.Uri, addr.Uri)
		}
	}
}

func TestAddrString(t *testing.T) {
	for _, test := range addrTests {
		addr := test.addr.String()
		if test.s != addr {
			t.Error(test.s, "!=", addr)
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
