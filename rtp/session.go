// RTP Session Transport Layer.

package rtp

import (
	"github.com/jart/gosip/dsp"
	"github.com/jart/gosip/sdp"
	"log"
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
	C <-chan *Frame
	R chan<- *Frame
	E <-chan error

	// Underlying UDP socket.
	Sock *net.UDPConn

	// Address of remote endpoint. This might change mid-session. If it's nil,
	// then egress packets are dropped.
	Peer *net.UDPAddr

	// Header is the current header that gets mutated with each transmit.
	Header Header

	obuf []byte
}

// Creates a new RTP µLaw 20ptime session listening on host with a random port
// selected from the range [16384,32768].
func NewSession(host string) (rs *Session, err error) {
	conn, err := listenRTP(host)
	if err != nil {
		return nil, err
	}
	sock := conn.(*net.UDPConn)
	c := make(chan *Frame)
	r := make(chan *Frame)
	e := make(chan error)
	go receiver(sock, c, e, r)
	return &Session{
		C:    c,
		E:    e,
		R:    r,
		Sock: sock,
		Header: Header{
			PT:   sdp.ULAWCodec.PT,
			Seq:  666,
			TS:   0,
			Ssrc: rand.Uint32(),
		},
		obuf: make([]byte, HeaderSize+160),
	}, nil
}

func (rs *Session) Send(frame *Frame) (err error) {
	if rs == nil || rs.Sock == nil || rs.Peer == nil {
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

func (rs *Session) SendRaw(data []byte, samps uint32) (err error) {
	if rs == nil || rs.Sock == nil || rs.Peer == nil {
		return nil
	}
	rs.Header.Write(rs.obuf)
	rs.Header.TS += samps
	rs.Header.Seq++
	_, err = rs.Sock.WriteTo(append(rs.obuf[:HeaderSize], data...), rs.Peer)
	return
}

func (rs *Session) Close() {
	if rs == nil || rs.Sock == nil {
		return
	}
	rs.Sock.Close()
	if frame, ok := <-rs.C; ok {
		rs.R <- frame
	}
	<-rs.E
	rs.Sock = nil
	close(rs.R)
}

func (rs *Session) CloseAfterError() {
	if rs == nil || rs.Sock == nil {
		return
	}
	rs.Sock.Close()
	rs.Sock = nil
	close(rs.R)
	<-rs.E
}

func receiver(sock *net.UDPConn, c chan<- *Frame, e chan<- error, r <-chan *Frame) {
	buf := make([]byte, 2048)
	frame := new(Frame)
	for {
		amt, _, err := sock.ReadFrom(buf)
		if err != nil {
			close(c)
			e <- err
			break
		}
		// TODO(jart): Verify source address?
		// TODO(jart): Packet reordering? Drop duplicate packets?
		// TODO(jart): DTMF?
		var phdr Header
		err = phdr.Read(buf)
		if err != nil {
			// TODO(jart): Best logging strategy?
			continue
		}
		if phdr.PT != sdp.ULAWCodec.PT {
			continue
		}
		if amt != HeaderSize+160 {
			continue
		}
		for n := 0; n < 160; n++ {
			frame[n] = int16(dsp.UlawToLinear(int64(buf[HeaderSize+n])))
		}
		c <- frame
		frame = <-r
	}
	close(e)
}

func listenRTP(host string) (sock net.PacketConn, err error) {
	if strings.Contains(host, ":") {
		return net.ListenPacket("udp", host)
	}
	for i := 0; i < rtpBindMaxAttempts; i++ {
		port := rtpBindPortMin + rand.Int63()%(rtpBindPortMax-rtpBindPortMin+1)
		saddr := net.JoinHostPort(host, strconv.FormatInt(port, 10))
		sock, err = net.ListenPacket("udp", saddr)
		if err == nil || !strings.Contains(err.Error(), "address already in use") {
			break
		}
		log.Println("RTP listen congestion:", saddr)
	}
	return
}
