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

// Additive White Gaussian Noise Generator
//
// This is based off the SpanDSP implementation, which is based off some
// research paper.

package dsp

import (
	"math"
)

const (
	m1           = 259200
	ia1          = 7141
	ic1          = 54773
	rm1          = 1.0 / m1
	m2           = 134456
	ia2          = 8121
	ic2          = 28411
	rm2          = 1.0 / m2
	m3           = 243000
	ia3          = 4561
	ic3          = 51349
	dbm0MaxPower = 3.14 + 3.02
)

// Awgn contains the white noise generator state information.
type AWGN struct {
	rms  float64
	ix1  int64
	ix2  int64
	ix3  int64
	r    [98]float64
	gset float64
	iset int
}

// Returns a new comfort noise generator. The volume param is in decibels and
// -50.0 is a good value. The closer you get to 0.0, the louder the noise will
// become.
func NewAWGN(volume float64) (s *AWGN) {
	return NewAWGN_DBM0(7162534, volume)
}

func NewAWGN_DBM0(idum int64, level float64) (s *AWGN) {
	return NewAWGN_DBOV(idum, level-dbm0MaxPower)
}

func NewAWGN_DBOV(idum int64, level float64) (s *AWGN) {
	s = new(AWGN)
	if idum < 0.0 {
		idum = -idum
	}
	s.rms = math.Pow(10.0, level/20.0) * 32768.0
	s.ix1 = (ic1 + idum) % m1
	s.ix1 = (ia1*s.ix1 + ic1) % m1
	s.ix2 = s.ix1 % m2
	s.ix1 = (ia1*s.ix1 + ic1) % m1
	s.ix3 = s.ix1 % m3
	s.r[0] = 0.0
	for j := 1; j <= 97; j++ {
		s.ix1 = (ia1*s.ix1 + ic1) % m1
		s.ix2 = (ia2*s.ix2 + ic2) % m2
		s.r[j] = (float64(s.ix1) + float64(s.ix2)*rm2) * rm1
	}
	s.gset = 0.0
	s.iset = 0
	return
}

func (s *AWGN) Get() int16 {
	var fac, r, v1, v2, amp float64
	if s.iset == 0 {
		for {
			v1 = 2.0*s.ran1() - 1.0
			v2 = 2.0*s.ran1() - 1.0
			r = v1*v1 + v2*v2
			if r < 1.0 {
				break
			}
		}
		fac = math.Sqrt(-2.0 * math.Log(r) / r)
		s.gset = v1 * fac
		s.iset = 1
		amp = v2 * fac * s.rms
	} else {
		s.iset = 0
		amp = s.gset * s.rms
	}
	return fsaturate(amp)
}

func (s *AWGN) ran1() (res float64) {
	var j int64
	s.ix1 = (ia1*s.ix1 + ic1) % m1
	s.ix2 = (ia2*s.ix2 + ic2) % m2
	s.ix3 = (ia3*s.ix3 + ic3) % m3
	j = 1 + ((97 * s.ix3) / m3)
	if j > 97 || j < 1 {
		return -1
	}
	res = s.r[j]
	s.r[j] = (float64(s.ix1) + float64(s.ix2)*rm2) * rm1
	return
}

// TODO(jart): How does FISTP work in Plan9 assembly?
func fsaturate(damp float64) int16 {
	if damp > float64(math.MaxInt16) {
		return math.MaxInt16
	}
	if damp < float64(math.MinInt16) {
		return math.MinInt16
	}
	return int16(damp)
	// return int16(lrint(damp))
}
