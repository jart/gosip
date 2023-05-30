// Copyright 2020 Justine Alexandra Roberts Tunney
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// SIP Protocol Status Definitions
//
// http://www.iana.org/assignments/sip-parameters

package sip

const (
	// 1xx: Provisional -- request received, continuing to process the request.
	StatusTrying               = 100 // Indicates server is not totally pwnd.
	StatusRinging              = 180 // Remote phone is definitely ringing.
	StatusCallIsBeingForwarded = 181
	StatusQueued               = 182
	StatusSessionProgress      = 183 // Establish early media (PSTN ringback)

	// 2xx: Success -- the action was successfully received, understood,
	//      and accepted;
	StatusOK             = 200 // Call is answered
	StatusAccepted       = 202 // [RFC3265]
	StatusNoNotification = 204 // [RFC5839]

	// 3xx: Redirection -- further action needs to be taken in order to
	//      complete the request;
	StatusMultipleChoices    = 300
	StatusMovedPermanently   = 301
	StatusMovedTemporarily   = 302 // Send your call there instead kthx.
	StatusUseProxy           = 305 // You fool! Send your call there instead.
	StatusAlternativeService = 380

	// 4xx: Client Error -- the request contains bad syntax or cannot be
	//      fulfilled at this server;
	StatusBadRequest                   = 400 // Missing headers, bad format, etc.
	StatusUnauthorized                 = 401 // Resend request with auth header.
	StatusPaymentRequired              = 402 // I am greedy.
	StatusForbidden                    = 403 // gtfo
	StatusNotFound                     = 404 // wat?
	StatusMethodNotAllowed             = 405 // I don't support that type of request.
	StatusNotAcceptable                = 406
	StatusProxyAuthenticationRequired  = 407
	StatusRequestTimeout               = 408
	StatusConflict                     = 409
	StatusGone                         = 410 // Shaniqua don't live here no more.
	StatusLengthRequired               = 411
	StatusConditionalRequestFailed     = 412 // [RFC3903]
	StatusRequestEntityTooLarge        = 413
	StatusRequestURITooLong            = 414
	StatusUnsupportedMediaType         = 415
	StatusUnsupportedURIScheme         = 416
	StatusUnknownResourcePriority      = 417
	StatusBadExtension                 = 420
	StatusExtensionRequired            = 421
	StatusSessionIntervalTooSmall      = 422 // [RFC4028]
	StatusIntervalTooBrief             = 423
	StatusUseIdentityHeader            = 428 // [RFC4474]
	StatusProvideReferrerIdentity      = 429 // [RFC3892]
	StatusFlowFailed                   = 430 // [RFC5626]
	StatusAnonymityDisallowed          = 433 // [RFC5079]
	StatusBadIdentityInfo              = 436 // [RFC4474]
	StatusUnsupportedCertificate       = 437 // [RFC4474]
	StatusInvalidIdentityHeader        = 438 // [RFC4474]
	StatusFirstHopLacksOutboundSupport = 439 // [RFC5626]
	StatusMaxBreadthExceeded           = 440 // [RFC5393]
	StatusConsentNeeded                = 470 // [RFC5360]
	StatusTemporarilyUnavailable       = 480 // fast busy or soft fail
	StatusCallTransactionDoesNotExist  = 481 // Bad news
	StatusLoopDetected                 = 482 // Froot looping
	StatusTooManyHops                  = 483 // Froot looping
	StatusAddressIncomplete            = 484
	StatusAmbiguous                    = 485
	StatusBusyHere                     = 486
	StatusRequestTerminated            = 487
	StatusNotAcceptableHere            = 488
	StatusBadEvent                     = 489 // [RFC3265]
	StatusRequestPending               = 491
	StatusUndecipherable               = 493
	StatusSecurityAgreementRequired    = 494 // [RFC3329]

	// 5xx: Server Error -- the server failed to fulfill an apparently
	//      valid request;
	StatusInternalServerError = 500
	StatusNotImplemented      = 501
	StatusBadGateway          = 502
	StatusServiceUnavailable  = 503
	StatusGatewayTimeout      = 504
	StatusVersionNotSupported = 505
	StatusMessageTooLarge     = 513
	StatusPreconditionFailure = 580 // [RFC3312]

	// 6xx: Global Failure -- the request cannot be fulfilled at any
	//      server.
	StatusBusyEverywhere       = 600
	StatusDecline              = 603
	StatusDoesNotExistAnywhere = 604
	StatusNotAcceptable606     = 606
	StatusDialogTerminated     = 687
)

