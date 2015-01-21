// -*-go-*-

package sip

import (
	"bytes"
	"errors"
	"fmt"
)

%% machine uri;
%% write data;

// ParseURI turns a a SIP URI into a data structure.
func ParseURI(s string) (uri *URI, err error) {
	if s == "" {
		return nil, errors.New("Empty URI")
	}
	return ParseURIBytes([]byte(s))
}

// ParseURI turns a a SIP URI byte slice into a data structure.
func ParseURIBytes(data []byte) (uri *URI, err error) {
	if data == nil {
		return nil, nil
	}
	uri = new(URI)
	cs := 0
	p := 0
	pe := len(data)
	eof := len(data)
	buf := make([]byte, len(data))
	amt := 0
	var b1, b2 string
	var hex byte

	%%{
		action start {
			amt = 0
		}

		action append {
			buf[amt] = fc
			amt++
		}

		action hexHi {
			hex = unhex(fc) * 16
		}

		action hexLo {
			hex += unhex(fc)
			buf[amt] = hex
			amt++
		}

		action scheme {
			uri.Scheme = string(buf[0:amt])
		}

		action user {
			uri.User = string(buf[0:amt])
		}

		action pass {
			uri.Pass = string(buf[0:amt])
		}

		action host {
			uri.Host = string(buf[0:amt])
		}

		action port {
			uri.Port = uri.Port * 10 + uint16(fc - 0x30)
		}

		action b1 {
			b1 = string(buf[0:amt])
			amt = 0
		}

		action b2 {
			b2 = string(buf[0:amt])
			amt = 0
		}

		action lower {
			if 'A' <= fc && fc <= 'Z' {
				buf[amt] = fc + 0x20
			} else {
				buf[amt] = fc
			}
			amt++
		}

		action param {
			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		}

		action header {
			if uri.Headers == nil {
				uri.Headers = URIHeaders{}
			}
			uri.Headers[b1] = b2
		}

		# Byte character definitions.
		mark            = "-" | "_" | "." | "!" | "~" | "*" | "'" | "(" | ")" ;
		reserved        = ";" | "/" | "?" | ":" | "@" | "&" | "=" | "+" | "$" | "," ;
		unreserved      = alpha | digit | mark ;
		ipv4            = digit | "." ;
		ipv6            = xdigit | "." | ":" ;
		hostname        = alpha | digit | "-" | "." ;
		tel             = digit | "+" | "-" ;
		schmchars       = alpha | digit | "+" | "-" | "." ;
		userchars       = unreserved | "&" | "=" | "+" | "$" | "," | ";" | "?" | "/" ;
		passchars       = unreserved | "&" | "=" | "+" | "$" | "," ;
		paramchars      = unreserved | "[" | "]" | "/" | ":" | "&" | "+" | "$" ;
		headerchars     = unreserved | "[" | "]" | "/" | "?" | ":" | "+" | "$" ;

		# Multibyte character definitions.
		escaped         = "%" ( xdigit @hexHi ) ( xdigit @hexLo ) ;
		userchar        = escaped | ( userchars @append ) ;
		passchar        = escaped | ( passchars @append ) ;
		paramchar       = escaped | ( paramchars @lower ) ;
		headerchar      = escaped | ( headerchars @append ) ;

		# URI component definitions.
		scheme          = ( alpha schmchars* ) >start @lower %scheme ;
		user            = userchar+ >start %user ;
		pass            = passchar+ >start %pass ;
		host6           = "[" ( ipv6+ >start @lower %host ) "]" ;
		host            = host6 | ( ( ipv4 | hostname | tel )+ >start @lower %host ) ;
		port            = digit+ @port ;
		paramkey        = paramchar+ >start >b2 %b1 ;
		paramval        = paramchar+ >start %b2 ;
		param           = space* ";" paramkey ( "=" paramval )? %param ;
		headerkey       = headerchar+ >start >b2 %b1 ;
		headerval       = headerchar+ >start %b2 ;
		header          = headerkey ( "=" headerval )? %header ;
		headers         = "?" header ( "&" header )* ;
		userpass        = user ( ":" pass )? ;
		hostport        = host ( ":" port )? ;
		uriSansUser    := space* scheme ":" hostport param* space* headers? space* ;
		uriWithUser    := space* scheme ":" userpass "@" hostport param* space* headers? space* ;
	}%%

	%% write init;
	if bytes.IndexByte(data, '@') == -1 {
		cs = uri_en_uriSansUser;
	} else {
		cs = uri_en_uriWithUser;
	}
	%% write exec;

	if cs < uri_first_final {
		if p == pe {
			return nil, errors.New(fmt.Sprintf("Incomplete URI: %s", data))
		} else {
			return nil, errors.New(fmt.Sprintf("Error in URI at pos %d: %s", p, data))
		}
	}
	return uri, nil
}
