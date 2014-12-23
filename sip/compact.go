// SIP Compact Header Definitions

package sip

func uncompactHeader(k string) string {
	if header, ok := compactHeaders[k]; ok {
		return header
	}
	return k
}

// http://www.cs.columbia.edu/sip/compact.html
var compactHeaders = map[string]string{
	"a": "Accept-Contact",
	"b": "Referred-By",
	"c": "Content-Type",
	"e": "Content-Encoding",
	"f": "From",
	"i": "Call-ID",
	"k": "Supported",
	"l": "Content-Length",
	"m": "Contact",
	"o": "Event",
	"r": "Refer-To",
	"s": "Subject",
	"t": "To",
	"u": "Allow-Events",
	"v": "Via",
}
