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

func trace(dir, pkt string, addr net.Addr) {
	size := len(pkt)
	bar := strings.Repeat("-", 72)
	suffix := "\n"
	if pkt != "" && pkt[len(pkt)-1] == '\n' {
		suffix = ""
	}
	log.Printf(
		"%s %d bytes from %s/%s\n"+
			"%s\n"+
			"%s%s"+
			"%s\n",
		dir, size, addr.Network(), addr.String(),
		bar,
		pkt, suffix,
		bar)
}
