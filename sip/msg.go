// SIP Message Library

package sip

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type Headers map[string]string

// Msg represents a SIP message. This can either be a request or a response.
// These fields are never nil unless otherwise specified.
type Msg struct {
	// Special non-SIP fields.
	SourceAddr *net.UDPAddr // Set by transport layer as received address

	// Fields that aren't headers.
	IsResponse bool   // This is a response (like 404 Not Found)
	Method     string // Indicates type of request (if request)
	Request    *URI   // dest URI (nil if response)
	Status     int    // Indicates happiness of response (if response)
	Phrase     string // Explains happiness of response (if response)
	Payload    string // Stuff that comes after two line breaks

	// Mandatory headers.
	Via         *Via   // Linked list of agents traversed (must have one)
	Route       *Addr  // Used for goose routing and loose routing
	RecordRoute *Addr  // Used for loose routing
	From        *Addr  // Logical sender of message
	To          *Addr  // Logical destination of message
	CallID      string // Identifies call from invite to bye
	CSeq        int    // Counter for network packet ordering
	CSeqMethod  string // Helps with matching to orig message

	// Convenience headers.
	MaxForwards int   // 0 has context specific meaning
	MinExpires  int   // Registrars need this when responding
	Expires     int   // Seconds registration should expire
	Paid        *Addr // P-Asserted-Identity or nil (used for PSTN ANI)
	Rpid        *Addr // Remote-Party-Id or nil
	Contact     *Addr // Where we send response packets or nil

	// All the other headers (never nil)
	Headers Headers
}

func (msg *Msg) String() string {
	if msg == nil {
		return ""
	}
	var b bytes.Buffer
	if err := msg.Append(&b); err != nil {
		log.Println("Bad SIP message!", err)
		return ""
	}
	return b.String()
}

