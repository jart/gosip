package sip

import (
	"bytes"
	"testing"
)

type escapeTest struct {
	name string
	in   string
	out  string
	p    func(byte) bool
}

var escapeTests = []escapeTest{

	escapeTest{
		name: "Param Normal",
		in:   "hello",
		out:  "hello",
		p:    paramc,
	},

	escapeTest{
		name: "User Normal",
		in:   "hello",
		out:  "hello",
		p:    userc,
	},

	escapeTest{
		name: "Param Spacing",
		in:   "hello there",
		out:  "hello%20there",
		p:    paramc,
	},

	escapeTest{
		name: "User Spacing",
		in:   "hello there",
		out:  "hello%20there",
		p:    userc,
	},
}

func TestEscape(t *testing.T) {
	for _, test := range escapeTests {
		var b bytes.Buffer
		appendEscaped(&b, []byte(test.in), test.p)
		out := b.String()
		if test.out != out {
			t.Errorf("%s: %s != %s", test.name, test.out, out)
		}
	}
}
