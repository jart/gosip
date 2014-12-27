// RTP Session Transport Layer.

package rtp

import (
	"errors"
	"fmt"
	"github.com/jart/gosip/dsp"
	"github.com/jart/gosip/sdp"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

const (
	rtpBindMaxAttempts = 10
	rtpBindPortMin     = 16384
	rtpBindPortMax     = 32768
)

type Frame [160]int16

// Session allows sending and receiving slinear frames for a single SIP media
// session. These frames are encoded as µLaw and transmitted over UDP. No
// support for RTCP is provided.
type Session struct {

	// Channel to which received RTP frames and errors are published.
	C chan *Frame
	E chan error

	// Underlying UDP socket.
	Sock *net.UDPConn

	// Address of remote endpoint. This might change mid-session. If it's nil,
	// then egress packets are dropped.
	Peer *net.UDPAddr

	// Header is the current header that gets mutated with each transmit.
	Header Header

	ibuf []byte
	obuf []byte
}

// Creates a new RTP µLaw 20ptime session listening on host with a random port
// selected from the range [16384,32768].
func NewSession(host string) (rs *Session, err error) {
	sock, err := listenRTP(host)
	if err != nil {
		return nil, err
	}
	rs = &Session{
		C:    make(chan *Frame, 32),
		E:    make(chan error, 1),
		Sock: sock.(*net.UDPConn),
		Header: Header{
			PT:   sdp.ULAWCodec.PT,
			Seq:  666,
			TS:   0,
			Ssrc: rand.Uint32(),
		},
		obuf: make([]byte, HeaderSize+160),
		ibuf: make([]byte, 2048),
	}
	rs.launchConsumer()
	return
}

func (rs *Session) Send(frame *Frame) (err error) {
	if rs.Peer == nil {
		return nil
	}
	rs.Header.Write(rs.obuf)
	rs.Header.TS += 160
	rs.Header.Seq++
	for n := 0; n < 160; n++ {
		rs.obuf[HeaderSize+n] = byte(dsp.LinearToUlaw(int64(frame[n])))
	}
	_, err = rs.Sock.WriteTo(rs.obuf, rs.Peer)
	return
}

func (rs *Session) launchConsumer() {
	go func() {
		for {
			frame, err := rs.recv()
			if err != nil {
				rs.E <- err
				return
			}
			rs.C <- frame
		}
	}()
}

func (rs *Session) recv() (frame *Frame, err error) {
	frame = new(Frame)
	for {
		amt, _, err := rs.Sock.ReadFrom(rs.ibuf)
		if err != nil {
			return nil, err
		}
		// TODO(jart): Verify source address?
		// TODO(jart): Packet reordering? Drop duplicate packets?
		// TODO(jart): DTMF?
		var phdr Header
		err = phdr.Read(rs.ibuf)
		if err != nil {
			return nil, err
		}
		if phdr.PT != sdp.ULAWCodec.PT {
			continue
		}
		if amt != HeaderSize+160 {
			return nil, errors.New(fmt.Sprintf("Unexpected RTP packet size: %d", amt))
		}
		for n := 0; n < 160; n++ {
			frame[n] = int16(dsp.UlawToLinear(int64(rs.ibuf[HeaderSize+n])))
		}
		return frame, nil
	}
}

func listenRTP(host string) (sock net.PacketConn, err error) {
	for i := 0; i < rtpBindMaxAttempts; i++ {
		port := rtpBindPortMin + rand.Int63()%(rtpBindPortMax-rtpBindPortMin+1)
		saddr := net.JoinHostPort(host, strconv.FormatInt(port, 10))
		sock, err = net.ListenPacket("udp", saddr)
		if err != nil {
			if !strings.Contains(err.Error(), "address already in use") {
				break
			}
		}
	}
	return
}
