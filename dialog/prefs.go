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

// Global Settings.  You can change these at startup to fine tune
// certain behaviors.

package dialog

import (
	"os"
	"time"
)

var (
	// How often to check for shutdowns.
	DeathClock = 200 * time.Millisecond

	// SIP egress msg length must not exceed me. If you are brave and use jumbo
	// frames you can increase this value.
	MTU = 1450

	// Maximum number SRV/Redirects/Etc. to entertain. More accurately, this is
	// the maximum number of time's we're allowed to enter the "calling" state.
	MaxAttempts = 10

	// This is how long to wait (in nanoseconds) for a 100 trying response before
	// retransmitting the invite. Multiply this by RetransmitAttempts and that's
	// how long it'll be before we try another server. That might seem like a
	// very long time but happy networks facilitate fast failover by sending hard
	// errors (ICMP, 480, 500-599)
	TryingTimeout = 500 * time.Millisecond

	// How many times to attempt to retransmit before moving on.
	RetransmitAttempts = 2

	// Number of nanoseconds (across all attempts) before we give up trying to
	// connect a call. This doesn't mean the call has to have been answered but
	// rather has moved beyond the "calling" state and doesn't go back.
	GiveupTimeout = 3 * time.Second

	// TODO(jart): How long to wait before refreshing a call with a re-INVITE
	//             message.
	RefreshTimeout = 15 * time.Minute

	// How many times are proxies allowed to forward a message.
	MaxForwards int = 70

	// Approx. how long the transaction engine remembers old Call-IDs to prevent
	// a retransmit from accidentally opening a new dialog.
	CallIDBanDuration = 35 * time.Second

	// Approx. how often to sweep old transactions from ban list.
	CallIDBanSweepDuration = 5 * time.Second

	// Should the transport layer add timestamps to Via headers (µsi/µse = UTC
	// microseconds since unix epoch ingress/egress).
	ViaTimestampTagging = true

	// Use this feature to print out raw sip messages the moment they are
	// sent/received by the transport layer.
	TransportTrace = (os.Getenv("TPORT_LOG") == "true")
)
