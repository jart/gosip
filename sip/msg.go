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

// SIP Message Library

package sip

import (
	"bytes"
	"net"
	"strconv"
)

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
	CallID      string // Identifies call from invite to bye
	CSeq        int    // Counter for network packet ordering
	CSeqMethod  string // Helps with matching to orig message
	MaxForwards int    // 0 has context specific meaning
	UserAgent   string

	// All the other RFC 3261 headers in plus some extras.
	Accept             string
	AcceptContact      string
	AcceptEncoding     string
	AcceptLanguage     string
	AlertInfo          string
	Allow              string
	AllowEvents        string
	AuthenticationInfo string
	Authorization      string
	CallInfo           string
	ContentDisposition string
	ContentEncoding    string
	ContentLanguage    string
	Date               string
	ErrorInfo          string
	Event              string
	Expires            int // Seconds registration should expire.
	InReplyTo          string
	MIMEVersion        string
	MinExpires         int // Registrars need this when responding
	Organization       string
	PAssertedIdentity  *Addr // P-Asserted-Identity or nil (used for PSTN ANI)
	Priority           string
	ProxyAuthenticate  string
	ProxyAuthorization string
	ProxyRequire       string
	ReferTo            string
	ReferredBy         string
	RemotePartyID      *Addr // Evil twin of P-Asserted-Identity.
	ReplyTo            string
	Require            string
	RetryAfter         string
	Server             string
	Subject            string
	Supported          string
	Timestamp          string
	Unsupported        string
	WWWAuthenticate    string
	Warning            string

	// Extension headers.
	XHeader *XHeader
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
	res.XHeader = msg.XHeader
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
		if msg.Phrase == "" {
			msg.Phrase = Phrase(msg.Status)
		}
		msg.appendVersion(b)
		b.WriteString(" ")
		b.WriteString(strconv.Itoa(msg.Status))
		b.WriteString(" ")
		b.WriteString(msg.Phrase)
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
	b.WriteString(msg.CallID)
	b.WriteString("\r\n")

	b.WriteString("CSeq: ")
	b.WriteString(strconv.Itoa(msg.CSeq))
	b.WriteString(" ")
	b.WriteString(msg.CSeqMethod)
	b.WriteString("\r\n")

	if msg.UserAgent != "" {
		b.WriteString("User-Agent: ")
		b.WriteString(msg.UserAgent)
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

	if msg.Accept != "" {
		b.WriteString("Accept: ")
		b.WriteString(msg.Accept)
		b.WriteString("\r\n")
	}

	if msg.AcceptEncoding != "" {
		b.WriteString("Accept-Encoding: ")
		b.WriteString(msg.AcceptEncoding)
		b.WriteString("\r\n")
	}

	if msg.AcceptLanguage != "" {
		b.WriteString("Accept-Language: ")
		b.WriteString(msg.AcceptLanguage)
		b.WriteString("\r\n")
	}

	if msg.AlertInfo != "" {
		b.WriteString("Alert-Info: ")
		b.WriteString(msg.AlertInfo)
		b.WriteString("\r\n")
	}

	if msg.Allow != "" {
		b.WriteString("Allow: ")
		b.WriteString(msg.Allow)
		b.WriteString("\r\n")
	}

	if msg.AllowEvents != "" {
		b.WriteString("Allow-Events: ")
		b.WriteString(msg.AllowEvents)
		b.WriteString("\r\n")
	}

	if msg.AuthenticationInfo != "" {
		b.WriteString("Authentication-Info: ")
		b.WriteString(msg.AuthenticationInfo)
		b.WriteString("\r\n")
	}

	if msg.Authorization != "" {
		b.WriteString("Authorization: ")
		b.WriteString(msg.Authorization)
		b.WriteString("\r\n")
	}

	if msg.CallInfo != "" {
		b.WriteString("Call-Info: ")
		b.WriteString(msg.CallInfo)
		b.WriteString("\r\n")
	}

	if msg.ContentDisposition != "" {
		b.WriteString("Content-Disposition: ")
		b.WriteString(msg.ContentDisposition)
		b.WriteString("\r\n")
	}

	if msg.ContentEncoding != "" {
		b.WriteString("Content-Encoding: ")
		b.WriteString(msg.ContentEncoding)
		b.WriteString("\r\n")
	}

	if msg.ContentLanguage != "" {
		b.WriteString("Content-Language: ")
		b.WriteString(msg.ContentLanguage)
		b.WriteString("\r\n")
	}

	if msg.Date != "" {
		b.WriteString("Date: ")
		b.WriteString(msg.Date)
		b.WriteString("\r\n")
	}

	if msg.ErrorInfo != "" {
		b.WriteString("Error-Info: ")
		b.WriteString(msg.ErrorInfo)
		b.WriteString("\r\n")
	}

	if msg.Event != "" {
		b.WriteString("Event: ")
		b.WriteString(msg.Event)
		b.WriteString("\r\n")
	}

	// Expires is allowed to be 0 for for REGISTER stuff.
	if msg.Expires > 0 || msg.Method == "REGISTER" || msg.CSeqMethod == "REGISTER" {
		b.WriteString("Expires: ")
		b.WriteString(strconv.Itoa(msg.Expires))
		b.WriteString("\r\n")
	}

	if msg.InReplyTo != "" {
		b.WriteString("In-Reply-To: ")
		b.WriteString(msg.InReplyTo)
		b.WriteString("\r\n")
	}

	if msg.MIMEVersion != "" {
		b.WriteString("MIME-Version: ")
		b.WriteString(msg.MIMEVersion)
		b.WriteString("\r\n")
	}

	if msg.MinExpires > 0 {
		b.WriteString("Min-Expires: ")
		b.WriteString(strconv.Itoa(msg.MinExpires))
		b.WriteString("\r\n")
	}

	if msg.Organization != "" {
		b.WriteString("Organization: ")
		b.WriteString(msg.Organization)
		b.WriteString("\r\n")
	}

	if msg.PAssertedIdentity != nil {
		b.WriteString("P-Asserted-Identity: ")
		msg.PAssertedIdentity.Append(b)
		b.WriteString("\r\n")
	}

	if msg.Priority != "" {
		b.WriteString("Priority: ")
		b.WriteString(msg.Priority)
		b.WriteString("\r\n")
	}

	if msg.ProxyAuthenticate != "" {
		b.WriteString("Proxy-Authenticate: ")
		b.WriteString(msg.ProxyAuthenticate)
		b.WriteString("\r\n")
	}

	if msg.ProxyAuthorization != "" {
		b.WriteString("Proxy-Authorization: ")
		b.WriteString(msg.ProxyAuthorization)
		b.WriteString("\r\n")
	}

	if msg.ProxyRequire != "" {
		b.WriteString("Proxy-Require: ")
		b.WriteString(msg.ProxyRequire)
		b.WriteString("\r\n")
	}

	if msg.ReferTo != "" {
		b.WriteString("Refer-To: ")
		b.WriteString(msg.ReferTo)
		b.WriteString("\r\n")
	}

	if msg.ReferredBy != "" {
		b.WriteString("Referred-By: ")
		b.WriteString(msg.ReferredBy)
		b.WriteString("\r\n")
	}

	if msg.RemotePartyID != nil {
		b.WriteString("Remote-Party-ID: ")
		msg.RemotePartyID.Append(b)
		b.WriteString("\r\n")
	}

	if msg.ReplyTo != "" {
		b.WriteString("Reply-To: ")
		b.WriteString(msg.ReplyTo)
		b.WriteString("\r\n")
	}

	if msg.Require != "" {
		b.WriteString("Require: ")
		b.WriteString(msg.Require)
		b.WriteString("\r\n")
	}

	if msg.RetryAfter != "" {
		b.WriteString("RetryAfter: ")
		b.WriteString(msg.RetryAfter)
		b.WriteString("\r\n")
	}

	if msg.Server != "" {
		b.WriteString("Server: ")
		b.WriteString(msg.Server)
		b.WriteString("\r\n")
	}

	if msg.Subject != "" {
		b.WriteString("Subject: ")
		b.WriteString(msg.Subject)
		b.WriteString("\r\n")
	}

	if msg.Supported != "" {
		b.WriteString("Supported: ")
		b.WriteString(msg.Supported)
		b.WriteString("\r\n")
	}

	if msg.Timestamp != "" {
		b.WriteString("Timestamp: ")
		b.WriteString(msg.Timestamp)
		b.WriteString("\r\n")
	}

	if msg.Unsupported != "" {
		b.WriteString("Unsupported: ")
		b.WriteString(msg.Unsupported)
		b.WriteString("\r\n")
	}

	if msg.Warning != "" {
		b.WriteString("Warning: ")
		b.WriteString(msg.Warning)
		b.WriteString("\r\n")
	}

	if msg.WWWAuthenticate != "" {
		b.WriteString("WWW-Authenticate: ")
		b.WriteString(msg.WWWAuthenticate)
		b.WriteString("\r\n")
	}

	msg.XHeader.Append(b)

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
