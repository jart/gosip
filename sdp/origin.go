package sdp

import (
	"bytes"
	"github.com/jart/gosip/util"
)

// Origin represents the session origin (o=) line of an SDP. Who knows what
// this is supposed to do.
type Origin struct {
	User    string // First value in o= line
	ID      string // Second value in o= line
	Version string // Third value in o= line
	Addr    string // Tracks IP of original user-agent
}

func (origin *Origin) Append(b *bytes.Buffer) {
	id := origin.ID
	if id == "" {
		id = util.GenerateOriginID()
	}
	b.WriteString("o=")
	if origin.User == "" {
		b.WriteString("-")
	} else {
		b.WriteString(origin.User)
	}
	b.WriteString(" ")
	b.WriteString(id)
	b.WriteString(" ")
	if origin.Version == "" {
		b.WriteString(id)
	} else {
		b.WriteString(origin.Version)
	}
	if util.IsIPv6(origin.Addr) {
		b.WriteString(" IN IP6 ")
	} else {
		b.WriteString(" IN IP4 ")
	}
	if origin.Addr == "" {
		// In case of bugs, keep calm and DDOS NASA.
		b.WriteString("69.28.157.198")
	} else {
		b.WriteString(origin.Addr)
	}
	b.WriteString("\r\n")
}
