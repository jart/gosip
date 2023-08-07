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

func TestL16MixSat160(t *testing.T) {
	var x, y [160]int16
	for n := 0; n < 160; n++ {
		x[n] = int16(n)
		y[n] = int16(666)
	}
	dsp.L16MixSat160(&x[0], &y[0])
	for n := 0; n < 160; n++ {
		want := int16(n + 666)
		if x[n] != want {
			t.Errorf("x[%v] = %v (wanted: %v)", n, x[n], want)
			return
		}
		if y[n] != int16(666) {
			t.Errorf("side effect y[%v] = %v", n, y[n])
			return
		}
	}
}

var ulawTable = []int64{
	-32124, -31100, -30076, -29052, -28028, -27004, -25980, -24956, -23932,
	-22908, -21884, -20860, -19836, -18812, -17788, -16764, -15996, -15484,
	-14972, -14460, -13948, -13436, -12924, -12412, -11900, -11388, -10876,
	-10364, -9852, -9340, -8828, -8316, -7932, -7676, -7420, -7164,
	-6908, -6652, -6396, -6140, -5884, -5628, -5372, -5116, -4860,
	-4604, -4348, -4092, -3900, -3772, -3644, -3516, -3388, -3260,
	-3132, -3004, -2876, -2748, -2620, -2492, -2364, -2236, -2108,
	-1980, -1884, -1820, -1756, -1692, -1628, -1564, -1500, -1436,
	-1372, -1308, -1244, -1180, -1116, -1052, -988, -924, -876,
	-844, -812, -780, -748, -716, -684, -652, -620, -588,
	-556, -524, -492, -460, -428, -396, -372, -356, -340,
	-324, -308, -292, -276, -260, -244, -228, -212, -196,
	-180, -164, -148, -132, -120, -112, -104, -96, -88,
	-80, -72, -64, -56, -48, -40, -32, -24, -16,
	-8, 0, 32124, 31100, 30076, 29052, 28028, 27004, 25980,
	24956, 23932, 22908, 21884, 20860, 19836, 18812, 17788, 16764,
	15996, 15484, 14972, 14460, 13948, 13436, 12924, 12412, 11900,
	11388, 10876, 10364, 9852, 9340, 8828, 8316, 7932, 7676,
	7420, 7164, 6908, 6652, 6396, 6140, 5884, 5628, 5372,
	5116, 4860, 4604, 4348, 4092, 3900, 3772, 3644, 3516,
	3388, 3260, 3132, 3004, 2876, 2748, 2620, 2492, 2364,
	2236, 2108, 1980, 1884, 1820, 1756, 1692, 1628, 1564,
	1500, 1436, 1372, 1308, 1244, 1180, 1116, 1052, 988,
	924, 876, 844, 812, 780, 748, 716, 684, 652,
	620, 588, 556, 524, 492, 460, 428, 396, 372,
	356, 340, 324, 308, 292, 276, 260, 244, 228,
	212, 196, 180, 164, 148, 132, 120, 112, 104,
	96, 88, 80, 72, 64, 56, 48, 40, 32, 24, 16, 8, 0,
}

func TestLinearToUlaw(t *testing.T) {
	if dsp.LinearToUlaw(0) != 255 {
		t.Error("omg")
	}
	if dsp.LinearToUlaw(-100) != 114 {
		t.Error("omg")
	}
}

func TestUlawToLinear(t *testing.T) {
	for n := 0; n <= 255; n++ {
		if dsp.UlawToLinear(int64(n)) != ulawTable[n] {
			t.Error("omg")
			return
		}
	}
}

func TestLinearToUlawToLinear(t *testing.T) {
	for n := 0; n <= 255; n++ {
		if dsp.UlawToLinear(dsp.LinearToUlaw(ulawTable[n])) != ulawTable[n] {
			t.Error("omg")
			return
		}
	}
}
