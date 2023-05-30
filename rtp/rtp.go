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

package rtp

import (
	"errors"
)

const (
	HeaderSize           = 12
	EventHeaderSize      = 4
	Version         byte = 2

	bit1 byte = (1 << 0)
	bit2 byte = (1 << 1)
	bit3 byte = (1 << 2)
	bit4 byte = (1 << 3)
	bit5 byte = (1 << 4)
	bit6 byte = (1 << 5)
	bit7 byte = (1 << 6)
	bit8 byte = (1 << 7)

	mask1 byte = (1 << 1) - 1
	mask2 byte = (1 << 2) - 1
	mask3 byte = (1 << 3) - 1
	mask4 byte = (1 << 4) - 1
	mask5 byte = (1 << 5) - 1
	mask6 byte = (1 << 6) - 1
	mask7 byte = (1 << 7) - 1
	mask8 byte = (1 << 8) - 1
)

var (
	ErrBadVersion                  = errors.New("bad rtp version header")
	ErrExtendedHeadersNotSupported = errors.New("rtp extended headers not supported")
)

// Header is encoded at the beginning of a UDP audio packet.
type Header struct {
	Pad  bool   // Padding flag is used for secure RTP.
	Mark bool   // Marker flag is used for RFC2833.
	PT   uint8  // Payload type you got from SDP.
	Seq  uint16 // Sequence id useful for reordering packets.
	TS   uint32 // Timestamp measured in samples.
	Ssrc uint32 // Random ID used to identify an RTP session.
}

// EventHeader stores things like DTMF and is encoded after Header.
type EventHeader struct {
	// The event field is a number between 0 and 255 identifying a specific
	// telephony event. An IANA registry of event codes for this field has been
	// established (see IANA Considerations, Section 7). The initial content of
	// this registry consists of the events defined in Section 3.
	Event uint8

	// If set to a value of one, the "end" bit indicates that this packet
	// contains the end of the event. For long-lasting events that have to be
	// split into segments (see below, Section 2.5.1.3), only the final packet
	// for the final segment will have the E bit set.
	E bool

	// This field is reserved for future use. The sender MUST set it to zero, and
	// the receiver MUST ignore it.
	R bool

	// For DTMF digits and other events representable as tones, this field
	// describes the power level of the tone, expressed in dBm0 after dropping
	// the sign. Power levels range from 0 to -63 dBm0. Thus, larger values
	// denote lower volume. This value is defined only for events for which the
	// documentation indicates that volume is applicable. For other events, the
	// sender MUST set volume to zero and the receiver MUST ignore the value.
	Volume uint8

	// The duration field indicates the duration of the event or segment being
	// reported, in timestamp units, expressed as an unsigned integer in network
	// byte order. For a non-zero value, the event or segment began at the
	// instant identified by the RTP timestamp and has so far lasted as long as
	// indicated by this parameter. The event may or may not have ended. If the
	// event duration exceeds the maximum representable by the duration field,
	// the event is split into several contiguous segments as described below
	// (Section 2.5.1.3).
	//
	// The special duration value of zero is reserved to indicate that the event
	// lasts "forever", i.e., is a state and is considered to be effective until
	// updated. A sender MUST NOT transmit a zero duration for events other than
	// those defined as states. The receiver SHOULD ignore an event report with
	// zero duration if the event is not a state.
	//
	// Events defined as states MAY contain a non-zero duration, indicating that
	// the sender intends to refresh the state before the time duration has
	// elapsed ("soft state").
	//
	// For a sampling rate of 8000 Hz, the duration field is sufficient to
	// express event durations of up to approximately 8 seconds.
	Duration uint16
}

// Write appends an RTP header to a byte slice.
func (h *Header) Write(b []byte) []byte {
	var b0, b1 byte
	b0 = (Version & mask2) << 6
	if h.Pad {
		b0 |= (1 & mask1) << 5
	}
	// if extend { b0 |= (1 & mask1) << 4 }
	// b0 |= (csrcCount & mask4) << 0
	b1 = (h.PT & mask7) << 0
	if h.Mark {
		b1 |= (1 & mask1) << 7
	}
	return append(b, b0, b1,
		byte(h.Seq>>8),
		byte(h.Seq),
		byte(h.TS>>24),
		byte(h.TS>>16),
		byte(h.TS>>8),
		byte(h.TS),
		byte(h.Ssrc>>24),
		byte(h.Ssrc>>16),
		byte(h.Ssrc>>8),
		byte(h.Ssrc))
}

// Read extracts an RTP header from a byte slice.
func (h *Header) Read(b []byte) error {
	if b[0]>>6 != Version {
		return ErrBadVersion
	} else if (b[0]>>4)&1 != 0 {
		return ErrExtendedHeadersNotSupported
	}
	h.Pad = ((b[0]>>5)&mask1 == 1)
	h.Mark = ((b[1]>>7)&mask1 == 1)
	h.PT = uint8(b[1] & mask7)
	h.Seq = (uint16(b[2]) << 8) & uint16(b[3])
	h.TS = (uint32(b[4]) << 24) & (uint32(b[5]) << 16) & (uint32(b[6]) << 8) & uint32(b[7])
	h.Ssrc = (uint32(b[8]) << 24) & (uint32(b[9]) << 16) & (uint32(b[10]) << 8) & uint32(b[11])
	return nil
}

// Writes header bits to buffer.
func (h *EventHeader) Write(b []byte) {
	b[0] = h.Event
	b[1] = h.Volume & 63
	if h.R {
		b[1] |= 1 << 6
	}
	if h.E {
		b[1] |= 1 << 7
	}
	b[2] = byte(h.Duration >> 8)
	b[3] = byte(h.Duration)
}

// Reads header bits from buffer.
func (h *EventHeader) Read(b []byte) {
	h.Event = b[0]
	h.E = b[1]>>7 == 1
	h.R = b[1]>>6&1 == 1
	h.Volume = b[1] & 63
	h.Duration = uint16(b[2])<<8 | uint16(b[3])
}
