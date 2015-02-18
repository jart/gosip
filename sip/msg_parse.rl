// -*-go-*-

package sip

import (
	"errors"
	"fmt"
	"github.com/jart/gosip/sdp"
)

%% machine msg;
%% write data;

// ParseMsg turns a SIP message into a data structure.
func ParseMsg(s string) (msg *Msg, err error) {
	if s == "" {
		return nil, errors.New("Empty SIP message")
	}
	return ParseMsgBytes([]byte(s))
}

// ParseMsg turns a SIP message byte slice into a data structure.
func ParseMsgBytes(data []byte) (msg *Msg, err error) {
	if data == nil {
		return nil, nil
	}
	msg = new(Msg)
	viap := &msg.Via
	cs := 0
	p := 0
	pe := len(data)
	eof := len(data)
	line := 1
	linep := 0
	buf := make([]byte, len(data))
	amt := 0
	mark := 0
	clen := 0
	ctype := ""
	var name string
	var hex byte
	var value *string
	var addr **Addr

	%%{
		action break {
			fbreak;
		}

		action line {
			line++
			linep = p + 1
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

		action space {
			buf[amt] = ' '
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

		action goto_header {
			fgoto header;
		}

		action goto_value {
			fhold;
			fgoto value;
		}

		action goto_xheader {
			fhold;
			fgoto xheader;
		}

		# Shorthand notation.
		action gxh {
			fhold;
			fgoto xheader;
		}

		action store_name {
			name = string(data[mark:p])
		}

		action store_value {{
			b := data[mark:p - 1]
			if value != nil {
				*value = string(b)
			} else if addr != nil {
				*addr, err = ParseAddrBytes(b, *addr)
				if err != nil { return nil, err }
			} else {
				if msg.Headers == nil {
					msg.Headers = Headers{}
				}
				msg.Headers[name] = string(b)
			}
		}}

		action new_addr {
			addr = new(Addr)
		}

		action addr_display {
			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		}

		action addr_uri {
			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		}

		action addr_param {
			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		}

		action CallID {
			msg.CallID = string(data[mark:p])
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

		action MaxForwards {
			msg.MaxForwards = msg.MaxForwards * 10 + (int(fc) - 0x30)
		}

		action MinExpires {
			msg.MinExpires = msg.MinExpires * 10 + (int(fc) - 0x30)
		}

		action Via {
			*viap, err = ParseVia(string(data[mark:p]))
			if err != nil { return nil, err }
			for *viap != nil { viap = &(*viap).Next }
		}

		action lookAheadWSP { lookAheadWSP(data, p, pe) }

		# https://tools.ietf.org/html/rfc2234
		SP              = " ";
		HTAB            = "\t";
		CR              = "\r";
		LF              = "\n" @line;
		CRLF            = CR LF;
		WSP             = SP | HTAB;
		LWS             = ( WSP* ( CR when lookAheadWSP ) LF )? WSP+;
		SWS             = LWS?;

		UTF8_CONT       = 0x80..0xBF @append;
		UTF8_NONASCII   = 0xC0..0xDF @append UTF8_CONT {1}
		                | 0xE0..0xEF @append UTF8_CONT {2}
		                | 0xF0..0xF7 @append UTF8_CONT {3}
		                | 0xF8..0xFb @append UTF8_CONT {4}
		                | 0xFC..0xFD @append UTF8_CONT {5};
		UTF8            = 0x21..0x7F @append | UTF8_NONASCII;

		mUTF8_CONT      = 0x80..0xBF;
		mUTF8_NONASCII  = 0xC0..0xDF mUTF8_CONT {1}
		                | 0xE0..0xEF mUTF8_CONT {2}
		                | 0xF0..0xF7 mUTF8_CONT {3}
		                | 0xF8..0xFb mUTF8_CONT {4}
		                | 0xFC..0xFD mUTF8_CONT {5};
		mUTF8           = 0x21..0x7F | mUTF8_NONASCII;

		# https://tools.ietf.org/html/rfc3261#section-25.1
		reserved        = ";" | "/" | "?" | ":" | "@" | "&" | "=" | "+" | "$" | "," ;
		mark            = "-" | "_" | "." | "!" | "~" | "*" | "'" | "(" | ")" ;
		unreserved      = alnum | mark ;
		tokenc          = alnum | "-" | "." | "!" | "%" | "*" | "_" | "+" | "`"
		                | "'" | "~" ;
		separators      = "("  | ")" | "<" | ">" | "@" | "," | ";" | ":" | "\\"
		                | "\"" | "/" | "[" | "]" | "?" | "=" | "{" | "}" | SP
		                | HTAB ;
		wordc           = alnum | "-" | "." | "!" | "%" | "*" | "_" | "+" | "`"
		                | "'" | "~" | "(" | ")" | "<" | ">" | ":" | "\\" | "\""
		                | "/" | "[" | "]" | "?" | "{" | "}" ;
		schmchars       = alnum | "+" | "-" | "." ;
		word            = wordc+;
		STAR            = SWS "*" SWS;
		SLASH           = SWS "/" SWS;
		EQUAL           = SWS "=" SWS;
		LPAREN          = SWS "(" SWS;
		RPAREN          = SWS ")" SWS;
		RAQUOT          = ">" SWS;
		LAQUOT          = SWS "<";
		COMMA           = SWS "," SWS;
		SEMI            = SWS ";" SWS;
		COLON           = SWS ":" SWS;
		HCOLON          = WSP* ":" SWS;
		LDQUOT          = SWS "\"";
		RDQUOT          = "\"" SWS;
		ctext           = 0x21..0x27 | 0x2A..0x5B | 0x5D..0x7E | UTF8_NONASCII | LWS;
		quoted_pair     = "\\" ( 0x00..0x09 | 0x0B..0x0C | 0x0E..0x7F ) @append;
		comment         = LPAREN ( ctext | quoted_pair )* <: RPAREN;  # TODO(jart): Nested parens.
		qdtext          = UTF8_NONASCII | LWS @append | ( 0x21 | 0x23..0x5B | 0x5D..0x7E ) @append;
		escaped         = "%" ( xdigit @hexHi ) ( xdigit @hexLo ) ;
		uric            = reserved | unreserved | "%" | "[" | "]";
		uric_no_slash   = unreserved | escaped | ";" | "?" | ":" | "@" | "&" | "="
		                | "+" | "$" | "," ;
		uri             = alpha schmchars* ":" uric+;
		token           = tokenc+;
		tokenhost       = ( tokenc | "[" | "]" | ":" )+;
		reasonc         = UTF8_NONASCII | ( reserved | unreserved | SP | HTAB ) @append;
		reasonmc        = escaped | reasonc;
		cid             = word ( "@" word )?;

		Method          = token >mark %Method;
		SIPVersionNo    = digit+ @VersionMajor "." digit+ @VersionMinor;
		RequestURI      = ^SP+ >mark %RequestURI;
		StatusCode      = ( digit @StatusCode ) {3};
		ReasonPhrase    = reasonmc+ >start %ReasonPhrase;
		hval            = ( mUTF8 | LWS )* >mark;

		# Address Headers
		aname    = ("Contact"i | "m"i) %{addr=&msg.Contact}
		         | ("From"i | "f"i) %{addr=&msg.From}
		         | "P-Asserted-Identity"i %{addr=&msg.PAssertedIdentity}
		         | "Record-Route"i %{addr=&msg.RecordRoute}
		         | "Remote-Party-ID"i %{addr=&msg.RemotePartyID}
		         | "Route"i %{addr=&msg.Route}
		         | ("To"i | "t"i) %{addr=&msg.To}
		         ;

		# String Headers
		sname    = "Accept"i %{value=&msg.Accept}
		         | ("Accept-Contact"i | "a"i) %{value=&msg.AcceptContact}
		         | "Accept-Encoding"i %{value=&msg.AcceptEncoding}
		         | "Accept-Language"i %{value=&msg.AcceptLanguage}
		         | ("Allow"i | "u"i) %{value=&msg.Allow}
		         | ("Allow-Events"i | "u"i) %{value=&msg.AllowEvents}
		         | "Alert-Info"i %{value=&msg.AlertInfo}
		         | "Authentication-Info"i %{value=&msg.AuthenticationInfo}
		         | "Authorization"i %{value=&msg.Authorization}
		         | "Content-Disposition"i %{value=&msg.ContentDisposition}
		         | "Content-Language"i %{value=&msg.ContentLanguage}
		         | ("Content-Encoding"i | "e"i) %{value=&msg.ContentEncoding}
		         | "Call-Info"i %{value=&msg.CallInfo}
		         | "Date"i %{value=&msg.Date}
		         | "Error-Info"i %{value=&msg.ErrorInfo}
		         | ("Event"i | "o"i) %{value=&msg.Event}
		         | "In-Reply-To"i %{value=&msg.InReplyTo}
		         | "Reply-To"i %{value=&msg.ReplyTo}
		         | "MIME-Version"i %{value=&msg.MIMEVersion}
		         | "Organization"i %{value=&msg.Organization}
		         | "Priority"i %{value=&msg.Priority}
		         | "Proxy-Authenticate"i %{value=&msg.ProxyAuthenticate}
		         | "Proxy-Authorization"i %{value=&msg.ProxyAuthorization}
		         | "Proxy-Require"i %{value=&msg.ProxyRequire}
		         | ("Refer-To"i | "r"i) %{value=&msg.ReferTo}
		         | ("Referred-By"i | "b"i) %{value=&msg.ReferredBy}
		         | "Require"i %{value=&msg.Require}
		         | "Retry-After"i %{value=&msg.RetryAfter}
		         | "Server"i %{value=&msg.Server}
		         | ("Subject"i | "s"i) %{value=&msg.Subject}
		         | ("Supported"i | "k"i) %{value=&msg.Supported}
		         | "Timestamp"i %{value=&msg.Timestamp}
		         | "Unsupported"i %{value=&msg.Unsupported}
		         | "User-Agent"i %{value=&msg.UserAgent}
		         | "Warning"i %{value=&msg.Warning}
		         | "WWW-Authenticate"i %{value=&msg.WWWAuthenticate}
		         ;

		# Custom Headers
		cheader  = ("Call-ID"i | "i"i) @!gxh HCOLON cid >mark %CallID
		         | ("Content-Length"i | "l"i) @!gxh HCOLON digit+ >{clen=0} @ContentLength
		         | ("Content-Type"i | "c"i) @!gxh HCOLON <: hval %ContentType
		         | "CSeq"i @!gxh HCOLON (digit+ @CSeq) LWS token >mark %CSeqMethod
		         | ("Expires"i | "l"i) @!gxh HCOLON digit+ >{msg.Expires=0} @Expires
		         | ("Max-Forwards"i | "l"i) @!gxh HCOLON digit+ >{msg.MaxForwards=0} @MaxForwards
		         | ("Min-Expires"i | "l"i) @!gxh HCOLON digit+ >{msg.MinExpires=0} @MinExpires
		         | ("Via"i | "v"i) @!gxh HCOLON <: hval %Via
		         ;

		value   := hval <: CRLF @store_value @goto_header;
		xheader := token >mark %store_name HCOLON <: any @{value=nil;addr=nil} @goto_value;
		header  := CRLF @break
		         | cheader <: CRLF @goto_header
		         | aname >mark @err(goto_xheader) HCOLON <: any @{value=nil} @goto_value
		         | sname >mark @err(goto_xheader) HCOLON <: any @{addr=nil} @goto_value
		         ;

		SIPVersion    = "SIP/" SIPVersionNo;
		Request       = Method SP RequestURI SP SIPVersion CRLF @goto_header;
		Response      = SIPVersion SP StatusCode SP ReasonPhrase CRLF @goto_header;
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

func lookAheadWSP(data []byte, p, pe int) bool {
	return p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')
}
