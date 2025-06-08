package rtp_test

import (
	"reflect"
	"testing"

	"github.com/jart/gosip/rtp"
)

func TestRTPRoundTrip(t *testing.T) {
	h := rtp.Header{
		Pad:  false,
		Mark: true,
		PT:   42,
		Seq:  1234,
		TS:   567891234,
		Ssrc: 6789,
	}

	b := h.Write(nil)
	var h2 rtp.Header
	err := h2.Read(b)
	if err != nil {
		t.Fatalf("Read RTP: %v", err)
	}

	if !reflect.DeepEqual(h, h2) {
		t.Errorf("%#v != %#v", h2, h)
	}
}
