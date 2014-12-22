// Global Settings.  You can change these at startup to fine tune
// certain behaviors.

package gosip

import (
	"os"
	"time"
)

var (
	// how often to check for shutdowns
	DeathClock = 200 * time.Millisecond

	// sip egress msg length must not exceed me.  if you are brave and
	// use jumbo frames you can increase this
	MTU = 1450

	// maximum number SRV/Redirects/Etc. to entertain.  more
	// accurately, this is the maximum number of time's we're allowed
	// to enter the "calling" state.
	MaxAttempts = 10

	// this is how long to wait (in nanoseconds) for a 100 trying
	// response before retransmitting the invite.  multiply this by
	// RetransmitAttempts and that's how long it'll be before we try
	// another server.  that might seem like a very long time but
	// happy networks facilitate fast failover by sending hard errors
	// (ICMP, 480, 500-599)
	TryingTimeout = 500 * time.Millisecond

	// how many times to attempt to retransmit before moving on
	RetransmitAttempts = 2

	// number of nanoseconds (across all attempts) before we give up
	// trying to connect a call.  this doesn't mean the call has to
	// have been answered but rather has moved beyond the "calling"
	// state and doesn't go back.
	GiveupTimeout = 3 * time.Second

	// ToDo: how long to wait before refreshing a call with a
	// re-INVITE message.
	RefreshTimeout = 15 * time.Minute

	// how many times are proxies allowed to forward a message
	MaxForwards int = 70

	// aprox. how long the transaction engine remembers old Call-IDs
	// to prevent a retransmit from accidentally opening a new dialog.
	CallIDBanDuration = 35 * time.Second

	// aprox. how often to sweep old transactions from ban list
	CallIDBanSweepDuration = 5 * time.Second

	// should the transport layer add timestamps to Via headers
	// (µsi/µse = utc microseconds since unix epoch ingress/egress)
	ViaTimestampTagging = true

	// use this feature to print out raw sip messages the moment they
	// are sent/received by the transport layer
	TransportTrace = (os.Getenv("TPORT_LOG") == "true")
)
