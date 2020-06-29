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

package util

import (
	"encoding/hex"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"syscall"
)

// Return true if error is ICMP connection refused.
func IsRefused(err error) bool {
	operr, ok := err.(*net.OpError)
	return ok && operr.Err == syscall.ECONNREFUSED
}

// Return true if error was caused by reading from a closed socket.
func IsUseOfClosed(err error) bool {
	return strings.Contains(err.Error(), "use of closed network connection")
}

// Return true if IP contains a colon.
func IsIPv6(ip string) bool {
	return strings.Index(ip, ":") >= 0
}

// Generate a secure random number between 0 and 50,000.
func GenerateCSeq() int {
	return rand.Int() % 50000
}

// Generate a 48-bit secure random string like: 27c97271d363.
func GenerateTag() string {
	return hex.EncodeToString(randomBytes(6))
}

// Generate a SIP 2.0 Via branch ID. This is probably not suitable for use by
// stateless proxies.
func GenerateBranch() string {
	return "z9hG4bK-" + GenerateTag()
}

// Generate a secure UUID4, e.g.f47ac10b-58cc-4372-a567-0e02b2c3d479
func GenerateCallID() string {
	lol := randomBytes(15)
	digs := hex.EncodeToString(lol)
	uuid4 := digs[0:8] +
		"-" + digs[8:12] +
		"-4" + digs[12:15] +
		"-a" + digs[15:18] +
		"-" + digs[18:]
	return uuid4
}

// Generate a random ID for an SDP.
func GenerateOriginID() string {
	return strconv.FormatUint(uint64(rand.Uint32()), 10)
}

func randomBytes(l int) (b []byte) {
	b = make([]byte, l)
	for i := 0; i < l; i++ {
		b[i] = byte(rand.Int())
	}
	return
}

func append(buf []byte, s string) []byte {
	lenb, lens := len(buf), len(s)
	if lenb+lens <= cap(buf) {
		buf = buf[0 : lenb+lens]
	} else {
		panic("mtu exceeded D:")
	}
	for i := 0; i < lens; i++ {
		buf[lenb+i] = byte(s[i])
	}
	return buf
}

func Portstr(port uint16) string {
	return strconv.FormatInt(int64(port), 10)
}
