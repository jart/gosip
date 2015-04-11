// SIP Message Library

package sip

import (
	"bytes"
	"net"
	"strconv"
)

type Headers map[string]string

// Msg represents a SIP message. This can either be a request or a response.
// These fields are never nil unless otherwise specified.
type Msg struct {
	// Fields that aren't headers.
	VersionMajor uint8
	VersionMinor uint8
	Method       string  // Indicates type of request (if request)
	Request      *URI    // dest URI (nil if response)
	Status       int     // Indicates happiness of response (if response)
	Phrase       string  // Explains happiness of response (if response)
	Payload      Payload // Stuff that comes after two line breaks

	// Special non-SIP fields.
	SourceAddr *net.UDPAddr // Set by transport layer as received address.

	// Important headers should be further up in the struct.
	From        *Addr  // Logical sender of message
	To          *Addr  // Logical destination of message
	Via         *Via   // Linked list of agents traversed (must have one)
	Route       *Addr  // Used for goose routing and loose routing
	RecordRoute *Addr  // Used for loose routing
	Contact     *Addr  // Where we send response packets or nil
	CallID      []byte // Identifies call from INVITE to BYE; never empty, only absent
	CSeq        int    // Counter for network packet ordering
	CSeqMethod  string // Helps with matching to orig message
	MaxForwards int    // 0 has context specific meaning
	UserAgent   []byte // Name of the SIP stack

	// All the other RFC 3261 headers in plus some extras.
	//
	// These byte slices will be nil if absent, in which case the header never
	// appeared. If they are empty, then the header did appear, but without a
	// value.
	Accept             []byte
	AcceptContact      []byte
	AcceptEncoding     []byte
	AcceptLanguage     []byte
	AlertInfo          []byte
	Allow              []byte
	AllowEvents        []byte
	AuthenticationInfo []byte
	Authorization      []byte
	CallInfo           []byte
	ContentDisposition []byte
	ContentEncoding    []byte
	ContentLanguage    []byte
	Date               []byte
	ErrorInfo          []byte
	Event              []byte
	Expires            int // Seconds registration should expire.
	InReplyTo          []byte
	MIMEVersion        []byte
	MinExpires         int // Registrars need this when responding
	Organization       []byte
	PAssertedIdentity  *Addr // P-Asserted-Identity or nil (used for PSTN ANI)
	Priority           []byte
	ProxyAuthenticate  []byte
	ProxyAuthorization []byte
	ProxyRequire       []byte
	ReferTo            []byte
	ReferredBy         []byte
	RemotePartyID      *Addr // Evil twin of P-Asserted-Identity.
	ReplyTo            []byte
	Require            []byte
	RetryAfter         []byte
	Server             []byte
	Subject            []byte
	Supported          []byte
	Timestamp          []byte
	Unsupported        []byte
	WWWAuthenticate    []byte
	Warning            []byte

	// Extension headers.
	Headers Headers
}

//go:generate ragel -Z -G2 -o msg_parse.go msg_parse.rl

func (msg *Msg) IsResponse() bool {
	return msg.Status > 0
}

func (msg *Msg) String() string {
	var b bytes.Buffer
	msg.Append(&b)
	return b.String()
}

func (msg *Msg) Copy() *Msg {
	if msg == nil {
		return nil
	}
	res := new(Msg)
	*res = *msg
	res.Request = msg.Request.Copy()
	res.To = msg.To.Copy()
	res.From = msg.From.Copy()
	res.Via = msg.Via.Copy()
	res.PAssertedIdentity = msg.PAssertedIdentity.Copy()
	res.RemotePartyID = msg.RemotePartyID.Copy()
	res.Route = msg.Route.Copy()
	res.Contact = msg.Contact.Copy()
	res.RecordRoute = msg.RecordRoute.Copy()
	res.Headers = make(Headers, len(msg.Headers))
	for k, v := range msg.Headers {
		res.Headers[k] = v
	}
	return res
}

