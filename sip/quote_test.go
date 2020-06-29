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
	{
		name: "Normal value",
		in:   "hello",
		out:  "hello",
	},
	{
		name: "Space doesn't quotes",
		in:   "hello there",
		out:  "hello there",
	},
	{
		name: "Less than adds quotes",
		in:   "hello there<",
		out:  "\"hello there<\"",
	},
	{
		name: "CRLF with space after works",
		in:   "hello\r\n there!",
		out:  "\"hello\r\n there!\"",
	},
	{
		name: "CRLF without space truncates",
		in:   "hello\r\nthere!",
		out:  "\"hello\"",
	},
	{
		name: "Escapable character escapes",
		in:   "hello\"there",
		out:  "\"hello\\\"there\"",
	},
	{
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
