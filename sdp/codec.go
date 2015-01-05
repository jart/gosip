package sdp

import (
	"bytes"
	"strconv"
)

// Codec describes one of the codec lines in an SDP. This data will be
// magically filled in if the rtpmap wasn't provided (assuming it's a well
// known codec having a payload type less than 96.)
type Codec struct {
	PT    uint8  // 7-bit payload type we need to put in our RTP packets
	Name  string // e.g. PCMU, G729, telephone-event, etc.
	Rate  int    // frequency in hertz.  usually 8000
	Param string // sometimes used to specify number of channels
	Fmtp  string // some extra info; i.e. dtmf might set as "0-16"
}

func (codec *Codec) Append(b *bytes.Buffer) {
	b.WriteString("a=rtpmap:")
	b.WriteString(strconv.FormatInt(int64(codec.PT), 10))
	b.WriteString(" ")
	b.WriteString(codec.Name)
	b.WriteString("/")
	b.WriteString(strconv.Itoa(codec.Rate))
	if codec.Param != "" {
		b.WriteString("/")
		b.WriteString(codec.Param)
	}
	b.WriteString("\r\n")
	if codec.Fmtp != "" {
		b.WriteString("a=fmtp:")
		b.WriteString(strconv.FormatInt(int64(codec.PT), 10))
		b.WriteString(" ")
		b.WriteString(codec.Fmtp)
		b.WriteString("\r\n")
	}
}
