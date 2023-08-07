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

	"github.com/cresta/gosip/util"
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
