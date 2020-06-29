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

type escapeTest struct {
	name string
	in   string
	out  string
	p    func(byte) bool
}

var escapeTests = []escapeTest{
	{
		name: "Param Normal",
		in:   "hello",
		out:  "hello",
		p:    paramc,
	},
	{
		name: "User Normal",
		in:   "hello",
		out:  "hello",
		p:    userc,
	},
	{
		name: "Param Spacing",
		in:   "hello there",
		out:  "hello%20there",
		p:    paramc,
	},
	{
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
