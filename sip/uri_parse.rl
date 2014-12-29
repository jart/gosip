package sip

import (
  "bytes"
  "errors"
  "fmt"
  "strconv"
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
    action strStart   { amt = 0                          }
    action strChar    { buf[amt] = fc; amt++             }
    action strLower   { buf[amt] = fc + 0x20; amt++      }
    action hexHi      { hex = unhex(fc) * 16             }
    action hexLo      { hex += unhex(fc)
                        buf[amt] = hex; amt++            }
    action user       { uri.User = string(buf[0:amt])    }
    action pass       { uri.Pass = string(buf[0:amt])    }
    action host       { uri.Host = string(buf[0:amt])    }
    action scheme     { uri.Scheme = string(buf[0:amt])  }
    action b1         { b1 = string(buf[0:amt]); amt = 0 }
    action b2         { b2 = string(buf[0:amt]); amt = 0 }

    action port {
      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
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

    # TODO(jart): Use BNF from SIP RFC: https://tools.ietf.org/html/rfc3261#section-25.1

    # Define what a single character is allowed to be.
    toxic         = ( cntrl | 127 ) ;
    scary         = ( toxic | space | "\"" | "#" | "%" | "<" | ">" | "=" ) ;
    schmchars     = ( lower | digit | "+" | "-" | "." ) ;
    authdelims    = ( "/" | "?" | "#" | ":" | "@" | ";" | "[" | "]" | "&" ) ;
    userchars     = any -- ( authdelims | scary ) ;
    passchars     = userchars ;
    hostchars     = passchars -- upper;
    hostcharsEsc  = ( hostchars | ":" ) -- upper;
    portchars     = digit ;
    paramchars    = userchars -- upper ;
    headerchars   = userchars ;

    # Define how characters trigger actions.
    escape        = "%" xdigit xdigit ;
    unescape      = "%" ( xdigit @hexHi ) ( xdigit @hexLo ) ;
    schmfirst     = ( upper @strLower ) | ( lower @strChar ) ;
    schmchar      = ( upper @strLower ) | ( schmchars @strChar ) ;
    userchar      = unescape | ( userchars @strChar ) ;
    passchar      = unescape | ( passchars @strChar ) ;
    hostchar      = unescape | ( upper @strLower ) | ( hostchars @strChar ) ;
    hostcharEsc   = unescape | ( upper @strLower ) | ( hostcharsEsc @strChar ) ;
    portchar      = unescape | ( portchars @strChar ) ;
    paramchar     = unescape | ( upper @strLower ) | ( paramchars @strChar ) ;
    headerchar    = unescape | ( headerchars @strChar ) ;

    # Define multi-character patterns.
    scheme        = ( schmfirst schmchar* ) >strStart %scheme ;
    user          = userchar+ >strStart %user ;
    pass          = passchar+ >strStart %pass ;
    hostPlain     = hostchar+ >strStart %host ;
    hostQuoted    = "[" ( hostcharEsc+ >strStart %host ) "]" ;
    host          = hostQuoted | hostPlain ;
    port          = portchar* >strStart %port ;
    paramkey      = paramchar+ >strStart >b2 %b1 ;
    paramval      = paramchar+ >strStart %b2 ;
    param         = space* ";" paramkey ( "=" paramval )? %param ;
    headerkey     = headerchar+ >strStart >b2 %b1 ;
    headerval     = headerchar+ >strStart %b2 ;
    header        = headerkey ( "=" headerval )? %header ;
    headers       = "?" header ( "&" header )* ;
    userpass      = user ( ":" pass )? ;
    hostport      = host ( ":" port )? ;
    uriSansUser  := space* scheme ":" hostport param* space* headers? space* ;
    uriWithUser  := space* scheme ":" userpass "@" hostport param* space* headers? space* ;
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
      return nil, errors.New(fmt.Sprintf("Unexpected EOF: %s", data))
    } else {
      return nil, errors.New(fmt.Sprintf("Error in URI at pos %d: %s", p, data))
    }
  }
  return uri, nil

fail:
  return nil, errors.New(fmt.Sprintf("Bad URI: %s", data))
}
