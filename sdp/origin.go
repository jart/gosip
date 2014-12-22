package sdp

import (
	"bytes"
	"errors"
	"github.com/jart/gosip/util"
)

// Origin represents the session origin (o=) line of an SDP. Who knows what
// this is supposed to do.
type Origin struct {
	User    string // first value in o= line
	ID      string // second value in o= line
	Version string // third value in o= line
	Addr    string // tracks ip of original user-agent
}

func (origin *Origin) Append(b *bytes.Buffer) error {
	if origin.ID == "" {
		return errors.New("sdp missing origin id")
	}
	if origin.Version == "" {
		return errors.New("sdp missing origin version")
	}
	if origin.Addr == "" {
		return errors.New("sdp missing origin address")
	}
	if origin.User == "" {
		origin.User = "-"
	}
	b.WriteString("o=")
	b.WriteString(origin.User + " ")
	b.WriteString(origin.ID + " ")
	b.WriteString(origin.Version)
	if util.IsIPv6(origin.Addr) {
		b.WriteString(" IN IP6 " + origin.Addr + "\r\n")
	} else {
		b.WriteString(" IN IP4 " + origin.Addr + "\r\n")
	}
	return nil
}
