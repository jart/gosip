# gosip

Version: 0.1  

## About

gosip (pronounced "gossip") is a library that lets you make phone calls using
the Go programming language. It provides a full-stack SIP/RTP implementation
that's tailored towards making calls over the PSTN through services such as
Flowroute.

gosip is most suitable for backend telephony applications. This is especially
true for apps that do interesting things with audio. gosip supports DSP out of
the box by providing SSE optimised audio mixing, an assembly implementation of
the µLaw codec (no other codecs are supported), and a comfort noise generator.

Telephony applications have traditionally been written on top of PBX systems
like Asterisk via an extension interface. Frameworks such as Adhearsion have
introduced further layers of abstraction to this unwieldy regime. These systems
are slow, difficult to administer, and in many cases superfluous to the needs
of the telephony app developer. gosip sets you free from Asterisk because you
can have everything in a single easy-to-deploy binary.

gosip has excellent support for SRV/NAPTR failover by way of timeouts, ICMP
refusal, and 502 Service Unavailable responses. It also supports SIP redirects
as well as changing the audio/signalling path mid-call. But most importantly,
gosip is lightweight enough that you can actually design your app to use a
single process for each phone call (assuming your app is in the audio path.) By
taking this non-monolithic approach to design, you can actually build a higher
availability global telephony service with zero interruptions during
software upgrades.

gosip parses SIP messages using the Ragel finite state machine compiler. Much
of the BNF was copied directly from the SIP RFC. This implementation approach
offers a fair amount of confidence that the parser is not only fast, but also
correct and secure.

gosip is less suitable for PBX, IVR, and VoIP reselling. For these things, you
should consider Asterisk, FreeSWITCH, or SER.

## RFCs

- [RFC3261: SIP](https://tools.ietf.org/html/rfc3261)
- [RFC4566: SDP](https://tools.ietf.org/html/rfc4566)
- [RFC3550: RTP](https://tools.ietf.org/html/rfc3550)
- [RFC4733: RTP DTMF](https://tools.ietf.org/html/rfc4733)
- [RFC3262: SIP 100rel](https://tools.ietf.org/html/rfc3262)
- [RFC3263: SIP: Locating a Server](https://tools.ietf.org/html/rfc3263)

## Example RFCs

- [RFC4475: SIP Torture Test Messages](https://tools.ietf.org/html/rfc4475)
- [RFC3665: SIP basic call flows](https://tools.ietf.org/html/rfc3665)
- [RFC3666: SIP PSTN call flows](https://tools.ietf.org/html/rfc3666)
- [RFC6216: Example call flows using SIP security mechanisms](https://tools.ietf.org/html/rfc6216)
