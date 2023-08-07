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

package sdp_test

import (
	"fmt"
	"testing"

	"github.com/cresta/gosip/sdp"
)

type sdpTest struct {
	name string   // arbitrary name for test
	s    string   // raw sdp input to parse
	s2   string   // non-blank if sdp looks different when we format it
	sdp  *sdp.SDP // memory structure of 's' after parsing
	err  error
}

var sdpTests = []sdpTest{

	{
		name: "Asterisk PCMU+DTMF",
		s: ("v=0\r\n" +
			"o=root 31589 31589 IN IP4 10.0.0.38\r\n" +
			"s=session\r\n" +
			"c=IN IP4 10.0.0.38\r\n" +
			"t=0 0\r\n" +
			"m=audio 30126 RTP/AVP 0 101\r\n" +
			"a=rtpmap:0 PCMU/8000\r\n" +
			"a=rtpmap:101 telephone-event/8000\r\n" +
			"a=fmtp:101 0-16\r\n" +
			"a=silenceSupp:off - - - -\r\n" +
			"a=ptime:20\r\n" +
			"a=sendrecv\r\n"),
		sdp: &sdp.SDP{
			Origin: sdp.Origin{
				User:    "root",
				ID:      "31589",
				Version: "31589",
				Addr:    "10.0.0.38",
			},
			Session: "session",
			Time:    "0 0",
			Addr:    "10.0.0.38",
			Audio: &sdp.Media{
				Proto: "RTP/AVP",
				Port:  30126,
				Codecs: []sdp.Codec{
					{PT: 0, Name: "PCMU", Rate: 8000},
					{PT: 101, Name: "telephone-event", Rate: 8000, Fmtp: "0-16"},
				},
			},
			Attrs: [][2]string{
				{"silenceSupp", "off - - - -"},
			},
			Ptime: 20,
		},
	},

	{
		name: "Audio+Video+Implicit+Fmtp",
		s: "v=0\r\n" +
			"o=- 3366701332 3366701332 IN IP4 1.2.3.4\r\n" +
			"c=IN IP4 1.2.3.4\r\n" +
			"m=audio 32898 RTP/AVP 18\r\n" +
			"m=video 32900 RTP/AVP 34\r\n" +
			"a=fmtp:18 annexb=yes",
		s2: "v=0\r\n" +
			"o=- 3366701332 3366701332 IN IP4 1.2.3.4\r\n" +
			"s=pokémon\r\n" +
			"c=IN IP4 1.2.3.4\r\n" +
			"t=0 0\r\n" +
			"m=audio 32898 RTP/AVP 18\r\n" +
			"a=rtpmap:18 G729/8000\r\n" +
			"a=fmtp:18 annexb=yes\r\n" +
			"m=video 32900 RTP/AVP 34\r\n" +
			"a=rtpmap:34 H263/90000\r\n" +
			"a=sendrecv\r\n",
		sdp: &sdp.SDP{
			Origin: sdp.Origin{
				User:    "-",
				ID:      "3366701332",
				Version: "3366701332",
				Addr:    "1.2.3.4",
			},
			Addr:    "1.2.3.4",
			Session: "pokémon",
			Time:    "0 0",
			Audio: &sdp.Media{
				Proto: "RTP/AVP",
				Port:  32898,
				Codecs: []sdp.Codec{
					{PT: 18, Name: "G729", Rate: 8000, Fmtp: "annexb=yes"},
				},
			},
			Video: &sdp.Media{
				Proto: "RTP/AVP",
				Port:  32900,
				Codecs: []sdp.Codec{
					{PT: 34, Name: "H263", Rate: 90000},
				},
			},
			Attrs: [][2]string{},
		},
	},

	{
		name: "Implicit Codecs",
		s: "v=0\r\n" +
			"o=- 3366701332 3366701332 IN IP4 1.2.3.4\r\n" +
			"s=-\r\n" +
			"c=IN IP4 1.2.3.4\r\n" +
			"t=0 0\r\n" +
			"m=audio 32898 RTP/AVP 9 18 0 101\r\n" +
			"a=rtpmap:101 telephone-event/8000\r\n" +
			"a=ptime:20\r\n",
		s2: "v=0\r\n" +
			"o=- 3366701332 3366701332 IN IP4 1.2.3.4\r\n" +
			"s=-\r\n" +
			"c=IN IP4 1.2.3.4\r\n" +
			"t=0 0\r\n" +
			"m=audio 32898 RTP/AVP 9 18 0 101\r\n" +
			"a=rtpmap:9 G722/8000\r\n" +
			"a=rtpmap:18 G729/8000\r\n" +
			"a=rtpmap:0 PCMU/8000\r\n" +
			"a=rtpmap:101 telephone-event/8000\r\n" +
			"a=ptime:20\r\n" +
			"a=sendrecv\r\n",
		sdp: &sdp.SDP{
			Origin: sdp.Origin{
				User:    "-",
				ID:      "3366701332",
				Version: "3366701332",
				Addr:    "1.2.3.4",
			},
			Session: "-",
			Time:    "0 0",
			Addr:    "1.2.3.4",
			Audio: &sdp.Media{
				Proto: "RTP/AVP",
				Port:  32898,
				Codecs: []sdp.Codec{
					{PT: 9, Name: "G722", Rate: 8000},
					{PT: 18, Name: "G729", Rate: 8000},
					{PT: 0, Name: "PCMU", Rate: 8000},
					{PT: 101, Name: "telephone-event", Rate: 8000},
				},
			},
			Ptime: 20,
			Attrs: [][2]string{},
		},
	},

	{
		name: "IPv6",
		s: "v=0\r\n" +
			"o=- 3366701332 3366701332 IN IP6 dead:beef::666\r\n" +
			"s=-\r\n" +
			"c=IN IP6 dead:beef::666\r\n" +
			"t=0 0\r\n" +
			"m=audio 32898 RTP/AVP 9 18 0 101\r\n" +
			"a=rtpmap:101 telephone-event/8000\r\n" +
			"a=ptime:20\r\n",
		s2: "v=0\r\n" +
			"o=- 3366701332 3366701332 IN IP6 dead:beef::666\r\n" +
			"s=-\r\n" +
			"c=IN IP6 dead:beef::666\r\n" +
			"t=0 0\r\n" +
			"m=audio 32898 RTP/AVP 9 18 0 101\r\n" +
			"a=rtpmap:9 G722/8000\r\n" +
			"a=rtpmap:18 G729/8000\r\n" +
			"a=rtpmap:0 PCMU/8000\r\n" +
			"a=rtpmap:101 telephone-event/8000\r\n" +
			"a=ptime:20\r\n" +
			"a=sendrecv\r\n",
		sdp: &sdp.SDP{
			Origin: sdp.Origin{
				User:    "-",
				ID:      "3366701332",
				Version: "3366701332",
				Addr:    "dead:beef::666",
			},
			Session: "-",
			Time:    "0 0",
			Addr:    "dead:beef::666",
			Audio: &sdp.Media{
				Proto: "RTP/AVP",
				Port:  32898,
				Codecs: []sdp.Codec{
					{PT: 9, Name: "G722", Rate: 8000},
					{PT: 18, Name: "G729", Rate: 8000},
					{PT: 0, Name: "PCMU", Rate: 8000},
					{PT: 101, Name: "telephone-event", Rate: 8000},
				},
			},
			Ptime: 20,
			Attrs: [][2]string{},
		},
	},

	{
		name: "pjmedia long sdp is long",
		s: ("v=0\r\n" +
			"o=- 3457169218 3457169218 IN IP4 10.11.34.37\r\n" +
			"s=pjmedia\r\n" +
			"c=IN IP4 10.11.34.37\r\n" +
			"t=0 0\r\n" +
			"m=audio 4000 RTP/AVP 103 102 104 113 3 0 8 9 101\r\n" +
			"a=rtpmap:103 speex/16000\r\n" +
			"a=rtpmap:102 speex/8000\r\n" +
			"a=rtpmap:104 speex/32000\r\n" +
			"a=rtpmap:113 iLBC/8000\r\n" +
			"a=fmtp:113 mode=30\r\n" +
			"a=rtpmap:3 GSM/8000\r\n" +
			"a=rtpmap:0 PCMU/8000\r\n" +
			"a=rtpmap:8 PCMA/8000\r\n" +
			"a=rtpmap:9 G722/8000\r\n" +
			"a=rtpmap:101 telephone-event/8000\r\n" +
			"a=fmtp:101 0-15\r\n" +
			"a=rtcp:4001 IN IP4 10.11.34.37\r\n" +
			"a=X-nat:0\r\n" +
			"a=ptime:20\r\n" +
			"a=sendrecv\r\n"),
		sdp: &sdp.SDP{
			Origin: sdp.Origin{
				User:    "-",
				ID:      "3457169218",
				Version: "3457169218",
				Addr:    "10.11.34.37",
			},
			Session: "pjmedia",
			Time:    "0 0",
			Addr:    "10.11.34.37",
			Audio: &sdp.Media{
				Proto: "RTP/AVP",
				Port:  4000,
				Codecs: []sdp.Codec{
					{PT: 103, Name: "speex", Rate: 16000},
					{PT: 102, Name: "speex", Rate: 8000},
					{PT: 104, Name: "speex", Rate: 32000},
					{PT: 113, Name: "iLBC", Rate: 8000, Fmtp: "mode=30"},
					{PT: 3, Name: "GSM", Rate: 8000},
					{PT: 0, Name: "PCMU", Rate: 8000},
					{PT: 8, Name: "PCMA", Rate: 8000},
					{PT: 9, Name: "G722", Rate: 8000},
					{PT: 101, Name: "telephone-event", Rate: 8000, Fmtp: "0-15"},
				},
			},
			Ptime: 20,
			Attrs: [][2]string{
				{"rtcp", "4001 IN IP4 10.11.34.37"},
				{"X-nat", "0"},
			},
		},
	},

	{
		name: "mp3 tcp",
		s: ("v=0\r\n" +
			"o=- 3366701332 3366701334 IN IP4 10.11.34.37\r\n" +
			"s=squigglies\r\n" +
			"c=IN IP6 dead:beef::666\r\n" +
			"t=0 0\r\n" +
			"m=audio 80 TCP/IP 111\r\n" +
			"a=rtpmap:111 MP3/44100/2\r\n" +
			"a=sendonly\r\n"),
		sdp: &sdp.SDP{
			Origin: sdp.Origin{
				User:    "-",
				ID:      "3366701332",
				Version: "3366701334",
				Addr:    "10.11.34.37",
			},
			Session:  "squigglies",
			Time:     "0 0",
			Addr:     "dead:beef::666",
			SendOnly: true,
			Audio: &sdp.Media{
				Proto: "TCP/IP",
				Port:  80,
				Codecs: []sdp.Codec{
					{PT: 111, Name: "MP3", Rate: 44100, Param: "2"},
				},
			},
			Attrs: [][2]string{},
		},
	},
}

