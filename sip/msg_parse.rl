// -*-go-*-

package sip

import (
	"errors"
	"fmt"
	"github.com/jart/gosip/sdp"
	"strings"
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
	fromp := &msg.From
	top := &msg.To
	viap := &msg.Via
	routep := &msg.Route
	rroutep := &msg.RecordRoute
	contactp := &msg.Contact
	paidp := &msg.PAssertedIdentity
	rpidp := &msg.RemotePartyID
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
	var addr *Addr
	var addrpp ***Addr

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

		action space {
			buf[amt] = ' '
			amt++
		}

		action erase {
			amt--
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

		action goto_avalue {
			fhold;
			fgoto avalue;
		}

		action goto_header {
			fgoto header;
		}

		action goto_svalue {
			fhold;
			fgoto svalue;
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
			} else {
				if msg.Headers == nil {
					msg.Headers = Headers{}
				}
				msg.Headers[name] = string(b)
			}
		}}

		action store_addr {
			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		}

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
		UTF8_CONT       = 0x80..0xBF @append;
		UTF8_NONASCII   = 0xC0..0xDF @append UTF8_CONT {1}
		                | 0xE0..0xEF @append UTF8_CONT {2}
		                | 0xF0..0xF7 @append UTF8_CONT {3}
		                | 0xF8..0xFb @append UTF8_CONT {4}
		                | 0xFC..0xFD @append UTF8_CONT {5};
		UTF8            = 0x21..0x7F @append | UTF8_NONASCII;
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
		schmchars       = alpha | digit | "+" | "-" | "." ;
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
		ctext           = 0x21..0x27 | 0x2A..0x5B | 0x5D..0x7E | UTF8_NONASCII | LWS;
		quoted_pair     = "\\" ( 0x00..0x09 | 0x0B..0x0C | 0x0E..0x7F ) @append;
		comment         = LPAREN ( ctext | quoted_pair )* <: RPAREN;  # TODO(jart): Nested parens.
		qdtext          = UTF8_NONASCII | LWS @append | ( 0x21 | 0x23..0x5B | 0x5D..0x7E ) @append;
		escaped         = "%" ( xdigit @hexHi ) ( xdigit @hexLo ) ;
		uric            = reserved | unreserved | "%" ;
		uric_no_slash   = unreserved | escaped | ";" | "?" | ":" | "@" | "&" | "="
		                | "+" | "$" | "," ;
		uri             = alpha schmchars* ":" uric+;
		token           = tokenc+;
		tokenhost       = ( tokenc | "[" | "]" | ":" )+;
		reasonc         = UTF8_NONASCII | ( reserved | unreserved | SP | HTAB ) @append;
		reasonmc        = escaped | reasonc;
		cid             = word ( "@" word )?;

		quoted_content  = ( qdtext | quoted_pair )* >start;
		quoted_string   = LDQUOT quoted_content <: RDQUOT;
		param_name      = token >mark %store_name;
		param_content   = tokenhost >start @append;
		param_value     = param_content | quoted_string;
		param           = ( param_name ( EQUAL param_value )? ) %addr_param;
		display         = ( token @append LWS %space )* >start;
		display_name    = ( display | quoted_string ) %addr_display;
		addr_spec       = uri >mark %addr_uri;
		name_addr       = display_name? LAQUOT addr_spec RAQUOT;
		addr_omg        = ( ( addr_spec -- ";" ) | name_addr ) ( SEMI param )*;
		addr            = addr_omg >new_addr %store_addr;

		Method          = token >mark %Method;
		SIPVersionNo    = digit+ @VersionMajor "." digit+ @VersionMinor;
		RequestURI      = ^SP+ >mark %RequestURI;
		StatusCode      = ( digit @StatusCode ) {3};
		ReasonPhrase    = reasonmc+ >start %ReasonPhrase;
		hval            = ( UTF8 | LWS )* >mark;

		# Address Headers
		aname    = ("Contact"i | "m"i) %{addrpp=&contactp}
		         | ("From"i | "f"i) %{addrpp=&fromp}
		         | "P-Asserted-Identity"i %{addrpp=&paidp}
		         | "Record-Route"i %{addrpp=&rroutep}
		         | "Remote-Party-ID"i %{addrpp=&rpidp}
		         | "Route"i %{addrpp=&routep}
		         | ("To"i | "t"i) %{addrpp=&top}
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

		avalue  := addr ( COMMA addr )* CR LF @goto_header;
		svalue  := hval CR LF @store_value @goto_header;
		xheader := token >mark %store_name HCOLON <: any @{value=nil} @goto_svalue;
		header  := CR LF @break
		         | cheader CR LF @goto_header
		         | aname >mark @err(goto_xheader) HCOLON <: any @goto_avalue
		         | sname >mark @err(goto_xheader) HCOLON <: any @goto_svalue
		         ;

		SIPVersion    = "SIP/" SIPVersionNo;
		Request       = Method SP RequestURI SP SIPVersion CR LF @goto_header;
		Response      = SIPVersion SP StatusCode SP ReasonPhrase CR LF @goto_header;
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
