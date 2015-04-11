package sip

import (
	"testing"
)

func BenchmarkParamc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		paramc('a')
	}
}
