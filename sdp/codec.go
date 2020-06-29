// Copyright 2020 Justine Alexandra Roberts Tunney
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
