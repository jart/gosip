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

package main

// #cgo pkg-config: ncurses libpulse-simple
// #include <stdlib.h>
// #include <ncurses.h>
// #include <pulse/simple.h>
// #include <pulse/error.h>
import "C"

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
	"unsafe"

	"github.com/cresta/gosip/dialog"
	"github.com/cresta/gosip/dsp"
	"github.com/cresta/gosip/rtp"
	"github.com/cresta/gosip/sdp"
	"github.com/cresta/gosip/sip"
	"github.com/cresta/gosip/util"
)

const (
	hz     = 8000
	chans  = 1
	ptime  = 20
	ssize  = 2
	psamps = hz / (1000 / ptime) * chans
	pbytes = psamps * ssize
)

var (
	addressFlag  = flag.String("address", "", "Public IP (or hostname) of the local machine. Defaults to asking an untrusted webserver.")
	paServerFlag = flag.String("paServer", "", "PulseAudio server name")
	paSinkFlag   = flag.String("paSink", "", "PulseAudio device or sink name")
	muteFlag     = flag.Bool("mute", false, "Send comfort noise rather than microphone input")
	paName       = C.CString("fone")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s URI\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	// Whom Are We Calling?
	requestURIString := flag.Args()[0]
	requestURI, err := sip.ParseURI([]byte(requestURIString))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad Request URI: %s\n", err.Error())
		os.Exit(1)
	}

	// Computer Speaker
	speaker, err := makePulseAudio(C.PA_STREAM_PLAYBACK, requestURIString)
	if err != nil {
		panic(err)
	}
	defer C.pa_simple_free(speaker)
	defer C.pa_simple_flush(speaker, nil)

	// Computer Microphone
	mic, err := makePulseAudio(C.PA_STREAM_RECORD, requestURIString)
	if err != nil {
		panic(err)
	}
	defer C.pa_simple_free(mic)

	// Get Public IP Address
	publicIP := *addressFlag
	if publicIP == "" {
		publicIP, err = getPublicIP()
		if err != nil {
			panic(err)
		}
	}

	// Create RTP Session
	rs, err := rtp.NewSession("")
	if err != nil {
		panic(err)
	}
	defer rs.Close()
	rtpPort := uint16(rs.Sock.LocalAddr().(*net.UDPAddr).Port)

	// Construct SIP INVITE
	invite := &sip.Msg{
		Method:  sip.MethodInvite,
		Request: requestURI,
		Via:     &sip.Via{Host: publicIP},
		To:      &sip.Addr{Uri: requestURI},
		From:    &sip.Addr{Uri: &sip.URI{Host: publicIP, User: os.Getenv("USER")}},
		Contact: &sip.Addr{Uri: &sip.URI{Host: publicIP}},
		Payload: &sdp.SDP{
			Addr: publicIP,
			Origin: sdp.Origin{
				ID:   util.GenerateOriginID(),
				Addr: publicIP,
			},
			Audio: &sdp.Media{
				Port:   rtpPort,
				Codecs: []sdp.Codec{sdp.ULAWCodec, sdp.DTMFCodec},
			},
		},
		Warning: "dark lord funk you up",
	}

	// Create SIP Dialog State Machine
	dl, err := dialog.NewDialog(invite)
	if err != nil {
		panic(err)
	}

	// Send Audio Every 20ms
	var frame rtp.Frame
	awgn := dsp.NewAWGN(-45.0)
	ticker := time.NewTicker(ptime * time.Millisecond)
	defer ticker.Stop()

	// Ctrl+C or Kill Graceful Shutdown
	death := make(chan os.Signal, 1)
	signal.Notify(death, os.Interrupt, os.Kill)

	// DTMF Terminal Input
	keyboard := make(chan byte)
	keyboardStart := func() {
		C.cbreak()
		C.noecho()
		go func() {
			var buf [1]byte
			for {
				amt, err := os.Stdin.Read(buf[:])
				if err != nil || amt != 1 {
					log.Printf("Keyboard: %s\r\n", err)
					return
				}
				keyboard <- buf[0]
			}
		}()
	}

	C.initscr()
	defer C.endwin()

	// Let's GO!
	var answered bool
	var paerr C.int
	for {
		select {

		// Send Audio
		case <-ticker.C:
			if *muteFlag {
				for n := 0; n < psamps; n++ {
					frame[n] = awgn.Get()
				}
			} else {
				if C.pa_simple_read(mic, unsafe.Pointer(&frame[0]), pbytes, &paerr) != 0 {
					log.Printf("Microphone: %s\r\n", C.GoString(C.pa_strerror(paerr)))
					break
				}
			}
			if err := rs.Send(&frame); err != nil {
				log.Printf("RTP: %s\r\n", err.Error())
			}

		// Send DTMF
		case ch := <-keyboard:
			if err := rs.SendDTMF(ch); err != nil {
				log.Printf("DTMF: %s\r\n", err.Error())
				break
			}
			log.Printf("DTMF: %c\r\n", ch)

		// Receive Audio
		case frame := <-rs.C:
			if len(frame) != psamps {
				log.Printf("RTP: Received undersized frame: %d != %d\r\n", len(frame), psamps)
			} else {
				if C.pa_simple_write(speaker, unsafe.Pointer(&frame[0]), pbytes, &paerr) != 0 {
					log.Printf("Speaker: %s\r\n", C.GoString(C.pa_strerror(paerr)))
				}
			}
			rs.R <- frame

		// Signalling
		case rs.Peer = <-dl.OnPeer:
		case state := <-dl.OnState:
			switch state {
			case dialog.Answered:
				answered = true
				keyboardStart()
			case dialog.Hangup:
				if answered {
					return
				} else {
					os.Exit(1)
				}
			}

		// Errors and Interruptions
		case err := <-dl.OnErr:
			log.Fatalf("SIP: %s\r\n", err.Error())
		case err := <-rs.E:
			log.Printf("RTP: %s\r\n", err.Error())
			rs.CloseAfterError()
			dl.Hangup <- true
		case <-death:
			dl.Hangup <- true
		}
	}
}

func makePulseAudio(direction C.pa_stream_direction_t, streamName string) (*C.pa_simple, error) {
	var ss C.pa_sample_spec
	ss.format = C.PA_SAMPLE_S16NE
	ss.rate = hz
	ss.channels = chans

	var ba C.pa_buffer_attr
	if direction == C.PA_STREAM_PLAYBACK {
		ba.maxlength = pbytes * 4
		ba.tlength = pbytes
		ba.prebuf = pbytes * 2
		ba.minreq = pbytes
		ba.fragsize = 0xffffffff
	} else {
		ba.maxlength = pbytes * 4
		ba.tlength = 0xffffffff
		ba.prebuf = 0xffffffff
		ba.minreq = 0xffffffff
		ba.fragsize = pbytes
	}

	var paServer *C.char
	if *paServerFlag != "" {
		paServer = C.CString(*paServerFlag)
		defer C.free(unsafe.Pointer(paServer))
	}

	var paSink *C.char
	if *paSinkFlag != "" {
		paSink = C.CString(*paSinkFlag)
		defer C.free(unsafe.Pointer(paSink))
	}

	paStreamName := C.CString(streamName)
	defer C.free(unsafe.Pointer(paStreamName))

	var paerr C.int
	pa := C.pa_simple_new(paServer, paName, direction, paSink, paStreamName, &ss, nil, &ba, &paerr)
	if pa == nil {
		return nil, errors.New(C.GoString(C.pa_strerror(paerr)))
	}
	return pa, nil
}

func getPublicIP() (string, error) {
	resp, err := http.Get("http://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
