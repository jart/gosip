package util

import (
	"bytes"
	"encoding/hex"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"syscall"
)

// Returns true if error is an i/o timeout.
func IsTimeout(err error) bool {
	operr, ok := err.(*net.OpError)
	return ok && operr.Timeout()
}

// Returns true if error is ICMP connection refused.
func IsRefused(err error) bool {
	operr, ok := err.(*net.OpError)
	return ok && operr.Err == syscall.ECONNREFUSED
}

// Returns true if error was caused by reading from a closed socket.
func IsUseOfClosed(err error) bool {
	return strings.Contains(err.Error(), "use of closed network connection")
}

// Returns true if IP contains a colon.
func IsIPv6(ip string) bool {
	n := strings.Index(ip, ":")
	return n >= 0
}

// Returns true if IPv4 and is a LAN address as defined by RFC 1918.
func IsIPPrivate(ip net.IP) bool {
	if ip != nil {
		ip = ip.To4()
		if ip != nil {
			switch ip[0] {
			case 10:
				return true
			case 172:
				if ip[1] >= 16 && ip[1] <= 31 {
					return true
				}
			case 192:
				if ip[1] == 168 {
					return true
				}
			}
		}
	}
	return false
}

// Generates a secure random number between 0 and 50,000.
func GenerateCSeq() int {
	return rand.Int() % 50000
}

// Generates a 48-bit secure random string like: 27c97271d363.
func GenerateTag() string {
	return hex.EncodeToString(randomBytes(6))
}

// This is used in the Via tag. Probably not suitable for use by stateless
// proxies.
func GenerateBranch() string {
	return "z9hG4bK-" + GenerateTag()
}

// Generates a secure UUID4, e.g. f47ac10b-58cc-4372-a567-0e02b2c3d479
func GenerateCallID() []byte {
	lol := randomBytes(15)
	digs := hex.EncodeToString(lol)
	var b bytes.Buffer
	b.WriteString(digs[0:8])
	b.WriteByte('-')
	b.WriteString(digs[8:12])
	b.WriteString("-4")
	b.WriteString(digs[12:15])
	b.WriteString("-a")
	b.WriteString(digs[15:18])
	b.WriteByte('-')
	b.WriteString(digs[18:])
	return b.Bytes()
}

// Generates a random ID for an SDP.
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
