# gosip

Version: 0.1  
Copyright: Copyright (c) 2010-2014 Justine Tunney  
License: MIT

## About

gosip (pronounced like the word "gossip") is a VoIP telephony library
written in Google's Go programming language.

gosip provides a barebones sip/sdp/rtp implementation suitable for use
over a trusted network.  To talk to the outside world you should
deploy your gosip applications behind a session border controller
(like tube, FreeSWITCH, or OpenSER) which are more capable of dealing
with security, network quality issues and SIP interop.

I was originally going to write bindings for sofia-sip but ultimately
decided it'd be quicker, less buggy, and faster performing to write a
lightweight SIP stack from scratch.

## Installation

Once Go is installed, just run ``make`` to build/test/install.

## Learning

The following unit tests also serve as tutorials to help you
understand SIP and the abstractions provided by this library.

- sip/rawsip_test.go: How to do SIP the hard way

- sip/url_test.go: Shows you what the SIP URL data structures look like

- sip/addr_test.go: Addresses are pretty much URLs inside angle brackets

- sip/msg_test.go: How the data structure for SIP packets works

- sip/manualsip_test.go: Make SIP easier with parser/formatter objects

- sip/echo_test.go: Manually make a test call to an echo application

## Overview

This is what a sip stack looks like:

   +-----------------------------------------------------------------------+
   |  9. Application Layer (your code)                                     |
   +-----------------------------------------------------------------------+
   |  8. Telephony API (tel/...)                                           |
   +---------------------------------------+-------------------------------+
   |  6. SIP Transaction (sip/transact.go) | 6. Media Codecs (sip/dsp.go)  |
   +---------------------------------------+-------------------------------+
   |  5. SIP Transport (sip/transport.go)  | 5. RTP Transport (sip/rtp.go) |
   +---------------------------------------+-------------------------------+
   |  2/3/4. Network Transport Layer                                       |
   +-----------------------------------------------------------------------+
   |  1. Tubes                                                             |
   +-----------------------------------------------------------------------+
   |  0. Electrons and Photons                                             |
   +-----------------------------------------------------------------------+

## RFCs

- [SIP (RFC 3261)](https://tools.ietf.org/html/rfc3261)
- [SDP (RFC 4566)](https://tools.ietf.org/html/rfc4566)
- [RTP (RFC 3550)](https://tools.ietf.org/html/rfc3550)
- [RTP DTMF (RFC 4733)](https://tools.ietf.org/html/rfc4733)
- [SIP 100rel (RFC 3262)](https://tools.ietf.org/html/rfc3262)