// I turn a SIP message back into a packet.
func (msg *Msg) Append(b *bytes.Buffer) {
	if msg == nil {
		return
	}
	if !msg.IsResponse() {
		if msg.Method == "" {
			b.WriteString("INVITE ")
		} else {
			b.WriteString(msg.Method)
			b.WriteString(" ")
		}
		if msg.Request == nil {
			// In case of bugs, keep calm and DDOS NASA.
			b.WriteString("sip:www.nasa.gov:80")
		} else {
			msg.Request.Append(b)
		}
		b.WriteString(" ")
		msg.appendVersion(b)
		b.WriteString("\r\n")
	} else {
		msg.appendVersion(b)
		b.WriteString(" ")
		b.WriteString(strconv.Itoa(msg.Status))
		b.WriteString(" ")
		if msg.Phrase == "" {
			b.WriteString(Phrase(msg.Status))
		} else {
			b.WriteString(msg.Phrase)
		}
		b.WriteString("\r\n")
	}

	b.WriteString("From: ")
	msg.From.Append(b)
	b.WriteString("\r\n")

	b.WriteString("To: ")
	msg.To.Append(b)
	b.WriteString("\r\n")

	for viap := msg.Via; viap != nil; viap = viap.Next {
		b.WriteString("Via: ")
		viap.Append(b)
		b.WriteString("\r\n")
	}

	if msg.Route != nil {
		b.WriteString("Route: ")
		msg.Route.Append(b)
		b.WriteString("\r\n")
	}

	if msg.RecordRoute != nil {
		b.WriteString("Record-Route: ")
		msg.RecordRoute.Append(b)
		b.WriteString("\r\n")
	}

	if msg.Contact != nil {
		b.WriteString("Contact: ")
		msg.Contact.Append(b)
		b.WriteString("\r\n")
	}

	b.WriteString("Call-ID: ")
	b.Write(msg.CallID)
	b.WriteString("\r\n")

	b.WriteString("CSeq: ")
	b.WriteString(strconv.Itoa(msg.CSeq))
	b.WriteString(" ")
	b.WriteString(msg.CSeqMethod)
	b.WriteString("\r\n")

	if msg.UserAgent != nil {
		b.WriteString("User-Agent: ")
		b.Write(msg.UserAgent)
		b.WriteString("\r\n")
	}

	if !msg.IsResponse() {
		if msg.MaxForwards == 0 {
			b.WriteString("Max-Forwards: 70\r\n")
		} else {
			b.WriteString("Max-Forwards: ")
			b.WriteString(strconv.Itoa(msg.MaxForwards))
			b.WriteString("\r\n")
		}
	}

	if msg.Accept != nil {
		b.WriteString("Accept: ")
		b.Write(msg.Accept)
		b.WriteString("\r\n")
	}

	if msg.AcceptEncoding != nil {
		b.WriteString("Accept-Encoding: ")
		b.Write(msg.AcceptEncoding)
		b.WriteString("\r\n")
	}

	if msg.AcceptLanguage != nil {
		b.WriteString("Accept-Language: ")
		b.Write(msg.AcceptLanguage)
		b.WriteString("\r\n")
	}

	if msg.AlertInfo != nil {
		b.WriteString("Alert-Info: ")
		b.Write(msg.AlertInfo)
		b.WriteString("\r\n")
	}

	if msg.Allow != nil {
		b.WriteString("Allow: ")
		b.Write(msg.Allow)
		b.WriteString("\r\n")
	}

	if msg.AllowEvents != nil {
		b.WriteString("Allow-Events: ")
		b.Write(msg.AllowEvents)
		b.WriteString("\r\n")
	}

	if msg.AuthenticationInfo != nil {
		b.WriteString("Authentication-Info: ")
		b.Write(msg.AuthenticationInfo)
		b.WriteString("\r\n")
	}

	if msg.Authorization != nil {
		b.WriteString("Authorization: ")
		b.Write(msg.Authorization)
		b.WriteString("\r\n")
	}

	if msg.CallInfo != nil {
		b.WriteString("Call-Info: ")
		b.Write(msg.CallInfo)
		b.WriteString("\r\n")
	}

	if msg.ContentDisposition != nil {
		b.WriteString("Content-Disposition: ")
		b.Write(msg.ContentDisposition)
		b.WriteString("\r\n")
	}

	if msg.ContentEncoding != nil {
		b.WriteString("Content-Encoding: ")
		b.Write(msg.ContentEncoding)
		b.WriteString("\r\n")
	}

	if msg.ContentLanguage != nil {
		b.WriteString("Content-Language: ")
		b.Write(msg.ContentLanguage)
		b.WriteString("\r\n")
	}

	if msg.Date != nil {
		b.WriteString("Date: ")
		b.Write(msg.Date)
		b.WriteString("\r\n")
	}

	if msg.ErrorInfo != nil {
		b.WriteString("Error-Info: ")
		b.Write(msg.ErrorInfo)
		b.WriteString("\r\n")
	}

	if msg.Event != nil {
		b.WriteString("Event: ")
		b.Write(msg.Event)
		b.WriteString("\r\n")
	}

	// Expires is allowed to be 0 for for REGISTER stuff.
	if msg.Expires > 0 || msg.Method == "REGISTER" || msg.CSeqMethod == "REGISTER" {
		b.WriteString("Expires: ")
		b.WriteString(strconv.Itoa(msg.Expires))
		b.WriteString("\r\n")
	}

	if msg.InReplyTo != nil {
		b.WriteString("In-Reply-To: ")
		b.Write(msg.InReplyTo)
		b.WriteString("\r\n")
	}

	if msg.MIMEVersion != nil {
		b.WriteString("MIME-Version: ")
		b.Write(msg.MIMEVersion)
		b.WriteString("\r\n")
	}

	if msg.MinExpires > 0 {
		b.WriteString("Min-Expires: ")
		b.WriteString(strconv.Itoa(msg.MinExpires))
		b.WriteString("\r\n")
	}

	if msg.Organization != nil {
		b.WriteString("Organization: ")
		b.Write(msg.Organization)
		b.WriteString("\r\n")
	}

	if msg.PAssertedIdentity != nil {
		b.WriteString("P-Asserted-Identity: ")
		msg.PAssertedIdentity.Append(b)
		b.WriteString("\r\n")
	}

	if msg.Priority != nil {
		b.WriteString("Priority: ")
		b.Write(msg.Priority)
		b.WriteString("\r\n")
	}

	if msg.ProxyAuthenticate != nil {
		b.WriteString("Proxy-Authenticate: ")
		b.Write(msg.ProxyAuthenticate)
		b.WriteString("\r\n")
	}

	if msg.ProxyAuthorization != nil {
		b.WriteString("Proxy-Authorization: ")
		b.Write(msg.ProxyAuthorization)
		b.WriteString("\r\n")
	}

	if msg.ProxyRequire != nil {
		b.WriteString("Proxy-Require: ")
		b.Write(msg.ProxyRequire)
		b.WriteString("\r\n")
	}

	if msg.ReferTo != nil {
		b.WriteString("Refer-To: ")
		b.Write(msg.ReferTo)
		b.WriteString("\r\n")
	}

	if msg.ReferredBy != nil {
		b.WriteString("Referred-By: ")
		b.Write(msg.ReferredBy)
		b.WriteString("\r\n")
	}

	if msg.RemotePartyID != nil {
		b.WriteString("Remote-Party-ID: ")
		msg.RemotePartyID.Append(b)
		b.WriteString("\r\n")
	}

	if msg.ReplyTo != nil {
		b.WriteString("Reply-To: ")
		b.Write(msg.ReplyTo)
		b.WriteString("\r\n")
	}

	if msg.Require != nil {
		b.WriteString("Require: ")
		b.Write(msg.Require)
		b.WriteString("\r\n")
	}

	if msg.RetryAfter != nil {
		b.WriteString("RetryAfter: ")
		b.Write(msg.RetryAfter)
		b.WriteString("\r\n")
	}

	if msg.Server != nil {
		b.WriteString("Server: ")
		b.Write(msg.Server)
		b.WriteString("\r\n")
	}

	if msg.Subject != nil {
		b.WriteString("Subject: ")
		b.Write(msg.Subject)
		b.WriteString("\r\n")
	}

	if msg.Supported != nil {
		b.WriteString("Supported: ")
		b.Write(msg.Supported)
		b.WriteString("\r\n")
	}

	if msg.Timestamp != nil {
		b.WriteString("Timestamp: ")
		b.Write(msg.Timestamp)
		b.WriteString("\r\n")
	}

	if msg.Unsupported != nil {
		b.WriteString("Unsupported: ")
		b.Write(msg.Unsupported)
		b.WriteString("\r\n")
	}

	if msg.Warning != nil {
		b.WriteString("Warning: ")
		b.Write(msg.Warning)
		b.WriteString("\r\n")
	}

	if msg.WWWAuthenticate != nil {
		b.WriteString("WWW-Authenticate: ")
		b.Write(msg.WWWAuthenticate)
		b.WriteString("\r\n")
	}

	if msg.Headers != nil {
		for k, v := range msg.Headers {
			b.WriteString(k)
			b.WriteString(": ")
			b.WriteString(v)
			b.WriteString("\r\n")
		}
	}

	if msg.Payload != nil {
		b.WriteString("Content-Type: ")
		b.WriteString(msg.Payload.ContentType())
		b.WriteString("\r\n")
		payload := msg.Payload.Data()
		b.WriteString("Content-Length: ")
		b.WriteString(strconv.Itoa(len(payload)))
		b.WriteString("\r\n\r\n")
		b.Write(payload)
	} else {
		b.WriteString("Content-Length: 0\r\n\r\n")
	}
}

func (msg *Msg) appendVersion(b *bytes.Buffer) {
	b.WriteString("SIP/")
	if msg.VersionMajor == 0 {
		b.WriteString("2.0")
	} else {
		b.WriteString(strconv.FormatUint(uint64(msg.VersionMajor), 10))
		b.WriteString(".")
		b.WriteString(strconv.FormatUint(uint64(msg.VersionMinor), 10))
	}
}
