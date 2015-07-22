package main

// #cgo pkg-config: libpulse-simple
// #include <stdlib.h>
// #include <pulse/simple.h>
// #include <pulse/error.h>
import "C"

import (
	"errors"
	"flag"
	"github.com/jart/gosip/dsp"
	"github.com/jart/gosip/rtp"
	"github.com/jart/gosip/sdp"
	"github.com/jart/gosip/sip"
	"github.com/jart/gosip/util"
	"net"
	"os"
	"time"
	"unsafe"
)

const (
	hz       = 8000
	chans    = 1
	ptime    = 20
	ssize    = 2
	psamps   = hz / (1000 / ptime) * chans
	pbytes   = psamps * ssize
	filename = "/var/lib/asterisk/sounds/en/cc-yougotpranked.s16"
)

var (
	address      = flag.String("sipAddress", ":9020", "Listen address")
	paServerFlag = flag.String("paServer", "", "Pulse Audio server name")
	paSinkFlag   = flag.String("paSink", "", "Pulse Audio device or sink name")
	paName       = C.CString("fone")
)

func main() {
	pa, err := makePulseAudio(C.PA_STREAM_PLAYBACK, filename)
	if err != nil {
		panic(err)
	}
	defer C.pa_simple_free(pa)
	defer C.pa_simple_flush(pa, nil)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tick := time.NewTicker(ptime * time.Millisecond)
	defer tick.Stop()
	func() {
		for {
			var buf [pbytes]byte
			select {
			case <-tick.C:
				got, _ := f.Read(buf[:])
				if got < pbytes {
					return
				}
				var paerr C.int
				if C.pa_simple_write(pa, unsafe.Pointer(&buf[0]), pbytes, &paerr) != 0 {
					panic(C.GoString(C.pa_strerror(paerr)))
				}
			}
		}
	}()
	os.Exit(0)

	// Create RTP audio session.
	rs, err := rtp.NewSession("")
	if err != nil {
		panic(err)
	}
	defer rs.Close()
	rtpPort := uint16(rs.Sock.LocalAddr().(*net.UDPAddr).Port)

	invite := &sip.Msg{
		Method:  sip.MethodInvite,
		Request: &sip.URI{User: "echo", Host: "127.0.0.1", Port: 5060},
		Payload: &sdp.SDP{
			Origin: sdp.Origin{ID: util.GenerateOriginID()},
			Audio: &sdp.Media{
				Port:   rtpPort,
				Codecs: []sdp.Codec{sdp.ULAWCodec, sdp.DTMFCodec},
			},
		},
	}

	// Create a SIP phone call.
	dl, err := sip.NewDialog(invite)
	if err != nil {
		panic(err)
	}

	// We're going to send white noise every 20ms.
	var frame rtp.Frame
	awgn := dsp.NewAWGN(-45.0)
	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()

	// Hangup after 200ms.
	death := time.After(200 * time.Millisecond)

	// Let's GO!
	var answered bool
	for {
		select {
		case <-ticker.C:
			for n := 0; n < 160; n++ {
				frame[n] = awgn.Get()
			}
			if err := rs.Send(&frame); err != nil {
				panic("RTP send failed: " + err.Error())
			}
		case err := <-dl.OnErr:
			panic(err)
		case state := <-dl.OnState:
			switch state {
			case sip.DialogAnswered:
				answered = true
			case sip.DialogHangup:
				if !answered {
					panic("Call didn't get answered!")
				}
				return
			}
		case rs.Peer = <-dl.OnPeer:
		case frame := <-rs.C:
			rs.R <- frame
		case err := <-rs.E:
			panic("RTP recv failed: " + err.Error())
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
	ba.maxlength = pbytes * 4
	ba.tlength = pbytes
	ba.prebuf = pbytes * 2
	ba.minreq = pbytes
	ba.fragsize = 0xffffffff

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
