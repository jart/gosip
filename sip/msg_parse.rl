// -*-go-*-
//
// Ragel SIP Message Parser
//
// This file is compiled into Go code by the Ragel State Machine Compiler for
// the purpose of converting SIP messages into a Msg data structure. This
// machine works in tandem with the Ragel machine defined in uri_parse.rl.
//
// Perhaps it would have been better if the authors of this protocol had chosen
// to use a binary serialization format like protocol buffers. But instead they
// chose to create a plaintext protocol that looks similar to HTTP requests,
// but are phenomenally more complicated.
//
// SIP messages are quite insane.
//
//   o Whitespace can be used liberally in a variety of different ways.
//
//     - Via host:port can have whitespace, e.g. "host \t:  port"
//
//   o UTF-8 is supported in some places but not others.
//
//   o Headers can span multiple lines.
//
//   o Header values can contain comments, e.g. Message: lol (i'm (hidden))
//
//   o Header names are case-insensitive and have shorthand notation.
//
//   o There's ~50 standard headers, many of which have custom parsing rules.
//
//   o URIs can have ;params;like=this
//
//     - Params can belong either to a URI or Addr object, e.g. <sip:uri;param>
//       cf. <sip:uri>;param
//
//     - Addresses may omit angle brackets, in which case params belong to the
//       Addr object.
//
//     - URI params ;are=escaped%20like%22this but params belonging to Addr
//       ;are="escaped like\"this"
//
//     - Backslash escaping is not like C, e.g. \t\n -> tn
//
//     - Address display name can have whitespace without quotes, which is
//       collapsed. Quoted form is not collapsed.
//
//   o Via and address headers can be repeated in two ways: repeating the
//     header, using commas within a single header, or both.
//
// See: http://www.colm.net/files/ragel/ragel-guide-6.9.pdf
// See: http://zedshaw.com/archive/ragel-state-charts/

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
	var via *Via

	%%{
		action hold {
			fhold;
		}

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

		action Via {
			*viap = via
			viap = &via.Next
			// via = nil
		}

		action ViaProtocol {
			via.Protocol = string(data[mark:p])
		}

		action ViaVersion {
			via.Version = string(data[mark:p])
		}

		action ViaTransport {
			via.Transport = string(data[mark:p])
		}

		action ViaHost {
			via.Host = string(data[mark:p])
		}

		action ViaPort {
			via.Port = via.Port * 10 + (uint16(fc) - 0x30)
		}

		action ViaParam {
			if via.Params == nil {
				via.Params = Params{}
			}
			via.Params[name] = string(buf[0:amt])
		}

		action goto_header {
			fgoto header;
		}

		action goto_value {
			fgoto value;
		}

		action goto_via {
			via = new(Via)
			fgoto via;
		}

		action goto_via_port {
			fgoto via_port;
		}

		action goto_via_pname {
			amt = 0  // Needed so ViaParam action works when there's no value.
			fgoto via_pname;
		}

		action goto_via_pvalue {
			fgoto via_pvalue;
		}

		action gxh {
			fhold;
			fgoto xheader;
		}

		action name {
			name = string(data[mark:p])
		}

		action value {{
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

		action lookAheadWSP { lookAheadWSP(data, p, pe) }

		# https://tools.ietf.org/html/rfc2234
		SP              = " ";
		HTAB            = "\t";
		CR              = "\r";
		LF              = "\n" @line;
		DQUOTE          = "\"";
		CRLF            = CR LF;
		WSP             = SP | HTAB;
		LWS             = ( WSP* ( CR when lookAheadWSP ) LF )? WSP+;
		SWS             = LWS?;

		LWSCRLF_append  = ( CR when lookAheadWSP ) @append LF @append;
		LWS_append      = ( WSP* @append LWSCRLF_append )? WSP+ @append;

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
		escaped         = "%" ( xdigit @hexHi ) ( xdigit @hexLo ) ;
		ipv4            = digit | "." ;
		ipv6            = xdigit | "." | ":" ;
		hostname        = alpha | digit | "-" | "." ;
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

		# Quoted strings can have just about anything, including backslash escapes,
		# which aren't quite as fancy as the ones you'd see in programming.
		qdtextc         = 0x21 | 0x23..0x5B | 0x5D..0x7E;
		qdtext          = UTF8_NONASCII | LWS_append | qdtextc @append;
		quoted_pair     = "\\" ( 0x00..0x09 | 0x0B..0x0C | 0x0E..0x7F ) @append;
		quoted_content  = ( qdtext | quoted_pair )* >start;
		quoted_string   = SWS DQUOTE quoted_content DQUOTE;

		# Via Parsing
		#
		# Parsing these is kind of difficult because infinite whitespace is allowed
		# between colons, semicolons, commas, and don't forget that lines can
		# continue. So we're going to break things down into four separate machines
		# that jump between each other.
		ViaProtocol     = token >mark %ViaProtocol;
		ViaVersion      = token >mark %ViaVersion;
		ViaTransport    = token >mark %ViaTransport;
		ViaSent         = ViaProtocol SLASH ViaVersion SLASH ViaTransport;
		ViaHostIPv4     = ( digit | "." )+ >mark %ViaHost;
		ViaHostIPv6     = "[" ( xdigit | "." | ":" )+ >mark %ViaHost "]";
		ViaHostName     = ( alnum | "." | "-" )+ >mark %ViaHost;
		ViaHost         = ViaHostIPv4 | ViaHostIPv6 | ViaHostName;
		ViaPort         = digit+ @ViaPort;
		ViaParamName    = token >mark %name;
		ViaParamContent = tokenhost >start @append;
		ViaParamValue   = ViaParamContent | quoted_string;
		via_pvalue_end  = ( CR when !lookAheadWSP ) LF @ViaParam @Via @goto_header
		                | SEMI <: any @hold @ViaParam @goto_via_pname
		                | COMMA <: any @hold @ViaParam @Via @goto_via;
		via_pvalue     := ViaParamValue via_pvalue_end;
		via_pname_end   = ( CR when !lookAheadWSP ) LF @ViaParam @Via @goto_header
		                | EQUAL <: any @hold @goto_via_pvalue
		                | SEMI <: any @hold @ViaParam @goto_via_pname
		                | COMMA <: any @hold @ViaParam @Via @goto_via;
		via_pname      := ViaParamName via_pname_end;
		via_port_end    = ( CR when !lookAheadWSP ) LF @Via @goto_header
		                | SEMI <: any @hold @goto_via_pname
		                | COMMA <: any @hold @Via @goto_via;
		via_port       := ViaPort via_port_end;
		via_end         = ( CR when !lookAheadWSP ) LF @Via @goto_header
		                | COLON <: any @hold @goto_via_port
		                | SEMI <: any @hold @goto_via_pname
		                | COMMA <: any @hold @Via @goto_via;
		via            := ViaSent LWS ViaHost via_end;

		# Address Header Name Definitions
		#
		# These headers set the addr pointer to tell the 'value' machine where to
		# store the value after using ParseAddrBytes().
		aname    = ("Contact"i | "m"i) %{addr=&msg.Contact}
		         | ("From"i | "f"i) %{addr=&msg.From}
		         | "P-Asserted-Identity"i %{addr=&msg.PAssertedIdentity}
		         | "Record-Route"i %{addr=&msg.RecordRoute}
		         | "Remote-Party-ID"i %{addr=&msg.RemotePartyID}
		         | "Route"i %{addr=&msg.Route}
		         | ("To"i | "t"i) %{addr=&msg.To}
		         ;

		# String Header Name Definitions
		#
		# These headers set the value pointer to tell the 'value' machine where to
		# store the resulting token string.
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

		# Custom Header Definitions
		#
		# These headers do not jump to the 'value' machine, but instead specify
		# their own special type of parsing.
		cheader  = ("Call-ID"i | "i"i) $!gxh HCOLON cid >mark %CallID
		         | ("Content-Length"i | "l"i) $!gxh HCOLON digit+ >{clen=0} @ContentLength
		         | ("Content-Type"i | "c"i) $!gxh HCOLON <: hval %ContentType
		         | "CSeq"i $!gxh HCOLON (digit+ @CSeq) LWS token >mark %CSeqMethod
		         | ("Expires"i | "l"i) $!gxh HCOLON digit+ >{msg.Expires=0} @Expires
		         | ("Max-Forwards"i | "l"i) $!gxh HCOLON digit+ >{msg.MaxForwards=0} @MaxForwards
		         | ("Min-Expires"i | "l"i) $!gxh HCOLON digit+ >{msg.MinExpires=0} @MinExpires
		         ;

		# Header Parsing
		#
		# The header machine parses a single header and then jumps to itself to
		# loop. When the final CRLF is observed, we then break out of the Ragel
		# parser and let the Go code handle payload extraction.
		#
		# Parsing standard header names is a prefix trie search in generated code.
		# Lookahead to set the mark on the header name. In order to support
		# extended headers, we'll use $!gxh to jump to the xheader machine when an
		# unrecognized character is detected in the header name.
		#
		# An independent machine has been created for generic header values, so
		# that it doesn't need to be duplicated for each leaf in the prefix
		# trie. When the value machine has finished reading a value, it'll be
		# parsed and stored based on whether the value/addr pointers are set.
		#
		# Header values can span multiple lines. Lookahead is used in the LWS
		# definition to check for whitespace at the start of the next line upon
		# encountering a line feed character, in order to determine if a line
		# continuation is present.
		#
		# In order to concatenate across machines, we use lookahead in conjunction
		# with the left-guarded concatenation operator. This pattern works is
		# defined as follows: `foo <: any @hold @goto_bar`.
		#
		# Header names are case insensitive. Each recognized header is assigned to
		# a specific field in the Msg data structure. Extended headers are stored
		# to a linked list data structure with the casing preserved. This is so
		# messages can be reproduced with roughly the same appearance. It is the
		# responsibility of the person using Msg.Headers to do case-insensitive
		# string comparisons.
		value   := hval <: CRLF @value @goto_header;
		xheader := token %name HCOLON <: any @{value=nil;addr=nil} @hold @goto_value;
		sheader  = cheader <: CRLF @goto_header
		         | aname $!gxh HCOLON <: any @{value=nil} @hold @goto_value
		         | sname $!gxh HCOLON <: any @{addr=nil} @hold @goto_value
		         | ("Via"i | "v"i) $!gxh HCOLON <: any @hold @goto_via;
		header  := CRLF @break
		         | tokenc @mark @hold sheader;

		# Start Line Parsing
		#
		# The Request and Response definitions are very straightforward, and the
		# main machine is the union of the two. Once the line feed character has
		# been observed, we then jump to the header machine.
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
