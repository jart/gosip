package sip

import (
	"bytes"
	"testing"
)

type quoteTest struct {
	name string
	in   string
	out  string
}

var quoteTests = []quoteTest{

	quoteTest{
		name: "Normal value",
		in:   "hello",
		out:  "hello",
	},

	quoteTest{
		name: "Space doesn't quotes",
		in:   "hello there",
		out:  "hello there",
	},

	quoteTest{
		name: "Less than adds quotes",
		in:   "hello there<",
		out:  "\"hello there<\"",
	},

	quoteTest{
		name: "CRLF with space after works",
		in:   "hello\r\n there!",
		out:  "\"hello\r\n there!\"",
	},

	quoteTest{
		name: "CRLF without space truncates",
		in:   "hello\r\nthere!",
		out:  "\"hello\"",
	},

	quoteTest{
		name: "Escapable character escapes",
		in:   "hello\"there",
		out:  "\"hello\\\"there\"",
	},

	quoteTest{
		name: "Unescapable character truncates",
		in:   "hello\xFFthere",
		out:  "\"hello\"",
	},
}

func TestQuote(t *testing.T) {
	for _, test := range quoteTests {
		var b bytes.Buffer
		appendQuoted(&b, []byte(test.in))
		out := b.String()
		if test.out != out {
			t.Errorf("%s: %s != %s", test.name, test.out, out)
		}
	}
}
