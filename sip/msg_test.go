package sip_test

import (
	"github.com/jart/gosip/sip"
	"reflect"
	"testing"
)

type msgTest struct {
	name string
	s    string
	msg  sip.Msg
	err  error
}

var msgTests = []msgTest{

	msgTest{
		name: "UTF8 Phrase",
		s: "SIP/2.0 200 ◕◡◕\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "◕◡◕",
		},
	},

	msgTest{
		name: "Left Padding",
		s: "SIP/2.0 200 OK\r\n" +
			"Expires:    666\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			Expires:      666,
		},
	},

	msgTest{
		name: "Extension Headers",
		s: "SIP/2.0 200 OK\r\n" +
			"X-LOL: omfg\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			Headers: sip.Headers{
				"X-LOL": "omfg",
			},
		},
	},

	msgTest{
		name: "Multiple Addresses",
		s: "SIP/2.0 200 OK\r\n" +
			"From:  <sip:lol.com> , <sip:bog.com> \r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			From: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "lol.com",
				},
				Next: &sip.Addr{
					Uri: &sip.URI{
						Scheme: "sip",
						Host:   "bog.com",
					},
				},
			},
		},
	},

	msgTest{
		name: "Line Continuation Warning",
		s: "SIP/2.0 200 OK\r\n" +
			"Warning: Morning and evening\r\n" +
			" Maids heard the goblins cry:\r\n" +
			" “Come buy our orchard fruits,\r\n" +
			" Come buy, come buy:\r\n" +
			" Apples and quinces,\r\n" +
			" Lemons and oranges\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			Warning: "Morning and evening\r\n" +
				" Maids heard the goblins cry:\r\n" +
				" “Come buy our orchard fruits,\r\n" +
				" Come buy, come buy:\r\n" +
				" Apples and quinces,\r\n" +
				" Lemons and oranges",
		},
	},

	msgTest{
		name: "Line Continuation Warning Followed By Extended Header",
		s: "SIP/2.0 200 OK\r\n" +
			"Warning: Morning and evening\r\n" +
			" Maids heard the goblins cry:\r\n" +
			" “Come buy our orchard fruits,\r\n" +
			" Come buy, come buy:\r\n" +
			" Apples and quinces,\r\n" +
			" Lemons and oranges\r\n" +
			"X-LOL: omfg\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			Warning: "Morning and evening\r\n" +
				" Maids heard the goblins cry:\r\n" +
				" “Come buy our orchard fruits,\r\n" +
				" Come buy, come buy:\r\n" +
				" Apples and quinces,\r\n" +
				" Lemons and oranges",
			Headers: sip.Headers{
				"X-LOL": "omfg",
			},
		},
	},

	msgTest{
		name: "Line Continuation Extended Followed By Extended",
		s: "SIP/2.0 200 OK\r\n" +
			"X-Warning: Come buy our orchard fruits,\r\n" +
			" Come buy, come buy\r\n" +
			"X-LOL: omfg\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			Headers: sip.Headers{
				"X-Warning": "Come buy our orchard fruits,\r\n" +
					" Come buy, come buy",
				"X-LOL": "omfg",
			},
		},
	},

	msgTest{
		name: "Line Continuation Extended Followed By Extended 2",
		s: "SIP/2.0 200 OK\r\n" +
			"NewFangledHeader:   newfangled value\r\n" +
			" continued newfangled value\r\n" +
			"UnknownHeaderWithUnusualValue: ;;,,;;,;\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			Headers: sip.Headers{
				"NewFangledHeader": "newfangled value\r\n" +
					" continued newfangled value",
				"UnknownHeaderWithUnusualValue": ";;,,;;,;",
			},
		},
	},

	msgTest{
		name: "Line Continuations Addr",
		s: "SIP/2.0 200 OK\r\n" +
			"From:\r\n" +
			" <sip:lol.com>,\r\n" +
			" <sip:bog.com>\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			From: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "lol.com",
				},
				Next: &sip.Addr{
					Uri: &sip.URI{
						Scheme: "sip",
						Host:   "bog.com",
					},
				},
			},
		},
	},

	msgTest{
		name: "Extended header looks like standard headers",
		s: "SIP/2.0 200 OK\r\n" +
			"Proxy-LOL: take\r\n" +
			"CSeq2: back\r\n" +
			"Contazt: the\r\n" +
			"P-Asserted-LOL: dance\r\n" +
			"viaz: floor\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			Headers: sip.Headers{
				"Proxy-LOL":      "take",
				"CSeq2":          "back",
				"Contazt":        "the",
				"P-Asserted-LOL": "dance",
				"viaz":           "floor",
			},
		},
	},

	msgTest{
		name: "Address Unquoted Display",
		s: "SIP/2.0 200 OK\r\n" +
			"From: Kitty <sip:lol.com>\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			From: &sip.Addr{
				Display: "Kitty",
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "lol.com",
				},
			},
		},
	},

	msgTest{
		name: "Address Quoted Display",
		s: "SIP/2.0 200 OK\r\n" +
			"From: \"Hello \\\"Kitty\\\" ◕◡◕\" <sip:lol.com>\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			From: &sip.Addr{
				Display: "Hello \"Kitty\" ◕◡◕",
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "lol.com",
				},
			},
		},
	},

	msgTest{
		name: "Address Quoted Display Multiline",
		s: "SIP/2.0 200 OK\r\n" +
			"From: \"oh\r\n" +
			"  my \r\n" +
			" goth\" <sip:lol.com>\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			From: &sip.Addr{
				Display: "oh\r\n" +
					"  my \r\n" +
					" goth",
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "lol.com",
				},
			},
		},
	},

	msgTest{
		name: "Address Unquoted Display Multiline",
		s: "SIP/2.0 200 OK\r\n" +
			"From: oh\r\n" +
			"  my \r\n" +
			" goth <sip:lol.com>\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			From: &sip.Addr{
				Display: "oh\r\n" +
					"  my \r\n" +
					" goth",
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "lol.com",
				},
			},
		},
	},

	msgTest{
		name: "Addr Tag",
		s: "SIP/2.0 200 OK\r\n" +
			"From: <sip:lol.com>;tag=omfg\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			From: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "lol.com",
				},
				Params: sip.Params{
					"tag": "omfg",
				},
			},
		},
	},

	msgTest{
		name: "Addr Tag Quoted",
		// TODO(jart): Crash when extra spacing in here.
		s: "SIP/2.0 200 OK\r\n" +
			"From: <sip:lol.com>;tag=\"◕◡◕\"\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			From: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "lol.com",
				},
				Params: sip.Params{
					"tag": "◕◡◕",
				},
			},
		},
	},

	msgTest{
		name: "Addr Tag Bare",
		s: "SIP/2.0 200 OK\r\n" +
			"From: <sip:lol.com>;tag\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			From: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "lol.com",
				},
				Params: sip.Params{
					"tag": "",
				},
			},
		},
	},

	msgTest{
		name: "Missing Angle Brackets With Tag Belongs to Addr Not URI",
		s: "SIP/2.0 200 OK\r\n" +
			"From: sip:lol.com;tag=omfg\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			From: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "lol.com",
				},
				Params: sip.Params{
					"tag": "omfg",
				},
			},
		},
	},

	// // TODO(jart): Implement me.
	// msgTest{
	// 	name: "Content Type Params",
	// 	s: "SIP/2.0 200 OK\r\n" +
	// 		"Content-Type: multipart/signed;\r\n" +
	// 		"        protocol=\"application/pkcs7-signature\";\r\n" +
	// 		"        micalg=sha1; boundary=boundary42\r\n" +
	// 		"\r\n",
	// 	msg: sip.Msg{
	// 		VersionMajor: 2,
	// 		Status:       200,
	// 		Phrase:       "OK",
	// 		Expires:      666,
	// 	},
	// },

	msgTest{
		name: "Via Line Continuation",
		s: "SIP/2.0 200 OK\r\n" +
			"Via: SIP/2.0/UDP 10.11.34.37 ,\r\n" +
			"     SIP/2.0/UDP 10.11.34.38\r\n" +
			"Warning: Maids heard the goblins cry\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			Warning:      "Maids heard the goblins cry",
			Via: &sip.Via{
				Protocol:  "SIP",
				Version:   "2.0",
				Transport: "UDP",
				Host:      "10.11.34.37",
				Next: &sip.Via{
					Protocol:  "SIP",
					Version:   "2.0",
					Transport: "UDP",
					Host:      "10.11.34.38",
				},
			},
		},
	},

	msgTest{
		name: "Via Multiple Lines",
		s: "SIP/2.0 200 OK\r\n" +
			"Via: SIP/2.0/UDP 10.11.34.37\r\n" +
			"Via: SIP/2.0/UDP 10.11.34.38\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			Via: &sip.Via{
				Protocol:  "SIP",
				Version:   "2.0",
				Transport: "UDP",
				Host:      "10.11.34.37",
				Next: &sip.Via{
					Protocol:  "SIP",
					Version:   "2.0",
					Transport: "UDP",
					Host:      "10.11.34.38",
				},
			},
		},
	},

	msgTest{
		name: "Via Multiple Lines Continuation",
		s: "SIP/2.0 200 OK\r\n" +
			"Via: SIP/2.0/UDP 10.11.34.37\r\n" +
			"v: SIP/2.0/UDP 10.11.34.38, SIP/2.0/UDP 10.11.34.39\r\n" +
			"m: <sip:love.com>\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			Via: &sip.Via{
				Protocol:  "SIP",
				Version:   "2.0",
				Transport: "UDP",
				Host:      "10.11.34.37",
				Next: &sip.Via{
					Protocol:  "SIP",
					Version:   "2.0",
					Transport: "UDP",
					Host:      "10.11.34.38",
					Next: &sip.Via{
						Protocol:  "SIP",
						Version:   "2.0",
						Transport: "UDP",
						Host:      "10.11.34.39",
					},
				},
			},
			Contact: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "love.com",
				},
			},
		},
	},

	msgTest{
		name: "Via Param",
		s: "SIP/2.0 200 OK\r\n" +
			"Via: SIP/ 2.0/TCP spindle.example.com ;branch=z9hG4bK9ikj8\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			Via: &sip.Via{
				Protocol:  "SIP",
				Version:   "2.0",
				Transport: "TCP",
				Host:      "spindle.example.com",
				Params:    sip.Params{"branch": "z9hG4bK9ikj8"},
			},
		},
	},

	msgTest{
		name: "Via Param Torture",
		s: "SIP/2.0 200 OK\r\n" +
			"v:  SIP  / 2.0  / TCP     spindle.example.com   ;\r\n" +
			"  branch  =   z9hG4bK9ikj8\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			Via: &sip.Via{
				Protocol:  "SIP",
				Version:   "2.0",
				Transport: "TCP",
				Host:      "spindle.example.com",
				Params:    sip.Params{"branch": "z9hG4bK9ikj8"},
			},
		},
	},

	msgTest{
		name: "Via Torture",
		s: "SIP/2.0 200 OK\r\n" +
			"Via  : SIP  /   2.0\r\n" +
			" /UDP\r\n" +
			"    192.0.2.2;branch=390skdjuw\r\n" +
			"v:  SIP  / 2.0  / TCP     spindle.example.com   ;\r\n" +
			"  branch  =   z9hG4bK9ikj8  ,\r\n" +
			" SIP  /    2.0   / UDP  192.168.255.111   ; branch=\r\n" +
			" z9hG4bK30239\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			Via: &sip.Via{
				Protocol:  "SIP",
				Version:   "2.0",
				Transport: "UDP",
				Host:      "192.0.2.2",
				Params:    sip.Params{"branch": "390skdjuw"},
				Next: &sip.Via{
					Protocol:  "SIP",
					Version:   "2.0",
					Transport: "TCP",
					Host:      "spindle.example.com",
					Params:    sip.Params{"branch": "z9hG4bK9ikj8"},
					Next: &sip.Via{
						Protocol:  "SIP",
						Version:   "2.0",
						Transport: "UDP",
						Host:      "192.168.255.111",
						Params:    sip.Params{"branch": "z9hG4bK30239"},
					},
				},
			},
		},
	},

	msgTest{
		name: "OPTIONS",
		s: "OPTIONS sip:10.11.34.37:42367 SIP/2.0\r\n" +
			"Via: SIP/2.0/UDP 10.11.34.37:42367;rport;branch=9dc39c3c3e84\r\n" +
			"Max-Forwards: 60\r\n" +
			"To: <sip:10.11.34.37:5060>\r\n" +
			"From: <sip:10.11.34.37:42367;laffo>;tag=11917cbc0513\r\n" +
			"Call-ID: e71a163e-c440-474d-a4ec-5cd85a0309c6\r\n" +
			"CSeq: 36612 OPTIONS\r\n" +
			"Contact: <sip:10.11.34.37:42367>\r\n" +
			"User-Agent: ghoul/0.1\r\n" +
			"Accept: application/sdp\r\n" +
			"Content-Length: 0\r\n" +
			"\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Method:       "OPTIONS",
			CSeqMethod:   "OPTIONS",
			MaxForwards:  60,
			CallID:       "e71a163e-c440-474d-a4ec-5cd85a0309c6",
			CSeq:         36612,
			Request: &sip.URI{
				Scheme: "sip",
				Host:   "10.11.34.37",
				Port:   42367,
			},
			Via: &sip.Via{
				Protocol:  "SIP",
				Version:   "2.0",
				Transport: "UDP",
				Host:      "10.11.34.37",
				Port:      42367,
				Params:    sip.Params{"rport": "", "branch": "9dc39c3c3e84"},
			},
			To: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "10.11.34.37",
					Port:   5060,
				},
			},
			From: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "10.11.34.37",
					Port:   42367,
					Params: sip.Params{"laffo": ""},
				},
				Params: sip.Params{"tag": "11917cbc0513"},
			},
			Contact: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "10.11.34.37",
					Port:   42367,
				},
			},
			UserAgent: "ghoul/0.1",
			Accept:    "application/sdp",
		},
	},

	msgTest{
		s: "SIP/2.0 200 OK\r\n" +
			"Via: SIP/2.0/UDP 127.0.0.1:52711;branch=z9hG4bK-03d1d81e94a0;received=127.0.0.1;rport=52711\r\n" +
			"From: <sip:127.0.0.1:52711;transport=udp>;tag=4568e274dbd8\r\n" +
			"To: <sip:echo@127.0.0.1:5060>;tag=as71a0fa72\r\n" +
			"Call-ID: 99042736-d40b-4d96-a81b-867321443ff5\r\n" +
			"CSeq: 16378 INVITE\r\n" +
			"Server: Asterisk PBX 10.11.1\r\n" +
			"Allow: INVITE, ACK, CANCEL, OPTIONS, BYE, REFER, SUBSCRIBE, NOTIFY, INFO, PUBLISH\r\n" +
			"Supported: replaces, timer\r\n" +
			"Contact: <sip:echo@127.0.0.1:5060>\r\n" +
			"Content-Type: application/sdp-lol\r\n" +
			"Content-Length: 255\r\n" +
			"\r\n" +
			"v=0\r\n" +
			"o=root 736606944 736606944 IN IP4 127.0.0.1\r\n" +
			"s=Asterisk PBX 10.11.1\r\n" +
			"c=IN IP4 127.0.0.1\r\n" +
			"t=0 0\r\n" +
			"m=audio 23452 RTP/AVP 0 101\r\n" +
			"a=rtpmap:0 PCMU/8000\r\n" +
			"a=rtpmap:101 telephone-event/8000\r\n" +
			"a=fmtp:101 0-16\r\n" +
			"a=silenceSupp:off - - - -\r\n" +
			"a=ptime:20\r\n" +
			"a=sendrecv\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			CallID:       "99042736-d40b-4d96-a81b-867321443ff5",
			CSeq:         16378,
			CSeqMethod:   "INVITE",
			Server:       "Asterisk PBX 10.11.1",
			Allow:        "INVITE, ACK, CANCEL, OPTIONS, BYE, REFER, SUBSCRIBE, NOTIFY, INFO, PUBLISH",
			Supported:    "replaces, timer",
			Via: &sip.Via{
				Protocol:  "SIP",
				Version:   "2.0",
				Transport: "UDP",
				Host:      "127.0.0.1",
				Port:      52711,
				Params: sip.Params{
					"branch":   "z9hG4bK-03d1d81e94a0",
					"received": "127.0.0.1",
					"rport":    "52711",
				},
			},
			From: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "127.0.0.1",
					Port:   52711,
					Params: sip.Params{"transport": "udp"},
				},
				Params: sip.Params{"tag": "4568e274dbd8"},
			},
			To: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					User:   "echo",
					Host:   "127.0.0.1",
					Port:   5060,
				},
				Params: sip.Params{"tag": "as71a0fa72"},
			},
			Contact: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					User:   "echo",
					Host:   "127.0.0.1",
					Port:   5060,
				},
			},
			Payload: &sip.MiscPayload{
				T: "application/sdp-lol",
				D: []byte("v=0\r\n" +
					"o=root 736606944 736606944 IN IP4 127.0.0.1\r\n" +
					"s=Asterisk PBX 10.11.1\r\n" +
					"c=IN IP4 127.0.0.1\r\n" +
					"t=0 0\r\n" +
					"m=audio 23452 RTP/AVP 0 101\r\n" +
					"a=rtpmap:0 PCMU/8000\r\n" +
					"a=rtpmap:101 telephone-event/8000\r\n" +
					"a=fmtp:101 0-16\r\n" +
					"a=silenceSupp:off - - - -\r\n" +
					"a=ptime:20\r\n" +
					"a=sendrecv\r\n"),
			},
		},
	},

	msgTest{
		s: "SIP/2.0 200 OK\r\n" +
			"Via: SIP/2.0/UDP 1.2.3.4:55345;branch=z9hG4bK-d1d81e94a099\r\n" +
			"From: <sip:+12126660420@fl.gg>;tag=68e274dbd83b\r\n" +
			"To: <sip:+12125650666@fl.gg>;tag=gK0cacc73a\r\n" +
			"Call-ID: 042736d4-0bd9-4681-ab86-7321443ff58a\r\n" +
			"CSeq: 31109 INVITE\r\n" +
			"Record-Route: <sip:216.115.69.133:5060;lr>\r\n" +
			"Record-Route: <sip:216.115.69.144:5060;lr>\r\n" +
			"Contact: <sip:+12125650666@4.55.22.99:5060>\r\n" +
			"Content-Type: application/sdp-lol\r\n" +
			"Content-Length:  168\r\n" +
			"\r\n" +
			"v=0\r\n" +
			"o=- 24294 7759 IN IP4 4.55.22.66\r\n" +
			"s=-\r\n" +
			"c=IN IP4 4.55.22.66\r\n" +
			"t=0 0\r\n" +
			"m=audio 19580 RTP/AVP 0 101\r\n" +
			"a=rtpmap:101 telephone-event/8000\r\n" +
			"a=fmtp:101 0-15\r\n" +
			"a=maxptime:20\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Status:       200,
			Phrase:       "OK",
			CallID:       "042736d4-0bd9-4681-ab86-7321443ff58a",
			CSeq:         31109,
			CSeqMethod:   "INVITE",
			Via: &sip.Via{
				Protocol:  "SIP",
				Version:   "2.0",
				Transport: "UDP",
				Host:      "1.2.3.4",
				Port:      55345,
				Params:    sip.Params{"branch": "z9hG4bK-d1d81e94a099"},
			},
			From: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					User:   "+12126660420",
					Host:   "fl.gg",
				},
				Params: sip.Params{"tag": "68e274dbd83b"},
			},
			To: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					User:   "+12125650666",
					Host:   "fl.gg",
				},
				Params: sip.Params{"tag": "gK0cacc73a"},
			},
			Contact: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					User:   "+12125650666",
					Host:   "4.55.22.99",
					Port:   5060,
				},
			},
			RecordRoute: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "216.115.69.133",
					Port:   5060,
					Params: sip.Params{"lr": ""},
				},
				Next: &sip.Addr{
					Uri: &sip.URI{
						Scheme: "sip",
						Host:   "216.115.69.144",
						Port:   5060,
						Params: sip.Params{"lr": ""},
					},
				},
			},
			Payload: &sip.MiscPayload{
				T: "application/sdp-lol",
				D: []byte("v=0\r\n" +
					"o=- 24294 7759 IN IP4 4.55.22.66\r\n" +
					"s=-\r\n" +
					"c=IN IP4 4.55.22.66\r\n" +
					"t=0 0\r\n" +
					"m=audio 19580 RTP/AVP 0 101\r\n" +
					"a=rtpmap:101 telephone-event/8000\r\n" +
					"a=fmtp:101 0-15\r\n" +
					"a=maxptime:20\r\n"),
			},
		},
	},

	msgTest{
		name: "INVITE",
		s: "INVITE sip:10.11.34.37 SIP/2.0\r\n" +
			"via: SIP/2.0/UDP 10.11.34.37:59516;rport;branch=z9hG4bKS308QB9UUpNyD\r\n" +
			"Max-Forwards: 70\r\n" +
			"From: <sip:10.11.34.37:59516>;tag=S1jX7UtK5Zerg\r\n" +
			"To: <sip:10.11.34.37>\r\n" +
			"Call-ID: 87704115-03b8-122e-08b5-001bfcce6bdf\r\n" +
			"CSeq: 133097268 INVITE\r\n" +
			// "Contact: <sip:10.11.34.37:59516>,\r\n" +
			// "         <sip:10.11.34.38:59516>\r\n" +
			"Contact: <sip:10.11.34.37:59516>\r\n" +
			"User-Agent: tube/0.1\r\n" +
			"Allow: INVITE, ACK, BYE, CANCEL, OPTIONS, PRACK, MESSAGE, SUBSCRIBE, NOTIFY, REFER, UPDATE, INFO\r\n" +
			"Supported: timer, 100rel\r\n" +
			"Allow-Events: talk\r\n" +
			"Content-Type: application/sdp-lol\r\n" +
			"Content-Disposition: session\r\n" +
			"Content-Length: 218\r\n" +
			"\r\n" +
			"v=0\r\n" +
			"o=- 2862054018559638081 6057228511765453924 IN IP4 10.11.34.37\r\n" +
			"s=-\r\n" +
			"c=IN IP4 10.11.34.37\r\n" +
			"t=0 0\r\n" +
			"m=audio 23448 RTP/AVP 0 101\r\n" +
			"a=rtpmap:0 PCMU/8000\r\n" +
			"a=rtpmap:101 telephone-event/8000\r\n" +
			"a=fmtp:101 0-16\r\n" +
			"a=ptime:20\r\n",
		msg: sip.Msg{
			VersionMajor: 2,
			Method:       "INVITE",
			CSeqMethod:   "INVITE",
			MaxForwards:  70,
			CallID:       "87704115-03b8-122e-08b5-001bfcce6bdf",
			CSeq:         133097268,
			Request: &sip.URI{
				Scheme: "sip",
				Host:   "10.11.34.37",
			},
			Via: &sip.Via{
				Protocol:  "SIP",
				Version:   "2.0",
				Transport: "UDP",
				Host:      "10.11.34.37",
				Port:      59516,
				Params:    sip.Params{"rport": "", "branch": "z9hG4bKS308QB9UUpNyD"},
			},
			To: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "10.11.34.37",
				},
			},
			From: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "10.11.34.37",
					Port:   59516,
				},
				Params: sip.Params{"tag": "S1jX7UtK5Zerg"},
			},
			Contact: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "10.11.34.37",
					Port:   59516,
				},
				// Next: &sip.Addr{
				// 	Uri: &sip.URI{
				// 		Scheme: "sip",
				// 		Host:   "10.11.34.38",
				// 		Port:   59516,
				// 	},
				// },
			},
			UserAgent:          "tube/0.1",
			Allow:              "INVITE, ACK, BYE, CANCEL, OPTIONS, PRACK, MESSAGE, SUBSCRIBE, NOTIFY, REFER, UPDATE, INFO",
			AllowEvents:        "talk",
			ContentDisposition: "session",
			Supported:          "timer, 100rel",
			Payload: &sip.MiscPayload{
				T: "application/sdp-lol",
				D: []byte("v=0\r\n" +
					"o=- 2862054018559638081 6057228511765453924 IN IP4 10.11.34.37\r\n" +
					"s=-\r\n" +
					"c=IN IP4 10.11.34.37\r\n" +
					"t=0 0\r\n" +
					"m=audio 23448 RTP/AVP 0 101\r\n" +
					"a=rtpmap:0 PCMU/8000\r\n" +
					"a=rtpmap:101 telephone-event/8000\r\n" +
					"a=fmtp:101 0-16\r\n" +
					"a=ptime:20\r\n"),
			},
		},
	},

	msgTest{
		name: "RFC4475 Torture Message #1",
		s: "INVITE sip:vivekg@chair-dnrc.example.com;unknownparam SIP/2.0\r\n" +
			"TO :\r\n" +
			" sip:vivekg@chair-dnrc.example.com ;   tag    = 1918181833n\r\n" +
			"from   : \"J Rosenberg \\\\\\\"\"       <sip:jdrosen@example.com>\r\n" +
			"  ;\r\n" +
			"  tag = 98asjd8\r\n" +
			"MaX-fOrWaRdS: 0068\r\n" +
			"Call-ID: wsinv.ndaksdj@192.0.2.1\r\n" +
			"Content-Length   : 150\r\n" +
			"cseq: 0009\r\n" +
			"  INVITE\r\n" +
			"Via  : SIP  /   2.0\r\n" +
			" /UDP\r\n" +
			"    192.0.2.2;branch=390skdjuw\r\n" +
			"s :\r\n" +
			"NewFangledHeader:   newfangled value\r\n" +
			" continued newfangled value\r\n" +
			"UnknownHeaderWithUnusualValue: ;;,,;;,;\r\n" +
			"Content-Type: application/sdp-JART\r\n" +
			"Route:\r\n" +
			" <sip:services.example.com;lr;unknownwith=value;unknown-no-value>\r\n" +
			"v:  SIP  / 2.0  / TCP     spindle.example.com   ;\r\n" +
			"  branch  =   z9hG4bK9ikj8  ,\r\n" +
			" SIP  /    2.0   / UDP  192.168.255.111   ; branch=\r\n" +
			" z9hG4bK30239\r\n" +
			"m:\"Quoted string \\\"\\\"\" <sip:jdrosen@example.com> ; newparam =\r\n" +
			"      newvalue ;\r\n" +
			"  secondparam ; q = 0.33\r\n" +
			"\r\n" +
			"v=0\r\n" +
			"o=mhandley 29739 7272939 IN IP4 192.0.2.3\r\n" +
			"s=-\r\n" +
			"c=IN IP4 192.0.2.4\r\n" +
			"t=0 0\r\n" +
			"m=audio 49217 RTP/AVP 0 12\r\n" +
			"m=video 3227 RTP/AVP 31\r\n" +
			"a=rtpmap:31 LPC\r\n",
		msg: sip.Msg{
			Method:       "INVITE",
			VersionMajor: 2,
			Request: &sip.URI{
				Scheme: "sip",
				User:   "vivekg",
				Host:   "chair-dnrc.example.com",
				Params: sip.Params{"unknownparam": ""},
			},
			To: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					User:   "vivekg",
					Host:   "chair-dnrc.example.com",
				},
				Params: sip.Params{"tag": "1918181833n"},
			},
			From: &sip.Addr{
				Display: "J Rosenberg \\\"",
				Uri: &sip.URI{
					Scheme: "sip",
					User:   "vivekg",
					Host:   "chair-dnrc.example.com",
					Params: sip.Params{"tag": "98asjd8"},
				},
			},
			MaxForwards: 68,
			CallID:      "wsinv.ndaksdj@192.0.2.1",
			CSeq:        9,
			CSeqMethod:  "INVITE",
			Via: &sip.Via{
				Protocol:  "SIP",
				Version:   "2.0",
				Transport: "UDP",
				Host:      "192.0.2.2",
				Params:    sip.Params{"branch": "390skdjuw"},
				Next: &sip.Via{
					Protocol:  "SIP",
					Version:   "2.0",
					Transport: "TCP",
					Host:      "spindle.example.com",
					Params:    sip.Params{"branch": "z9hG4bK9ikj8"},
					Next: &sip.Via{
						Protocol:  "SIP",
						Version:   "2.0",
						Transport: "UDP",
						Host:      "192.168.255.111",
						Params:    sip.Params{"branch": "z9hG4bK30239"},
					},
				},
			},
			Subject: "",
			Headers: sip.Headers{
				"NewFangledHeader": "newfangled value\r\n" +
					" continued newfangled value",
				"UnknownHeaderWithUnusualValue": ";;,,;;,;",
			},
			Route: &sip.Addr{
				Uri: &sip.URI{
					Scheme: "sip",
					Host:   "services.example.com",
					Params: sip.Params{
						"lr":               "",
						"unknownwith":      "value",
						"unknown-no-value": "",
					},
				},
			},
			Contact: &sip.Addr{
				Display: "Quoted string \"\"",
				Uri: &sip.URI{
					Scheme: "sip",
					User:   "jdrosen",
					Host:   "example.com",
				},
				Params: sip.Params{
					"newparam":    "newvalue",
					"secondparam": "",
					"q":           "0.33",
				},
			},
			Payload: &sip.MiscPayload{
				T: "application/sdp-JART",
				D: []byte("v=0\r\n" +
					"o=mhandley 29739 7272939 IN IP4 192.0.2.3\r\n" +
					"s=-\r\n" +
					"c=IN IP4 192.0.2.4\r\n" +
					"t=0 0\r\n" +
					"m=audio 49217 RTP/AVP 0 12\r\n" +
					"m=video 3227 RTP/AVP 31\r\n" +
					"a=rtpmap:31 LPC\r\n"),
			},
		},
	},
}

