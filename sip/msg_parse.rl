package sip

import (
//  "bytes"
  "errors"
  "fmt"
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
  cs := 0
  p := 0
  pe := len(data)
  line := 1
  linep := 0
  buf := make([]byte, len(data))
  amt := 0
  mark := 0
  clen := 0
//  var b1 string
  var hex byte

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

    action extHeaderName {
      b1 = string(bytes.ToLower(data[mark:p]))
    }

    action extHeaderValue {
      if msg.Headers == nil {
        msg.Headers = Headers{}
      }
      msg.Headers[b1] = string(data[mark:p])
    }

    action Accept {
      msg.Accept = string(data[mark:p])
    }

    action AcceptValue {
      msg.AcceptContact = string(data[mark:p])
    }

    action AcceptEncoding {
      msg.AcceptEncoding = string(data[mark:p])
    }

    action AcceptLanguage {
      msg.AcceptLanguage = string(data[mark:p])
    }

    action Allow {
      msg.Allow = string(data[mark:p])
    }

    action AllowEvents {
      msg.AllowEvents = string(data[mark:p])
    }

    action AlertInfo {
      msg.AlertInfo = string(data[mark:p])
    }

    action AuthenticationInfo {
      msg.AuthenticationInfo = string(data[mark:p])
    }

    action Authorization {
      msg.Authorization = string(data[mark:p])
    }

    action CallID {
      msg.CallID = string(data[mark:p])
    }

    action Contact {
      msg.Contact, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    }

    action ContentDisposition {
      msg.ContentDisposition = string(data[mark:p])
    }

    action ContentLanguage {
      msg.ContentLanguage = string(data[mark:p])
    }

    action ContentLength {
      clen = clen * 10 + (int(fc) - 0x30)
    }

    action ContentEncoding {
      msg.ContentEncoding = string(data[mark:p])
    }

    action ContentType {
      msg.ContentType = string(data[mark:p])
    }

    action CSeq {
      msg.CSeq = msg.CSeq * 10 + (int(fc) - 0x30)
    }

    action CSeqMethod {
      msg.CSeqMethod = string(data[mark:p])
    }

    action CallInfo {
      msg.CallInfo = string(data[mark:p])
    }

    action Date {
      msg.Date = string(data[mark:p])
    }

    action ErrorInfo {
      msg.ErrorInfo = string(data[mark:p])
    }

    action Event {
      msg.Event = string(data[mark:p])
    }

    action Expires {
      msg.Expires = msg.Expires * 10 + (int(fc) - 0x30)
    }

    action From {
      msg.From, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    }

    action InReplyTo {
      msg.InReplyTo = string(data[mark:p])
    }

    action MaxForwardsZero {
      msg.MaxForwards = 0
    }

    action MaxForwards {
      msg.MaxForwards = msg.MaxForwards * 10 + (int(fc) - 0x30)
    }

    action MinExpires {
      msg.MinExpires = msg.MinExpires * 10 + (int(fc) - 0x30)
    }

    action ReplyTo {
      msg.ReplyTo = string(data[mark:p])
    }

    action MIMEVersion {
      msg.MIMEVersion = string(data[mark:p])
    }

    action Organization {
      msg.Organization = string(data[mark:p])
    }

    action PAssertedIdentity {
      msg.PAssertedIdentity, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    }

    action Priority {
      msg.Priority = string(data[mark:p])
    }

    action ProxyAuthenticate {
      msg.ProxyAuthenticate = string(data[mark:p])
    }

    action ProxyAuthorization {
      msg.ProxyAuthorization = string(data[mark:p])
    }

    action ProxyRequire {
      msg.ProxyRequire = string(data[mark:p])
    }

    action RecordRoute {
      msg.RecordRoute, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    }

    action ReferTo {
      msg.ReferTo = string(data[mark:p])
    }

    action ReferredBy {
      msg.ReferredBy = string(data[mark:p])
    }

    action RemotePartyID {
      msg.RemotePartyID, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    }

    action Require {
      msg.Require = string(data[mark:p])
    }

    action RetryAfter {
      msg.RetryAfter = string(data[mark:p])
    }

    action Route {
      msg.Route, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    }

    action Server {
      msg.Server = string(data[mark:p])
    }

    action Subject {
      msg.Subject = string(data[mark:p])
    }

    action Supported {
      msg.Supported = string(data[mark:p])
    }

    action Timestamp {
      msg.Timestamp = string(data[mark:p])
    }

    action To {
      msg.To, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    }

    action Unsupported {
      msg.Unsupported = string(data[mark:p])
    }

    action UserAgent {
      msg.UserAgent = string(data[mark:p])
    }

    action Via {
      msg.Via, err = ParseVia(string(data[mark:p]))
      if err != nil { return nil, err }
    }

    action Warning {
      msg.Warning = string(data[mark:p])
    }

    action WWWAuthenticate {
      msg.WWWAuthenticate = string(data[mark:p])
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
    HCOLON          = WSP* ":" SWS ;
    tokenc          = alpha | digit | "-" | "." | "!" | "%" | "*" | "_"
                    | "+" | "`" | "'" | "~" ;
    token           = tokenc+ >mark ;
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
    reasonc         = reserved | unreserved | UTF8_NONASCII | SP | HTAB ;
    reasonmc        = escaped | ( reasonc @append ) ;
    cid             = word ( "@" word )? ;

    Method          = token %Method;
    SIPVersionNo    = digit+ @VersionMajor "." digit+ @VersionMinor;
    RequestURI      = ^SP+ >mark %RequestURI;
    StatusCode      = ( digit @StatusCode ) {3};
    ReasonPhrase    = reasonmc+ >start %ReasonPhrase;
    hval            = ( UTF8 | LWS )* >mark;
    extHeader       = token %extHeaderName HCOLON hval %extHeaderValue;

    # http://www.iana.org/assignments/sip-parameters/sip-parameters.xhtml
    stdHeader = "Accept"i HCOLON hval %Accept
              | ("Accept-Contact"i | "a"i) HCOLON hval %AcceptValue
              | "Accept-Encoding"i HCOLON hval %AcceptEncoding
              | "Accept-Language"i HCOLON hval %AcceptLanguage
              | ("Allow"i | "u"i) HCOLON hval %Allow
              | ("Allow-Events"i | "u"i) HCOLON hval %AllowEvents
              | "Alert-Info"i HCOLON hval %AlertInfo
              | "Authentication-Info"i HCOLON hval %AuthenticationInfo
              | "Authorization"i HCOLON hval %Authorization
              | ("Call-ID"i | "i"i) HCOLON cid >mark %CallID
              | ("Contact"i | "m"i) HCOLON hval %Contact
              | "Content-Disposition"i HCOLON hval %ContentDisposition
              | "Content-Language"i HCOLON hval %ContentLanguage
              | ("Content-Length"i | "l"i) HCOLON digit+ @ContentLength
              | ("Content-Encoding"i | "e"i) HCOLON hval %ContentEncoding
              | ("Content-Type"i | "c"i) HCOLON hval %ContentType
              | "CSeq"i HCOLON (digit+ @CSeq) LWS token >mark %CSeqMethod
              | "Call-Info"i HCOLON hval %CallInfo
              | "Date"i HCOLON hval %Date
              | "Error-Info"i HCOLON hval %ErrorInfo
              | ("Event"i | "o"i) HCOLON hval %Event
              | ("Expires"i | "l"i) HCOLON digit+ @Expires
              | ("From"i | "f"i) HCOLON hval %From
              | "In-Reply-To"i HCOLON hval %InReplyTo
              | ("Max-Forwards"i | "l"i) HCOLON digit+ >MaxForwardsZero @MaxForwards
              | ("Min-Expires"i | "l"i) HCOLON digit+ @MinExpires
              | "Reply-To"i HCOLON hval %ReplyTo
              | "MIME-Version"i HCOLON hval %MIMEVersion
              | "Organization"i HCOLON hval %Organization
              | "P-Asserted-Identity"i HCOLON hval %PAssertedIdentity
              | "Priority"i HCOLON hval %Priority
              | "Proxy-Authenticate"i HCOLON hval %ProxyAuthenticate
              | "Proxy-Authorization"i HCOLON hval %ProxyAuthorization
              | "Proxy-Require"i HCOLON hval %ProxyRequire
              | "Record-Route"i HCOLON hval %RecordRoute
              | ("Refer-To"i | "r"i) HCOLON hval %ReferTo
              | ("Referred-By"i | "b"i) HCOLON hval %ReferredBy
              | "Remote-Party-ID"i HCOLON hval %RemotePartyID
              | "Require"i HCOLON hval %Require
              | "Retry-After"i HCOLON hval %RetryAfter
              | "Route"i HCOLON hval %Route
              | "Server"i HCOLON hval %Server
              | ("Subject"i | "s"i) HCOLON hval %Subject
              | ("Supported"i | "k"i) HCOLON hval %Supported
              | "Timestamp"i HCOLON hval %Timestamp
              | ("To"i | "t"i) HCOLON hval %To
              | "Unsupported"i HCOLON hval %Unsupported
              | "User-Agent"i HCOLON hval %UserAgent
              | ("Via"i | "v"i) HCOLON hval %Via
              | "Warning"i HCOLON hval %Warning
              | "WWW-Authenticate"i HCOLON hval %WWWAuthenticate
              ;

    header        = stdHeader CRLF;
    headers       = header* CR LF @break;
    SIPVersion    = "SIP/" SIPVersionNo;
    RequestLine   = Method SP RequestURI SP SIPVersion CRLF;
    StatusLine    = SIPVersion SP StatusCode SP ReasonPhrase CRLF;
    Request       = RequestLine headers;
    Response      = StatusLine headers;
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
    msg.Payload = string(data[p:len(data)])
  }

  return msg, nil
}
