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

// Charset implementation using int64 bitmasks

package sip

var (
	tokencMask  [2]uint64
	qdtextcMask [2]uint64
	usercMask   [2]uint64
	passcMask   [2]uint64
	paramcMask  [2]uint64
	headercMask [2]uint64
)

func tokenc(c byte) bool {
	return charsetContains(&tokencMask, c)
}

func qdtextc(c byte) bool {
	return charsetContains(&qdtextcMask, c)
}

func userc(c byte) bool {
	return charsetContains(&usercMask, c)
}

func passc(c byte) bool {
	return charsetContains(&passcMask, c)
}

func paramc(c byte) bool {
	return charsetContains(&paramcMask, c)
}

func headerc(c byte) bool {
	return charsetContains(&headercMask, c)
}

func qdtextesc(c byte) bool {
	return 0x00 <= c && c <= 0x09 ||
		0x0B <= c && c <= 0x0C ||
		0x0E <= c && c <= 0x7F
}

func whitespacec(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}

func init() {
	charsetAddAlphaNumeric(&tokencMask)
	charsetAdd(&tokencMask, '-', '.', '!', '%', '*', '_', '+', '`', '\'', '~')

	charsetAddRange(&qdtextcMask, 0x23, 0x5B)
	charsetAddRange(&qdtextcMask, 0x5D, 0x7E)
	charsetAdd(&qdtextcMask, '\r', '\n', '\t', ' ', '!')

	charsetAddAlphaNumeric(&usercMask)
	charsetAddMark(&usercMask)
	charsetAdd(&usercMask, '&', '=', '+', '$', ',', ';', '?', '/')

	charsetAddAlphaNumeric(&passcMask)
	charsetAddMark(&passcMask)
	charsetAdd(&passcMask, '&', '=', '+', '$', ',')

	charsetAddAlphaNumeric(&paramcMask)
	charsetAddMark(&paramcMask)
	charsetAdd(&paramcMask, '[', ']', '/', ':', '&', '+', '$')

	charsetAddAlphaNumeric(&headercMask)
	charsetAddMark(&headercMask)
	charsetAdd(&headercMask, '[', ']', '/', '?', ':', '+', '$')
}

func charsetContains(mask *[2]uint64, i byte) bool {
	return i < 128 && mask[i/64]&(1<<(i%64)) != 0
}

func charsetAdd(mask *[2]uint64, vi ...byte) {
	for _, i := range vi {
		mask[i/64] |= 1 << (i % 64)
	}
}

func charsetAddRange(mask *[2]uint64, a, b byte) {
	for i := a; i <= b; i++ {
		charsetAdd(mask, i)
	}
}

func charsetAddMark(mask *[2]uint64) {
	charsetAdd(mask, '-', '_', '.', '!', '~', '*', '\'', '(', ')')
}

func charsetAddAlphaNumeric(mask *[2]uint64) {
	charsetAddRange(mask, 'a', 'z')
	charsetAddRange(mask, 'A', 'Z')
	charsetAddRange(mask, '0', '9')
}
