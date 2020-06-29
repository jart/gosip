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

// IANA Assigned VoIP Codec Payload Types
//
// The following codecs have been standardized by IANA thereby
// allowing their 'a=rtpmap' information to be omitted in SDP
// messages.  they've been hard coded to make your life easier.
//
// Many of these codecs (G711, G722, G726, GSM, LPC) have been
// implemented by Steve Underwood in his excellent SpanDSP library.
//
// Newer codecs like silk, broadvoice, speex, etc. use a dynamic
// payload type.
//
// Reference Material:
//
// - IANA Payload Types: http://www.iana.org/assignments/rtp-parameters
// - Explains well-known ITU codecs: http://tools.ietf.org/html/rfc3551
//

package sdp

var (
	ULAWCodec = Codec{PT: 0, Name: "PCMU", Rate: 8000}
	DTMFCodec = Codec{PT: 101, Name: "telephone-event", Rate: 8000, Fmtp: "0-16"}
	Opus      = Codec{PT: 111, Name: "opus", Rate: 48000, Param: "2"}

	StandardCodecs = map[uint8]Codec{
		0:  ULAWCodec,                                      // G.711 μ-Law is the de-facto codec (SpanDSP g711.h)
		3:  {PT: 3, Name: "GSM", Rate: 8000},               // Uncool codec asterisk ppl like (SpanDSP gsm0610.h)
		4:  {PT: 4, Name: "G723", Rate: 8000},              // Worthless.
		5:  {PT: 5, Name: "DVI4", Rate: 8000},              // Adaptive pulse code modulation (SpanDSP ima_adpcm.h)
		6:  {PT: 6, Name: "DVI4", Rate: 16000},             // Adaptive pulse code modulation 16khz (SpanDSP ima_adpcm.h)
		7:  {PT: 7, Name: "LPC", Rate: 8000},               // Chat with your friends ww2 field marshall style (SpanDSP lpc10.h)
		8:  {PT: 8, Name: "PCMA", Rate: 8000},              // G.711 variant of μ-Law used in yurop (SpanDSP g711.h)
		9:  {PT: 9, Name: "G722", Rate: 8000},              // Used for Polycom HD Voice; Rate actually 16khz LOL (SpanDSP g722.h)
		10: {PT: 10, Name: "L16", Rate: 44100, Param: "2"}, // 16-bit signed PCM stereo/mono (mind your MTU; adjust ptime)
		11: {PT: 11, Name: "L16", Rate: 44100},
		12: {PT: 12, Name: "QCELP", Rate: 8000},
		13: {PT: 13, Name: "CN", Rate: 8000}, // RFC3389 comfort noise
		14: {PT: 14, Name: "MPA", Rate: 90000},
		15: {PT: 15, Name: "G728", Rate: 8000},
		16: {PT: 16, Name: "DVI4", Rate: 11025},
		17: {PT: 17, Name: "DVI4", Rate: 22050},
		18: {PT: 18, Name: "G729", Rate: 8000}, // Best telephone voice codec (if you got $10 bucks)
		25: {PT: 25, Name: "CelB", Rate: 90000},
		26: {PT: 26, Name: "JPEG", Rate: 90000},
		28: {PT: 28, Name: "nv", Rate: 90000},
		31: {PT: 31, Name: "H261", Rate: 90000}, // RFC4587 Video
		32: {PT: 32, Name: "MPV", Rate: 90000},
		33: {PT: 33, Name: "MP2T", Rate: 90000},
		34: {PT: 34, Name: "H263", Rate: 90000}, // $$$ video
	}
)
