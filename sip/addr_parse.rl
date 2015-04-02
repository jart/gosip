// -*-go-*-

package sip

import (
	"errors"
	"fmt"
	"strings"
)

%% machine addr;
%% write data;

// ParseAddr turns a SIP address into a data structure.
func ParseAddr(s string) (addr *Addr, err error) {
	if s == "" {
		return nil, errors.New("Empty SIP message")
	}
	return ParseAddrBytes([]byte(s), nil)
}

// ParseAddr turns a SIP address slice into a data structure.
func ParseAddrBytes(data []byte, next *Addr) (addr *Addr, err error) {
	if data == nil {
		return nil, nil
	}
	addr = new(Addr)
	result := addr
	cs := 0
	p := 0
	pe := len(data)
	eof := len(data)
	buf := make([]byte, len(data))
	amt := 0
	mark := 0
	var name string

	%%{
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

		action name {
			name = string(data[mark:p])
		}

		action display {
			// TODO: Collapse this string.
			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		}

		action qdisplay {
			addr.Display = string(buf[0:amt])
		}

		action uri {
			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		}

		action param {
			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		}

		action lookAheadWSP { p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t') }

		action goto_name_addr {
			p = mark
			fhold;
			fgoto name_addr;
		}

		action goto_bare_addr {
			p = mark
			fhold;
			fgoto bare_addr;
		}

		action goto_params {
			fhold;
			fgoto params;
		}

		action goto_addr {
			addr.Next = new(Addr)
			addr = addr.Next
			fhold;
			fgoto addr;
		}

		# https://tools.ietf.org/html/rfc2234
		SP              = " ";
		HTAB            = "\t";
		CR              = "\r";
		LF              = "\n";
		DQUOTE          = "\"";
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

		LWSCRLF_append  = ( CR when lookAheadWSP ) @append LF @append;
		LWS_append      = ( WSP* @append LWSCRLF_append )? WSP+ @append;

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

		# The RFC BNF for SIP addresses is a bit nuts; we'll try our best.
		token           = tokenc+;
		tokenhost       = ( tokenc | "[" | "]" | ":" )+;
		schemec         = alnum | "+" | "-" | ".";
		scheme          = alpha schemec*;
		uric            = reserved | unreserved | "%" | "[" | "]";
		uri             = scheme ":" uric+;

		# Quoted strings can have just about anything, including backslash escapes,
		# which aren't quite as fancy as the ones you'd see in programming.
		qdtextc         = 0x21 | 0x23..0x5B | 0x5D..0x7E;
		qdtext          = UTF8_NONASCII | LWS_append | qdtextc @append;
		quoted_pair     = "\\" ( 0x00..0x09 | 0x0B..0x0C | 0x0E..0x7F ) @append;
		quoted_content  = ( qdtext | quoted_pair )* >start;
		quoted_string   = SWS DQUOTE quoted_content DQUOTE;

		# The display name can either be quoted or unquoted. The unquoted form is
		# much more limited in what it may contain, although it does permit spaces.
		display_name    = quoted_string %qdisplay
		                | SWS ( token @append LWS_append )* >start %display
		                ;

		# Parameter parsing is pretty straightforward. Unlike URI parameters, the
		# values on these can be quoted. The only thing that's weird is if the URI
		# doesn't appear in angle brackets, the parameters need to be owned by the
		# Addr object, not the URI. This has been placed in its own machine so it
		# doesn't have to be duplicated in the two machines below.
		param_name      = token >mark >start %name;
		param_content   = tokenhost @append;
		param_value     = param_content | quoted_string;
		param           = ( param_name ( EQUAL param_value )? ) %param;
		params         := ( SEMI param )* ( COMMA <: any @goto_addr )?;

		# Now we're going to define two separate machines. The first is for
		# addresses with angle brackets and the second is for ones without.
		addr_spec       = uri >mark %uri;
		name_addr      := display_name? LAQUOT addr_spec RAQUOT <: ( any @goto_params )?;
		bare_addr      := ( addr_spec -- ";" ) <: ( any @goto_params )?;

		# Now we perform lookahead to determine which machine we should use.
		look            = [^;<]
		                | ";" @goto_bare_addr
		                | "<" @goto_name_addr
		                ;
		addr           := look+ >mark $eof(goto_name_addr);

		write init;
		write exec;
	}%%

	if cs < addr_first_final {
		if p == pe {
			return nil, errors.New(fmt.Sprintf("Incomplete SIP address: %s", data))
		} else {
			return nil, errors.New(fmt.Sprintf("Error in SIP address at offset %d: %s", p, string(data)))
		}
	}

	if next != nil {
		next.Last().Next = result
		return next, nil
	}
	return result, nil
}
