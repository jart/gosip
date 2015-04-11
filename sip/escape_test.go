package sip

import (
	"testing"
)

type escapeTest struct {
	name string
	in   string
	out  string
	p    func([]byte) []byte
}

var escapeTests = []escapeTest{

	escapeTest{
		name: "Param Normal",
		in:   "hello",
		out:  "hello",
		p:    escapeParam,
	},

	escapeTest{
		name: "User Normal",
		in:   "hello",
		out:  "hello",
		p:    escapeUser,
	},

	escapeTest{
		name: "Param Spacing",
		in:   "hello there",
		out:  "hello%20there",
		p:    escapeParam,
	},

	escapeTest{
		name: "User Spacing",
		in:   "hello there",
		out:  "hello%20there",
		p:    escapeUser,
	},
}

func TestEscape(t *testing.T) {
	for _, test := range escapeTests {
		out := string(test.p([]byte(test.in)))
		if test.out != out {
			t.Errorf("%s: %s != %s", test.name, test.out, out)
		}
	}
}

func BenchmarkEscapeParam(b *testing.B) {
	for i := 0; i < b.N; i++ {
		escapeParam([]byte("hello there"))
	}
}
