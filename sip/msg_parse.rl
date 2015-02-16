// -*-go-*-

package sip

import (
	"errors"
	"fmt"
	"github.com/jart/gosip/sdp"
)

%% machine msg;
%% write data;

// ParseMsg turns a a SIP message into a data structure.
func ParseMsg(s string) (msg *Msg, err error) {
	if s == "" {
		return nil, errors.New("Empty SIP message")
	}
	return ParseMsgBytes([]byte(s))
}

// ParseMsg turns a a SIP message byte slice into a data structure.
func ParseMsgBytes(data []byte) (msg *Msg, err error) {
	if data == nil {
		return nil, nil
	}
	msg = new(Msg)
	viap := &msg.Via
	routep := &msg.Route
	rroutep := &msg.RecordRoute
	contactp := &msg.Contact
	cs := 0
	p := 0
	pe := len(data)
	//eof := len(data)
	//stack := make([]int, 2)
	//top := 0
	line := 1
	linep := 0
	buf := make([]byte, len(data))
	amt := 0
	mark := 0
	clen := 0
	ctype := ""
	var b1 string
	var hex byte
	var dest *string

	%%{
		action break {
			fbreak;
		}

		action mark {
			mark = p
		}

		action start {
			amt = 0
		}

		action append {
			buf[amt] = fc
			amt++
		}

		action collapse {
			amt = appendCollapse(buf, amt, fc)
		}

		action hexHi {
			hex = unhex(fc) * 16
		}

		action hexLo {
			hex += unhex(fc)
			buf[amt] = hex
			amt++
		}

		action lower {
			amt = appendLower(buf, amt, fc)
		}

		action Method {
			msg.Method = string(data[mark:p])
		}

		action VersionMajor {
			msg.VersionMajor = msg.VersionMajor * 10 + (fc - 0x30)
		}

		action VersionMinor {
			msg.VersionMinor = msg.VersionMinor * 10 + (fc - 0x30)
		}

		action RequestURI {
			msg.Request, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		}

		action StatusCode {
			msg.Status = msg.Status * 10 + (int(fc) - 0x30)
		}

		action ReasonPhrase {
			msg.Phrase = string(buf[0:amt])
		}

		action header {
			fgoto header;
		}

		action hname {
			b1 = string(bytes.ToLower(data[mark:p]))
		}

		action svalue {
			fhold;
			fgoto svalue;
		}

		action svalueDone {
			if dest != nil {
				*dest = string(data[mark:p - 1])
			} else {
				if msg.Headers == nil {
					msg.Headers = Headers{}
				}
				msg.Headers[b1] = string(data[mark:p])
			}
		}

		action CallID {
			msg.CallID = string(data[mark:p])
		}

		action Contact {
			*contactp, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *contactp != nil { contactp = &(*contactp).Next }
		}

		action ContentLength {
			clen = clen * 10 + (int(fc) - 0x30)
		}

		action ContentType {
			ctype = string(data[mark:p])
		}

		action CSeq {
			msg.CSeq = msg.CSeq * 10 + (int(fc) - 0x30)
		}

		action CSeqMethod {
			msg.CSeqMethod = string(data[mark:p])
		}

		action Expires {
			msg.Expires = msg.Expires * 10 + (int(fc) - 0x30)
		}

		action From {
			msg.From, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		}

		action MaxForwards {
			msg.MaxForwards = msg.MaxForwards * 10 + (int(fc) - 0x30)
		}

		action MinExpires {
			msg.MinExpires = msg.MinExpires * 10 + (int(fc) - 0x30)
		}

		action PAssertedIdentity {
			msg.PAssertedIdentity, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		}

		action RecordRoute {
			*rroutep, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *rroutep != nil { rroutep = &(*rroutep).Next }
		}

		action RemotePartyID {
			msg.RemotePartyID, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		}

		action Route {
			*routep, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *routep != nil { routep = &(*routep).Next }
		}

		action To {
			msg.To, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		}

		action Via {
			*viap, err = ParseVia(string(data[mark:p]))
			if err != nil { return nil, err }
			for *viap != nil { viap = &(*viap).Next }
		}

		action lookAheadWSP { p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t') }

		# https://tools.ietf.org/html/rfc2234
		SP              = " ";
		HTAB            = "\t";
		CR              = "\r";
		LF              = "\n" @{ line++; linep = p; };
		CRLF            = CR LF;
		WSP             = SP | HTAB;
		LWS             = ( WSP* ( CR when lookAheadWSP ) LF )? WSP+;
		SWS             = LWS?;
		UTF8_NONASCII   = 0x80..0xFD;
		UTF8            = 0x21..0x7F | UTF8_NONASCII;
		UTF8_TRIM       = ( UTF8+ (LWS* UTF8)* ) >start @collapse;

		# https://tools.ietf.org/html/rfc3261#section-25.1
		reserved        = ";" | "/" | "?" | ":" | "@" | "&" | "=" | "+" | "$" | "," ;
		mark            = "-" | "_" | "." | "!" | "~" | "*" | "'" | "(" | ")" ;
		unreserved      = alpha | digit | mark ;
		tokenc          = alpha | digit | "-" | "." | "!" | "%" | "*" | "_"
		                | "+" | "`" | "'" | "~" ;
		separators      = "("  | ")" | "<" | ">" | "@" | "," | ";" | ":" | "\\"
		                | "\"" | "/" | "[" | "]" | "?" | "=" | "{" | "}" | SP
		                | HTAB ;
		wordc           = alpha | digit | "-" | "." | "!" | "%" | "*" | "_"
		                | "+" | "`" | "'" | "~" | "(" | ")" | "<" | ">" | ":"
		                | "\\" | "\"" | "/" | "[" | "]" | "?" | "{" | "}" ;
		word            = wordc+ ;
		STAR            = SWS "*" SWS ;
		SLASH           = SWS "/" SWS ;
		EQUAL           = SWS "=" SWS ;
		LPAREN          = SWS "(" SWS ;
		RPAREN          = SWS ")" SWS ;
		RAQUOT          = ">" SWS ;
		LAQUOT          = SWS "<" ;
		COMMA           = SWS "," SWS ;
		SEMI            = SWS ";" SWS ;
		COLON           = SWS ":" SWS ;
		HCOLON          = WSP* ":" SWS ;
		LDQUOT          = SWS "\"" ;
		RDQUOT          = "\"" SWS ;
		ctext           = 0x21..0x27 | 0x2A..0x5B | 0x5D..0x7E | UTF8_NONASCII | LWS ;
		quoted_pair     = "\\" ( 0x00..0x09 | 0x0B..0x0C | 0x0E..0x7F ) ;
		comment         = LPAREN ( ctext | quoted_pair )* <: RPAREN ;  # TODO(jart): Nested parens
		qdtext          = LWS | 0x21 | 0x23..0x5B | 0x5D..0x7E | UTF8_NONASCII ;
		quoted_string   = SWS "\"" ( qdtext | quoted_pair )* <: "\"" ;
		escaped         = "%" ( xdigit @hexHi ) ( xdigit @hexLo ) ;
		uric            = reserved | unreserved | escaped ;
		uric_no_slash   = unreserved | escaped | ";" | "?" | ":" | "@" | "&" | "="
		                | "+" | "$" | "," ;
		token           = tokenc+ >mark ;
		reasonc         = reserved | unreserved | UTF8_NONASCII | SP | HTAB ;
		reasonmc        = escaped | ( reasonc @append ) ;
		cid             = word ( "@" word )? ;

		Method          = token %Method;
		SIPVersionNo    = digit+ @VersionMajor "." digit+ @VersionMinor;
		RequestURI      = ^SP+ >mark %RequestURI;
		StatusCode      = ( digit @StatusCode ) {3};
		ReasonPhrase    = reasonmc+ >start %ReasonPhrase;
		hval            = ( UTF8 | LWS )* >mark;

		cheader  = ("Call-ID"i | "i"i) HCOLON cid >mark %CallID
		         | ("Contact"i | "m"i) HCOLON <: hval %Contact
		         | ("Content-Length"i | "l"i) HCOLON digit+ >{clen=0} @ContentLength
		         | ("Content-Type"i | "c"i) HCOLON <: hval %ContentType
		         | "CSeq"i HCOLON (digit+ @CSeq) LWS token %CSeqMethod
		         | ("Expires"i | "l"i) HCOLON digit+ >{msg.Expires=0} @Expires
		         | ("From"i | "f"i) HCOLON <: hval %From
		         | ("Max-Forwards"i | "l"i) HCOLON digit+ >{msg.MaxForwards=0} @MaxForwards
		         | ("Min-Expires"i | "l"i) HCOLON digit+ >{msg.MinExpires=0} @MinExpires
		         | "P-Asserted-Identity"i HCOLON <: hval %PAssertedIdentity
		         | "Record-Route"i HCOLON <: hval %RecordRoute
		         | "Remote-Party-ID"i HCOLON <: hval %RemotePartyID
		         | "Route"i HCOLON <: hval %Route
		         | ("To"i | "t"i) HCOLON <: hval %To
		         | ("Via"i | "v"i) HCOLON <: hval %Via
		         ;

		sname    = "Accept"i %{dest=&msg.Accept}
		         | ("Accept-Contact"i | "a"i) %{dest=&msg.AcceptContact}
		         | "Accept-Encoding"i %{dest=&msg.AcceptEncoding}
		         | "Accept-Language"i %{dest=&msg.AcceptLanguage}
		         | ("Allow"i | "u"i) %{dest=&msg.Allow}
		         | ("Allow-Events"i | "u"i) %{dest=&msg.AllowEvents}
		         | "Alert-Info"i %{dest=&msg.AlertInfo}
		         | "Authentication-Info"i %{dest=&msg.AuthenticationInfo}
		         | "Authorization"i %{dest=&msg.Authorization}
		         | "Content-Disposition"i %{dest=&msg.ContentDisposition}
		         | "Content-Language"i %{dest=&msg.ContentLanguage}
		         | ("Content-Encoding"i | "e"i) %{dest=&msg.ContentEncoding}
		         | "Call-Info"i %{dest=&msg.CallInfo}
		         | "Date"i %{dest=&msg.Date}
		         | "Error-Info"i %{dest=&msg.ErrorInfo}
		         | ("Event"i | "o"i) %{dest=&msg.Event}
		         | "In-Reply-To"i %{dest=&msg.InReplyTo}
		         | "Reply-To"i %{dest=&msg.ReplyTo}
		         | "MIME-Version"i %{dest=&msg.MIMEVersion}
		         | "Organization"i %{dest=&msg.Organization}
		         | "Priority"i %{dest=&msg.Priority}
		         | "Proxy-Authenticate"i %{dest=&msg.ProxyAuthenticate}
		         | "Proxy-Authorization"i %{dest=&msg.ProxyAuthorization}
		         | "Proxy-Require"i %{dest=&msg.ProxyRequire}
		         | ("Refer-To"i | "r"i) %{dest=&msg.ReferTo}
		         | ("Referred-By"i | "b"i) %{dest=&msg.ReferredBy}
		         | "Require"i %{dest=&msg.Require}
		         | "Retry-After"i %{dest=&msg.RetryAfter}
		         | "Server"i %{dest=&msg.Server}
		         | ("Subject"i | "s"i) %{dest=&msg.Subject}
		         | ("Supported"i | "k"i) %{dest=&msg.Supported}
		         | "Timestamp"i %{dest=&msg.Timestamp}
		         | "Unsupported"i %{dest=&msg.Unsupported}
		         | "User-Agent"i %{dest=&msg.UserAgent}
		         | "Warning"i %{dest=&msg.Warning}
		         | "WWW-Authenticate"i %{dest=&msg.WWWAuthenticate}
		         ;

		svalue  := hval CR LF @svalueDone @header;
		xheader  = token %hname HCOLON @{dest=nil} @svalue;
		header  := CR LF @break
		         | sname HCOLON <: any @svalue
		         | cheader CR LF @header
		         ;

		SIPVersion    = "SIP/" SIPVersionNo;
		Request       = Method SP RequestURI SP SIPVersion CR LF @header;
		Response      = SIPVersion SP StatusCode SP ReasonPhrase CR LF @header;
		main         := Request | Response;

		write init;
		write exec;
	}%%

	if cs < msg_first_final {
		if p == pe {
			return nil, errors.New(fmt.Sprintf("Incomplete SIP message: %s", data))
		} else {
			return nil, errors.New(fmt.Sprintf("Error in SIP message at line %d offset %d:\n%s", line, p - linep, data))
		}
	}

	if clen > 0 {
		if clen != len(data) - p {
			return nil, errors.New(fmt.Sprintf("Content-Length incorrect: %d != %d", clen, len(data) - p))
		}
		if ctype == sdp.ContentType {
			ms, err := sdp.Parse(string(data[p:len(data)]))
			if err != nil { return nil, err }
			msg.Payload = ms
		} else {
			msg.Payload = &MiscPayload{T: ctype, D: data[p:len(data)]}
		}
	}

	return msg, nil
}