// Parses a SIP message into a data structure. This takes ~70 µs on average.
func ParseMsg(packet string) (msg *Msg, err error) {
	msg = new(Msg)
	if packet == "" {
		return nil, errors.New("Empty msg")
	}
	if n := strings.Index(packet, "\r\n\r\n"); n > 0 {
		packet, msg.Payload = packet[0:n], packet[n+4:]
	}
	lines := strings.Split(packet, "\r\n")
	if lines == nil || len(lines) < 2 {
		return nil, errors.New("Too few lines")
	}
	var k, v string
	var okVia, okTo, okFrom, okCallID, okComputer bool
	err = msg.parseFirstLine(lines[0])
	if err != nil {
		return nil, err
	}
	hdrs := lines[1:]
	msg.Headers = make(map[string]string, len(hdrs))
	msg.MaxForwards = 70
	viap := &msg.Via
	contactp := &msg.Contact
	routep := &msg.Route
	rroutep := &msg.RecordRoute
	for _, hdr := range hdrs {
		if hdr == "" {
			continue
		}
		if hdr[0] == ' ' || hdr[0] == '\t' {
			v = strings.Trim(hdr, "\t ") // Line continuation.
		} else {
			if i := strings.Index(hdr, ": "); i > 0 {
				k, v = hdr[0:i], hdr[i+2:]
				if k == "" || v == "" {
					log.Println("[NOTICE]", "blank header found", hdr)
				}
				k = uncompactHeader(k)
			} else {
				log.Println("[NOTICE]", "header missing delimiter", hdr)
				continue
			}
		}
		switch strings.ToLower(k) {
		case "call-id":
			okCallID = true
			msg.CallID = v
		case "via":
			okVia = true
			*viap, err = ParseVia(v)
			if err != nil {
				return nil, errors.New("Bad Via header: " + err.Error())
			} else {
				viap = &(*viap).Next
			}
		case "to":
			okTo = true
			msg.To, err = ParseAddr(v)
			if err != nil {
				return nil, errors.New("Bad To header: " + err.Error())
			}
		case "from":
			okFrom = true
			msg.From, err = ParseAddr(v)
			if err != nil {
				return nil, errors.New("Bad From header: " + err.Error())
			}
		case "contact":
			*contactp, err = ParseAddr(v)
			if err != nil {
				return nil, errors.New("Bad Contact header: " + err.Error())
			} else {
				contactp = &(*contactp).Last().Next
			}
		case "cseq":
			okComputer = false
			if n := strings.Index(v, " "); n > 0 {
				sseq, method := v[0:n], v[n+1:]
				if seq, err := strconv.Atoi(sseq); err == nil {
					msg.CSeq, msg.CSeqMethod = seq, method
					okComputer = true
				}
			}
			if !okComputer {
				return nil, errors.New("Bad CSeq Header")
			}
		case "content-length":
			if cl, err := strconv.Atoi(v); err == nil {
				if cl != len(msg.Payload) {
					return nil, errors.New(fmt.Sprintf(
						"Content-Length (%d) differs from payload length (%d)",
						cl, len(msg.Payload)))
				}
			} else {
				return nil, errors.New("Bad Content-Length header")
			}
		case "expires":
			if cl, err := strconv.Atoi(v); err == nil && cl >= 0 {
				msg.Expires = cl
			} else {
				return nil, errors.New("Bad Expires header")
			}
		case "min-expires":
			if cl, err := strconv.Atoi(v); err == nil && cl > 0 {
				msg.MinExpires = cl
			} else {
				return nil, errors.New("Bad Min-Expires header")
			}
		case "max-forwards":
			if cl, err := strconv.Atoi(v); err == nil && cl > 0 {
				msg.MaxForwards = cl
			} else {
				return nil, errors.New("Bad Max-Forwards header")
			}
		case "route":
			*routep, err = ParseAddr(v)
			if err != nil {
				return nil, errors.New("Bad Route header: " + err.Error())
			} else {
				routep = &(*routep).Last().Next
			}
		case "record-route":
			*rroutep, err = ParseAddr(v)
			if err != nil {
				return nil, errors.New("Bad Record-Route header: " + err.Error())
			} else {
				rroutep = &(*rroutep).Last().Next
			}
		case "p-asserted-identity":
			msg.Paid, err = ParseAddr(v)
			if err != nil {
				return nil, errors.New("Bad P-Asserted-Identity header: " + err.Error())
			}
		case "remote-party-id":
			msg.Rpid, err = ParseAddr(v)
			if err != nil {
				return nil, errors.New("Bad Remote-Party-ID header: " + err.Error())
			}
		default:
			msg.Headers[k] = v
		}
	}
	if !okVia || !okTo || !okFrom || !okCallID || !okComputer {
		return nil, errors.New("Missing mandatory headers")
	}
	return
}

func (msg *Msg) Copy() *Msg {
	if msg == nil {
		return nil
	}
	res := new(Msg)
	*res = *msg
	res.To = msg.To.Copy()
	res.From = msg.From.Copy()
	res.Via = msg.Via.Copy()
	res.Paid = msg.Paid.Copy()
	res.Rpid = msg.Rpid.Copy()
	res.Route = msg.Route.Copy()
	res.Request = msg.Request.Copy()
	res.Contact = msg.Contact.Copy()
	res.RecordRoute = msg.RecordRoute.Copy()
	res.Headers = make(map[string]string, len(msg.Headers))
	for k, v := range msg.Headers {
		res.Headers[k] = v
	}
	return res
}