func Phrase(status int) string {
	if phrase, ok := phrases[status]; ok {
		return phrase
	}
	return "Unknown Status Code"
}

var phrases = map[int]string{
	StatusTrying:                       "Trying",
	StatusRinging:                      "Ringing",
	StatusCallIsBeingForwarded:         "Call Is Being Forwarded",
	StatusQueued:                       "Queued",
	StatusSessionProgress:              "Session Progress",
	StatusOK:                           "OK",
	StatusAccepted:                     "Accepted",
	StatusNoNotification:               "No Notification",
	StatusMultipleChoices:              "Multiple Choices",
	StatusMovedPermanently:             "Moved Permanently",
	StatusMovedTemporarily:             "Moved Temporarily",
	StatusUseProxy:                     "Use Proxy",
	StatusAlternativeService:           "Alternative Service",
	StatusBadRequest:                   "Bad Request",
	StatusUnauthorized:                 "Unauthorized",
	StatusPaymentRequired:              "Payment Required",
	StatusForbidden:                    "Forbidden",
	StatusNotFound:                     "Not Found",
	StatusMethodNotAllowed:             "Method Not Allowed",
	StatusNotAcceptable:                "Not Acceptable",
	StatusProxyAuthenticationRequired:  "Proxy Authentication Required",
	StatusRequestTimeout:               "Request Timeout",
	StatusConflict:                     "Conflict",
	StatusGone:                         "Gone",
	StatusLengthRequired:               "Length Required",
	StatusConditionalRequestFailed:     "Conditional Request Failed",
	StatusRequestEntityTooLarge:        "Request Entity Too Large",
	StatusRequestURITooLong:            "Request-URI Too Long",
	StatusUnsupportedMediaType:         "Unsupported Media Type",
	StatusUnsupportedURIScheme:         "Unsupported URI Scheme",
	StatusUnknownResourcePriority:      "Unknown Resource-Priority",
	StatusBadExtension:                 "Bad Extension",
	StatusExtensionRequired:            "Extension Required",
	StatusSessionIntervalTooSmall:      "Session Interval Too Small",
	StatusIntervalTooBrief:             "Interval Too Brief",
	StatusUseIdentityHeader:            "Use Identity Header",
	StatusProvideReferrerIdentity:      "Provide Referrer Identity",
	StatusFlowFailed:                   "Flow Failed",
	StatusAnonymityDisallowed:          "Anonymity Disallowed",
	StatusBadIdentityInfo:              "Bad Identity-Info",
	StatusUnsupportedCertificate:       "Unsupported Certificate",
	StatusInvalidIdentityHeader:        "Invalid Identity Header",
	StatusFirstHopLacksOutboundSupport: "First Hop Lacks Outbound Support",
	StatusMaxBreadthExceeded:           "Max-Breadth Exceeded",
	StatusConsentNeeded:                "Consent Needed",
	StatusTemporarilyUnavailable:       "Temporarily Unavailable",
	StatusCallTransactionDoesNotExist:  "Call/Transaction Does Not Exist",
	StatusLoopDetected:                 "Loop Detected",
	StatusTooManyHops:                  "Too Many Hops",
	StatusAddressIncomplete:            "Address StatusIncomplete",
	StatusAmbiguous:                    "Ambiguous",
	StatusBusyHere:                     "Busy Here",
	StatusRequestTerminated:            "Request Terminated",
	StatusNotAcceptableHere:            "Not Acceptable Here",
	StatusBadEvent:                     "Bad Event",
	StatusRequestPending:               "Request Pending",
	StatusUndecipherable:               "Undecipherable",
	StatusSecurityAgreementRequired:    "Security Agreement Required",
	StatusInternalServerError:          "Internal Server Error",
	StatusNotImplemented:               "Not Implemented",
	StatusBadGateway:                   "Bad Gateway",
	StatusServiceUnavailable:           "Service Unavailable",
	StatusGatewayTimeout:               "Gateway Time-out",
	StatusVersionNotSupported:          "Version Not Supported",
	StatusMessageTooLarge:              "Message Too Large",
	StatusPreconditionFailure:          "Precondition Failure",
	StatusBusyEverywhere:               "Busy Everywhere",
	StatusDecline:                      "Decline",
	StatusDoesNotExistAnywhere:         "Does Not Exist Anywhere",
	StatusNotAcceptable606:             "Not Acceptable",
	StatusDialogTerminated:             "Dialog Terminated",
}
