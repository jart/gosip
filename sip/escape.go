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
)

const (
	hexChars = "0123456789abcdef"
)

// appendEscaped appends while URL encoding bytes that don't match the predicate..
func appendEscaped(b *bytes.Buffer, s []byte, p func(byte) bool) {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if p(c) {
			b.WriteByte(c)
		} else {
			b.WriteByte('%')
			b.WriteByte(hexChars[c>>4])
			b.WriteByte(hexChars[c%16])
		}
	}
}
