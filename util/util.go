package util

import (
	"encoding/hex"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

// Returns true if error is an i/o timeout.
func IsTimeout(err error) bool {
	if operr, ok := err.(*net.OpError); ok {
		return operr.Timeout()
	}
	return false
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

func HostPortToString(host string, port uint16) (saddr string) {
	sport := strconv.FormatInt(int64(port), 10)
	if IsIPv6(host) {
		saddr = "[" + host + "]:" + sport
	} else {
		saddr = host + ":" + sport
	}
	return saddr
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

// Generates a secure UUID4, e.g.f47ac10b-58cc-4372-a567-0e02b2c3d479
func GenerateCallID() string {
	lol := randomBytes(15)
	digs := hex.EncodeToString(lol)
	uuid4 := digs[0:8] + "-" + digs[8:12] + "-4" + digs[12:15] +
		"-a" + digs[15:18] + "-" + digs[18:]
	return uuid4
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