// i turn a sip message back into a packet
func (msg *Msg) Append(b *bytes.Buffer) error {
	if !msg.IsResponse {
		if msg.Method == "" {
			return errors.New("Msg.Method not set")
		}
		if msg.Request == nil {
			return errors.New("msg.Request not set")
		}
		b.WriteString(msg.Method)
		b.WriteString(" ")
		msg.Request.Append(b)
		b.WriteString(" SIP/2.0\r\n")
	} else {
		if msg.Status < 100 {
			return errors.New("Msg.Status < 100")
		}
		if msg.Status >= 700 {
			return errors.New("Msg.Status >= 700")
		}
		if msg.Phrase == "" {
			msg.Phrase = Phrase(msg.Status)
		}
		b.WriteString("SIP/2.0 ")
		b.WriteString(strconv.Itoa(msg.Status))
		b.WriteString(" ")
		b.WriteString(msg.Phrase)
		b.WriteString("\r\n")
	}

	if msg.Via == nil {
		return errors.New("Need moar Via headers")
	}
	for viap := msg.Via; viap != nil; viap = viap.Next {
		b.WriteString("Via: ")
		if err := viap.Append(b); err != nil {
			return err
		}
		b.WriteString("\r\n")
	}

	if msg.MaxForwards < 0 {
		return errors.New("MaxForwards is less than 0!!")
	} else if msg.MaxForwards == 0 {
		b.WriteString("Max-Forwards: 70\r\n")
	} else {
		b.WriteString("Max-Forwards: ")
		b.WriteString(strconv.Itoa(msg.MaxForwards))
		b.WriteString("\r\n")
	}

	b.WriteString("From: ")
	msg.From.Append(b)
	b.WriteString("\r\n")

	b.WriteString("To: ")
	msg.To.Append(b)
	b.WriteString("\r\n")

	if msg.CallID == "" {
		return errors.New("CallID is blank")
	}
	b.WriteString("Call-ID: ")
	b.WriteString(msg.CallID)
	b.WriteString("\r\n")

	if msg.CSeq < 0 || msg.CSeqMethod == "" {
		return errors.New("Bad CSeq")
	}
	b.WriteString("CSeq: ")
	b.WriteString(strconv.Itoa(msg.CSeq))
	b.WriteString(" ")
	b.WriteString(msg.CSeqMethod)
	b.WriteString("\r\n")

	if msg.Contact != nil {
		b.WriteString("Contact: ")
		msg.Contact.Append(b)
		b.WriteString("\r\n")
	}

	// Expires is allowed to be 0 for for REGISTER stuff.
	if msg.Expires > 0 ||
		msg.Method == "REGISTER" || msg.CSeqMethod == "REGISTER" {
		b.WriteString("Expires: ")
		b.WriteString(strconv.Itoa(msg.Expires))
		b.WriteString("\r\n")
	}

	if msg.MinExpires > 0 {
		b.WriteString("Min-Expires: ")
		b.WriteString(strconv.Itoa(msg.MinExpires))
		b.WriteString("\r\n")
	}

	if msg.Headers != nil {
		for k, v := range msg.Headers {
			if k == "" || v == "" {
				return errors.New("Header blank")
			}
			b.WriteString(k)
			b.WriteString(": ")
			b.WriteString(v)
			b.WriteString("\r\n")
		}
	}

	if msg.Paid != nil {
		b.WriteString("P-Asserted-Identity: ")
		msg.Paid.Append(b)
		b.WriteString("\r\n")
	}

	if msg.Rpid != nil {
		b.WriteString("Remote-Party-ID: ")
		msg.Rpid.Append(b)
		b.WriteString("\r\n")
	}

	b.WriteString("Content-Length: ")
	b.WriteString(strconv.Itoa(len(msg.Payload)))
	b.WriteString("\r\n\r\n")
	b.WriteString(msg.Payload)

	return nil
}

func (msg *Msg) parseFirstLine(s string) (err error) {
	toks := strings.Split(s, " ")
	if toks != nil && len(toks) == 3 && toks[2] == "SIP/2.0" {
		msg.Phrase = ""
		msg.Status = 0
		msg.Method = toks[0]
		msg.Request = new(URI)
		msg.Request, err = ParseURI(toks[1])
	} else if toks != nil && len(toks) == 3 && toks[0] == "SIP/2.0" {
		msg.IsResponse = true
		msg.Method = ""
		msg.Request = nil
		msg.Phrase = toks[2]
		msg.Status, err = strconv.Atoi(toks[1])
		if err != nil {
			return errors.New("Invalid status")
		}
	} else {
		err = errors.New("Bad protocol or request line")
	}
	return err
}
