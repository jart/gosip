package sip

import (
	"flag"
	"log"
	"net"
	"strings"
)

var (
	tracing          = flag.Bool("tracing", true, "Enable SIP message tracing")
	timestampTagging = flag.Bool("timestampTagging", false, "Add microsecond timestamps to Via tags")
)

func trace(dir string, pkt []byte, addr net.Addr) {
	size := len(pkt)
	bar := strings.Repeat("-", 72)
	suffix := "\r\n"
	if pkt != nil && len(pkt) > 0 && pkt[len(pkt)-1] == '\n' {
		suffix = ""
	}
	log.Printf(
		"%s %d bytes from %s/%s\r\n"+
			"%s\r\n"+
			"%s%s"+
			"%s\r\n",
		dir, size, addr.Network(), addr.String(),
		bar,
		pkt, suffix,
		bar)
}
