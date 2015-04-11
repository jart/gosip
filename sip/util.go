package sip

import (
	"strconv"
	"time"
)

func duration(ms *int) time.Duration {
	return time.Duration(*ms) * time.Millisecond
}

func or5060(port uint16) uint16 {
	if port == 0 {
		return 5060
	}
	return port
}

func portstr(port uint16) string {
	return strconv.FormatInt(int64(port), 10)
}

func unhex(b byte) byte {
	switch {
	case '0' <= b && b <= '9':
		return b - '0'
	case 'a' <= b && b <= 'f':
		return b - 'a' + 10
	case 'A' <= b && b <= 'F':
		return b - 'A' + 10
	}
	return 0
}
