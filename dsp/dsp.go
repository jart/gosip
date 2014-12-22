// Digital Signal Processing

package dsp

// Mixes together two audio frames containing 160 samples. Uses saturation
// adding so you don't hear clipping if ppl talk loud. Uses 128-bit SIMD
// instructions so we can add eight numbers at the same time.
func L16MixSat160(dst, src *int16)

// Compresses a PCM audio sample into a G.711 μ-Law sample. The BSR instruction
// is what makes this code fast.
func LinearToUlaw(linear int64) (ulaw int64)

// Turns a μ-Law byte back into an audio sample.
func UlawToLinear(ulaw int64) (linear int64)