func sdpCompareCodec(t *testing.T, name string, correct, codec *sdp.Codec) {
	if correct != nil && codec == nil {
		t.Error(name, "not found")
	}
	if correct == nil && codec != nil {
		t.Error(name, "DO NOT WANT", codec)
	}
	if codec == nil {
		return
	}

	if correct.PT != codec.PT {
		t.Error(name, "PT", correct.PT, "!=", codec.PT)
	}
	if correct.Name != codec.Name {
		t.Error(name, "Name", correct.Name, "!=", codec.Name)
	}
	if correct.Rate != codec.Rate {
		t.Error(name, "Rate", correct.Rate, "!=", codec.Rate)
	}
	if correct.Param != codec.Param {
		t.Error(name, "Param", correct.Param, "!=", codec.Param)
	}
	if correct.Fmtp != codec.Fmtp {
		t.Error(name, "Fmtp", correct.Fmtp, "!=", codec.Fmtp)
	}
}

func sdpCompareCodecs(t *testing.T, name string, corrects, codecs []sdp.Codec) {
	if corrects != nil && codecs == nil {
		t.Error(name, "codecs not found")
	}
	if corrects == nil && codecs != nil {
		t.Error(name, "codecs WUT", codecs)
	}
	if corrects == nil || codecs == nil {
		return
	}

	if len(corrects) != len(codecs) {
		t.Error(name, "len(Codecs)", len(corrects), "!=", len(codecs))
	} else {
		for i := range corrects {
			c1, c2 := &corrects[i], &codecs[i]
			if c1 == nil || c2 == nil {
				t.Error(name, "where my codecs at?")
			} else {
				sdpCompareCodec(t, name, c1, c2)
			}
		}
	}
}

