package dsp_test

import (
	"github.com/jart/gosip/dsp"
	"testing"
)

func TestAWGN(t *testing.T) {
	awgn := dsp.NewAWGN(-50.0)
	samp := []int16{64, -28, 1, 34, -73}
	for i, want := range samp {
		got := awgn.Get()
		if want != got {
			t.Errorf("sample #%d: wanted %d but got %d", i, want, got)
		}
	}
}
