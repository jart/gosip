// RTP Session Transport Layer.

package rtp

import (
	"errors"
	"fmt"
	"github.com/jart/gosip/dsp"
	"github.com/jart/gosip/sdp"
	"github.com/jart/gosip/util"
	"math/rand"
	"net"
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

	// Underlying UDP socket.
	Sock *net.UDPConn

	// Address of remote endpoint. This might change mid-session. If it's nil,
	// then egress packets are dropped.
	Peer *net.UDPAddr

	// Header is the current header that gets mutated with each transmit.
	Header Header

	ibuf []byte
	obuf []byte
	phdr Header
}

// Creates a new RTP µLaw 20ptime session listening on host with a random port
// selected from the range [16384,32768].
func NewSession(host string) (s *Session, err error) {
	sock, err := listenRTP(host)
	if err != nil {
		return nil, err
	}
	return &Session{
		Sock: sock.(*net.UDPConn),
		Header: Header{
			PT:   sdp.ULAWCodec.PT,
			Seq:  666,
			TS:   0,
			Ssrc: rand.Uint32(),
		},
	}, err
}

func (s *Session) Send(frame Frame) (err error) {
	if s.Peer == nil {
		return nil
	}
	if s.obuf == nil {
		s.obuf = make([]byte, HeaderSize+160)
	}
	s.Header.Write(s.obuf)
	s.Header.TS += 160
	s.Header.Seq++
	for n := 0; n < 160; n++ {
		s.obuf[HeaderSize+n] = byte(dsp.LinearToUlaw(int64(frame[n])))
	}
	_, err = s.Sock.WriteTo(s.obuf, s.Peer)
	return
}

func (s *Session) Recv(frame Frame) (err error) {
	if s.ibuf == nil {
		s.ibuf = make([]byte, 2048)
	}
	for {
		amt, _, err := s.Sock.ReadFrom(s.ibuf)
		if err != nil {
			return err
		}
		// TODO(jart): Verify source address?
		// TODO(jart): Packet reordering? Drop duplicate packets?
		// TODO(jart): DTMF?
		var phdr Header
		err = phdr.Read(s.ibuf)
		if err != nil {
			return err
		}
		if phdr.PT != sdp.ULAWCodec.PT {
			continue
		}
		if amt != HeaderSize+160 {
			return errors.New(fmt.Sprintf("Unexpected RTP packet size: %d", amt))
		}
		for n := 0; n < 160; n++ {
			frame[n] = int16(dsp.UlawToLinear(int64(s.ibuf[HeaderSize+n])))
		}
		return nil
	}
}

func listenRTP(host string) (sock net.PacketConn, err error) {
	for i := 0; i < rtpBindMaxAttempts; i++ {
		port := rtpBindPortMin + rand.Int()%(rtpBindPortMax-rtpBindPortMin+1)
		saddr := util.HostPortToString(host, uint16(port))
		sock, err = net.ListenPacket("udp", saddr)
		if err != nil {
			if !strings.Contains(err.Error(), "address already in use") {
				break
			}
		}
	}
	return
}
