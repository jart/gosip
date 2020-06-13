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

package sip

import (
	"flag"
	"log"
	"net"
	"strings"
)

var (
	tracing          = flag.Bool("trace", false, "Enable SIP message tracing")
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
