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

package dsp_test

import (
	"testing"

	"github.com/cresta/gosip/dsp"
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
