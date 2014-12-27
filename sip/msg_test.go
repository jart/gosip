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
			Method:      "OPTIONS",
			CSeqMethod:  "OPTIONS",
			MaxForwards: 60,
			CallID:      "e71a163e-c440-474d-a4ec-5cd85a0309c6",
			CSeq:        36612,
			Request: &sip.URI{
				Scheme: "sip",
				Host:   "10.11.34.37",
				Port:   42367,
			},
			Via: &sip.Via{
				Version: "2.0",
				Proto:   "UDP",
				Host:    "10.11.34.37",
				Port:    42367,
				Params:  sip.Params{"rport": "", "branch": "9dc39c3c3e84"},
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
			Headers: map[string]string{
				"User-Agent": "ghoul/0.1",
				"Accept":     "application/sdp",
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
			"Contact: <sip:10.11.34.37:59516>\r\n" +
			"         <sip:10.11.34.38:59516>\r\n" +
			"User-Agent: tube/0.1\r\n" +
			"Allow: INVITE, ACK, BYE, CANCEL, OPTIONS, PRACK, MESSAGE, SUBSCRIBE, NOTIFY, REFER, UPDATE, INFO\r\n" +
			"Supported: timer, 100rel\r\n" +
			"Allow-Events: talk\r\n" +
			"Content-Type: application/sdp\r\n" +
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
			Method:      "INVITE",
			CSeqMethod:  "INVITE",
			MaxForwards: 70,
			CallID:      "87704115-03b8-122e-08b5-001bfcce6bdf",
			CSeq:        133097268,
			Request: &sip.URI{
				Scheme: "sip",
				Host:   "10.11.34.37",
			},
			Via: &sip.Via{
				Version: "2.0",
				Proto:   "UDP",
				Host:    "10.11.34.37",
				Port:    59516,
				Params:  sip.Params{"rport": "", "branch": "z9hG4bKS308QB9UUpNyD"},
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
				Next: &sip.Addr{
					Uri: &sip.URI{
						Scheme: "sip",
						Host:   "10.11.34.38",
						Port:   59516,
					},
				},
			},
			Headers: map[string]string{
				"User-Agent":          "tube/0.1",
				"Allow":               "INVITE, ACK, BYE, CANCEL, OPTIONS, PRACK, MESSAGE, SUBSCRIBE, NOTIFY, REFER, UPDATE, INFO",
				"Allow-Events":        "talk",
				"Content-Disposition": "session",
				"Supported":           "timer, 100rel",
				"Content-Type":        "application/sdp",
			},
			Payload: "v=0\r\n" +
				"o=- 2862054018559638081 6057228511765453924 IN IP4 10.11.34.37\r\n" +
				"s=-\r\n" +
				"c=IN IP4 10.11.34.37\r\n" +
				"t=0 0\r\n" +
				"m=audio 23448 RTP/AVP 0 101\r\n" +
				"a=rtpmap:0 PCMU/8000\r\n" +
				"a=rtpmap:101 telephone-event/8000\r\n" +
				"a=fmtp:101 0-16\r\n" +
				"a=ptime:20\r\n",
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
			t.Errorf("Message:\n%#v !=\n%#v", &test.msg, msg)
			if !reflect.DeepEqual(test.msg.Payload, msg.Payload) {
				t.Errorf("Payload:\n%#v !=\n%#v", test.msg.Payload, msg.Payload)
			}
			if !reflect.DeepEqual(test.msg.Headers, msg.Headers) {
				t.Errorf("Headers:\n%#v !=\n%#v", test.msg.Headers, msg.Headers)
			}
			if !reflect.DeepEqual(test.msg.Via, msg.Via) {
				t.Errorf("Via:\n%#v !=\n%#v", test.msg.Via, msg.Via)
			}
			if !reflect.DeepEqual(test.msg.Request, msg.Request) {
				t.Errorf("Request:\n%#v !=\n%#v", test.msg.Request, msg.Request)
			}
			if !reflect.DeepEqual(test.msg.To, msg.To) {
				t.Errorf("To:\n%#v !=\n%#v", test.msg.To, msg.To)
			}
			if !reflect.DeepEqual(test.msg.From, msg.From) {
				t.Errorf("From:\n%#v !=\n%#v", test.msg.From, msg.From)
			}
			if !reflect.DeepEqual(test.msg.Contact, msg.Contact) {
				t.Errorf("Contact:\n%#v !=\n%#v", test.msg.Contact, msg.Contact)
			}
		}
	}
}
