# gosip

Version: 0.1  
Copyright (c) 2010-2014 Justine Alexandra Roberts Tunney

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
taking this non-monolithic approach to design, you can actually build a highly
available (five nines) global telephony service with zero interruptions during
software upgrades.

gosip parses SIP messages using the Ragel finite state machine compiler. Much
of the BNF was copied directly from the SIP RFC. This implementation approach
offers a fair amount of confidence that the parser is not only fast, but also
correct and secure.

gosip is less suitable for PBX, IVR, and VoIP reselling. For these things, you
should consider Asterisk, FreeSWITCH, or SER.

## RFCs

- [SIP (RFC 3261)](https://tools.ietf.org/html/rfc3261)
- [SDP (RFC 4566)](https://tools.ietf.org/html/rfc4566)
- [RTP (RFC 3550)](https://tools.ietf.org/html/rfc3550)
- [RTP DTMF (RFC 4733)](https://tools.ietf.org/html/rfc4733)
- [SIP 100rel (RFC 3262)](https://tools.ietf.org/html/rfc3262)
- [SIP: Locating a Server (RFC 3263)](https://tools.ietf.org/html/rfc3263)
