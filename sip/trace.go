package sip

import (
	"log"
	"net"
	"strings"
	"time"
)

func trace(dir, pkt string, addr net.Addr, t time.Time) {
	size := len(pkt)
	bar := strings.Repeat("-", 72)
	suffix := "\n"
	if pkt[len(pkt)-1] == '\n' {
		suffix = ""
	}
	log.Printf(
		"%s %d bytes to %s/%s at %s\n"+
			"%s\n"+
			"%s%s"+
			"%s\n",
		dir, size, addr.Network(), addr.String(),
		t.Format(time.RFC3339Nano),
		bar,
		pkt, suffix,
		bar)
}
