// Standard SIP Protocol Messages
//
// http://www.iana.org/assignments/sip-parameters

package sip

var Phrases = map[int]string{
	// 1xx: Provisional -- request received, continuing to process the
	//      request;
	100: "Trying",  // indicates server is not totally pwnd
	180: "Ringing", // i promise the remote phone is ringing
	181: "Call Is Being Forwarded",
	182: "Queued",
	183: "Session Progress", // establish early media (pstn ringback)

	// 2xx: Success -- the action was successfully received, understood,
	//      and accepted;
	200: "OK",              // call is answered
	202: "Accepted",        // [RFC3265]
	204: "No Notification", // [RFC5839]

	// 3xx: Redirection -- further action needs to be taken in order to
	//      complete the request;
	300: "Multiple Choices",
	301: "Moved Permanently",
	302: "Moved Temporarily", // send your call there instead kthx
	305: "Use Proxy",         // you fool! send your call there instead
	380: "Alternative Service",

	// 4xx: Client Error -- the request contains bad syntax or cannot be
	//      fulfilled at this server;
	400: "Bad Request",        // missing headers, bad format, etc.
	401: "Unauthorized",       // resend request with auth header
	402: "Payment Required",   // i am greedy
	403: "Forbidden",          // gtfo
	404: "Not Found",          // wat?
	405: "Method Not Allowed", // i don't support that type of request
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Timeout",
	409: "Conflict",
	410: "Gone", // shaniqua don't live here no more
	411: "Length Required",
	412: "Conditional Request Failed", // [RFC3903]
	413: "Request Entity Too Large",
	414: "Request-URI Too Long",
	415: "Unsupported Media Type",
	416: "Unsupported URI Scheme",
	417: "Unknown Resource-Priority",
	420: "Bad Extension",
	421: "Extension Required",
	422: "Session Interval Too Small", // [RFC4028]
	423: "Interval Too Brief",
	428: "Use Identity Header",              // [RFC4474]
	429: "Provide Referrer Identity",        // [RFC3892]
	430: "Flow Failed",                      // [RFC5626]
	433: "Anonymity Disallowed",             // [RFC5079]
	436: "Bad Identity-Info",                // [RFC4474]
	437: "Unsupported Certificate",          // [RFC4474]
	438: "Invalid Identity Header",          // [RFC4474]
	439: "First Hop Lacks Outbound Support", // [RFC5626]
	440: "Max-Breadth Exceeded",             // [RFC5393]
	470: "Consent Needed",                   // [RFC5360]
	480: "Temporarily Unavailable",          // fast busy or soft fail
	481: "Call/Transaction Does Not Exist",  // bad news
	482: "Loop Detected",                    // froot looping
	483: "Too Many Hops",                    // froot looping
	484: "Address Incomplete",
	485: "Ambiguous",
	486: "Busy Here",
	487: "Request Terminated",
	488: "Not Acceptable Here",
	489: "Bad Event", // [RFC3265]
	491: "Request Pending",
	493: "Undecipherable",
	494: "Security Agreement Required", // [RFC3329]

	// 5xx: Server Error -- the server failed to fulfill an apparently
	//      valid request;
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Time-out",
	505: "Version Not Supported",
	513: "Message Too Large",
	580: "Precondition Failure", // [RFC3312]

	// 6xx: Global Failure -- the request cannot be fulfilled at any
	//      server.
	600: "Busy Everywhere",
	603: "Decline",
	604: "Does Not Exist Anywhere",
	606: "Not Acceptable",
	687: "Dialog Terminated",
}