func TestParseMsg(t *testing.T) {
	for _, test := range msgTests {
		msg, err := sip.ParseMsg(test.s)
		if err != nil {
			if test.err == nil {
				t.Errorf("%v", err)
				continue
			} else { // test was supposed to fail
				panic("TODO(jart): Implement failing support.")
			}
		}
		if !reflect.DeepEqual(&test.msg, msg) {
			t.Errorf("%s:\n%#v !=\n%#v", test.name, &test.msg, msg)
			if !reflect.DeepEqual(test.msg.Payload, msg.Payload) {
				t.Errorf("Payload:\n%#v !=\n%#v", test.msg.Payload, msg.Payload)
			}
			if !reflect.DeepEqual(test.msg.Headers, msg.Headers) {
				t.Errorf("Headers:\n%#v !=\n%#v", test.msg.Headers, msg.Headers)
			}
			if !reflect.DeepEqual(test.msg.Via, msg.Via) {
				t.Errorf("Via:\n%#v !=\n%#v", test.msg.Via, msg.Via)
				t.Errorf("Via #2:\n%#v !=\n%#v", test.msg.Via.Next, msg.Via.Next)
				t.Errorf("Via #3:\n%#v !=\n%#v", test.msg.Via.Next.Next, msg.Via.Next.Next)
				// t.Errorf("Via #4:\n%#v !=\n%#v", test.msg.Via.Next.Next.Next, msg.Via.Next.Next.Next)
			}
			if !reflect.DeepEqual(test.msg.Request, msg.Request) {
				t.Errorf("Request:\n%#v !=\n%#v", test.msg.Request, msg.Request)
			}
			if !reflect.DeepEqual(test.msg.To, msg.To) {
				addrError(t, "To", test.msg.To, msg.To)
			}
			if !reflect.DeepEqual(test.msg.From, msg.From) {
				addrError(t, "From", test.msg.From, msg.From)
			}
			if !reflect.DeepEqual(test.msg.Contact, msg.Contact) {
				addrError(t, "Contact", test.msg.Contact, msg.Contact)
			}
			if !reflect.DeepEqual(test.msg.RecordRoute, msg.RecordRoute) {
				addrError(t, "RecordRoute", test.msg.RecordRoute, msg.RecordRoute)
			}
		}
	}
}
