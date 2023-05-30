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

// Digital Signal Processing

package dsp

// Mixes together two audio frames containing 160 samples. Uses saturation
// adding so you don't hear clipping if ppl talk loud. Uses 128-bit SIMD
// instructions so we can add eight numbers at the same time.
func L16MixSat160(dst, src *int16)

// Compresses a PCM audio sample into a G.711 μ-Law sample. The BSR instruction
// is what makes this code fast.
//
// TODO(jart): How do I make assembly use proper types?
func LinearToUlaw(linear int64) (ulaw int64)

// Turns a μ-Law byte back into an audio sample.
func UlawToLinear(ulaw int64) (linear int64)
