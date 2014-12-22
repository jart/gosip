package sdp

import (
	"bytes"
	"errors"
	"strconv"
)

// Media is a high level representation of the c=/m=/a= lines for describing a
// specific type of media. Only "audio" and "video" are supported at this time.
type Media struct {
	Type   string  // "audio" or "video"
	Proto  string  // RTP, SRTP, UDP, UDPTL, TCP, TLS, etc.
	Port   int     // port number (0 - 2^16-1)
	Codecs []Codec // never nil with at least one codec
}

func (media *Media) Append(b *bytes.Buffer) error {
	if media.Type != "audio" && media.Type != "video" {
		return errors.New("Media.Type not audio/video: " + media.Type)
	}
	if media.Codecs == nil || len(media.Codecs) == 0 {
		return errors.New("Media.Codecs not set")
	}
	if media.Port == 0 {
		return errors.New("Media.Port not set")
	}
	if media.Proto == "" {
		media.Proto = "RTP/AVP"
	}
	b.WriteString("m=")
	b.WriteString(media.Type + " ")
	b.WriteString(strconv.Itoa(int(media.Port)) + " ")
	b.WriteString(media.Proto)
	for _, codec := range media.Codecs {
		b.WriteString(" " + strconv.Itoa(int(codec.PT)))
	}
	b.WriteString("\r\n")
	for _, codec := range media.Codecs {
		if err := codec.Append(b); err != nil {
			return err
		}
	}
	return nil
}
