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

// Media is a high level representation of the c=/m=/a= lines for describing a
// specific type of media. Only "audio" and "video" are supported at this time.
type Media struct {
	Proto  string  // RTP, SRTP, UDP, UDPTL, TCP, TLS, etc.
	Port   uint16  // Port number (0 - 2^16-1)
	Codecs []Codec // Collection of codecs of a specific type.
}

func (media *Media) Append(type_ string, b *bytes.Buffer) {
	b.WriteString("m=")
	b.WriteString(type_)
	b.WriteString(" ")
	b.WriteString(strconv.FormatUint(uint64(media.Port), 10))
	b.WriteString(" ")
	if media.Proto == "" {
		b.WriteString("RTP/AVP")
	} else {
		b.WriteString(media.Proto)
	}
	for _, codec := range media.Codecs {
		b.WriteString(" ")
		b.WriteString(strconv.FormatInt(int64(codec.PT), 10))
	}
	b.WriteString("\r\n")
	for _, codec := range media.Codecs {
		codec.Append(b)
	}
}