func sdpCompareMedia(t *testing.T, name string, correct, media *sdp.Media) {
	if correct != nil && media == nil {
		t.Error(name, "not found")
	}
	if correct == nil && media != nil {
		t.Error(name, "DO NOT WANT", media)
	}
	if correct == nil || media == nil {
		return
	}

	if correct.Proto != media.Proto {
		t.Error(name, "Proto", correct.Proto, "!=", media.Proto)
	}
	if correct.Port != media.Port {
		t.Error(name, "Port", correct.Port, "!=", media.Port)
	}
	if media.Codecs == nil || len(media.Codecs) < 1 {
		t.Error(name, "Must have at least one codec")
	}

	sdpCompareCodecs(t, name, correct.Codecs, media.Codecs)
}

func TestParse(t *testing.T) {
	for _, test := range sdpTests {
		sdp, err := sdp.Parse(test.s)
		if err != nil {
			if test.err == nil {
				t.Errorf("%v", err)
				continue
			} else { // test was supposed to fail
				panic("todo")
			}
		}

		if test.sdp.Addr != sdp.Addr {
			t.Error(test.name, "Addr", test.sdp.Addr, "!=", sdp.Addr)
		}
		if test.sdp.Origin.User != sdp.Origin.User {
			t.Error(test.name, "Origin.User", test.sdp.Origin.User, "!=",
				sdp.Origin.User)
		}
		if test.sdp.Origin.ID != sdp.Origin.ID {
			t.Error(test.name, "Origin.ID doesn't match")
		}
		if test.sdp.Origin.Version != sdp.Origin.Version {
			t.Error(test.name, "Origin.Version doesn't match")
		}
		if test.sdp.Origin.Addr != sdp.Origin.Addr {
			t.Error(test.name, "Origin.Addr doesn't match")
		}
		if test.sdp.Session != sdp.Session {
			t.Error(test.name, "Session", test.sdp.Session, "!=", sdp.Session)
		}
		if test.sdp.Time != sdp.Time {
			t.Error(test.name, "Time", test.sdp.Time, "!=", sdp.Time)
		}
		if test.sdp.Ptime != sdp.Ptime {
			t.Error(test.name, "Ptime", test.sdp.Ptime, "!=", sdp.Ptime)
		}
		if test.sdp.RecvOnly != sdp.RecvOnly {
			t.Error(test.name, "RecvOnly doesn't match")
		}
		if test.sdp.SendOnly != sdp.SendOnly {
			t.Error(test.name, "SendOnly doesn't match")
		}

		if test.sdp.Attrs != nil {
			if sdp.Attrs == nil {
				t.Error(test.name, "Attrs weren't extracted")
			} else if len(sdp.Attrs) != len(test.sdp.Attrs) {
				t.Error(test.name, "Attrs length not same")
			} else {
				for i := range sdp.Attrs {
					p1, p2 := test.sdp.Attrs[i], sdp.Attrs[i]
					if p1[0] != p2[0] || p1[1] != p2[1] {
						t.Error(test.name, "attr", p1, "!=", p2)
						break
					}
				}
			}
		} else {
			if sdp.Attrs != nil {
				t.Error(test.name, "unexpected attrs", sdp.Attrs)
			}
		}

		sdpCompareMedia(t, "Audio", test.sdp.Audio, sdp.Audio)
		sdpCompareMedia(t, "Video", test.sdp.Video, sdp.Video)
	}
}

func TestFormatSDP(t *testing.T) {
	for _, test := range sdpTests {
		sdp := test.sdp.String()
		s := test.s
		if test.s2 != "" {
			s = test.s2
		}
		if s != sdp {
			t.Error("\n" + test.name + "\n\n" + s + "\nIS NOT\n\n" + sdp)
			fmt.Printf("%s", sdp)
			fmt.Printf("pokémon")
		}
	}
}
