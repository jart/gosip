package sip

import (
	"github.com/jart/gosip/sdp"
	"github.com/jart/gosip/util"
	"log"
)

const (
	GosipUserAgent = "gosip/1.o"
	GosipAllow     = "INVITE, ACK, CANCEL, BYE, OPTIONS"
)

func NewRequest(tp *Transport, method string, to, from *Addr) *Msg {
	return &Msg{
		Method:     method,
		Request:    to.Uri.Copy(),
		Via:        tp.Via.Copy().Branch(),
		From:       from.Or(tp.Contact).Tag(),
		To:         to.Copy(),
		Contact:    tp.Contact,
		CallID:     util.GenerateCallID(),
		CSeq:       util.GenerateCSeq(),
		CSeqMethod: method,
		Headers:    DefaultHeaders(),
	}
}

func NewResponse(msg *Msg, status int) *Msg {
	return &Msg{
		IsResponse:  true,
		Status:      status,
		Phrase:      Phrase(status),
		Via:         msg.Via,
		From:        msg.From,
		To:          msg.To,
		CallID:      msg.CallID,
		CSeq:        msg.CSeq,
		CSeqMethod:  msg.CSeqMethod,
		RecordRoute: msg.RecordRoute,
		Headers:     DefaultHeaders(),
	}
}

// http://tools.ietf.org/html/rfc3261#section-17.1.1.3
func NewAck(original, last *Msg) *Msg {
	return &Msg{
		Method:     "ACK",
		Request:    original.Request,
		Via:        original.Via.Copy().SetNext(nil),
		From:       original.From,
		To:         last.To,
		CallID:     original.CallID,
		CSeq:       original.CSeq,
		CSeqMethod: "ACK",
		Route:      last.RecordRoute.Reversed(),
		Headers:    DefaultHeaders(),
	}
}

func NewCancel(invite *Msg) *Msg {
	if invite.IsResponse || invite.Method != "INVITE" {
		log.Printf("Can't CANCEL anything non-INVITE:\n%s", invite)
	}
	return &Msg{
		Method:     "CANCEL",
		Request:    invite.Request,
		Via:        invite.Via,
		From:       invite.From,
		To:         invite.To,
		CallID:     invite.CallID,
		CSeq:       invite.CSeq,
		CSeqMethod: "CANCEL",
		Route:      invite.RecordRoute.Reversed(),
		Headers:    DefaultHeaders(),
	}
}

func NewBye(invite, last *Msg) *Msg {
	return &Msg{
		Method:     "BYE",
		Request:    last.Contact.Uri,
		Via:        invite.Via,
		From:       last.From,
		To:         last.To,
		CallID:     last.CallID,
		CSeq:       last.CSeq + 1,
		CSeqMethod: "BYE",
		Route:      last.RecordRoute.Reversed(),
		Headers:    DefaultHeaders(),
	}
}

// Returns true if `resp` can be considered an appropriate response to `msg`.
// Do not use for ACKs.
func ResponseMatch(msg, resp *Msg) bool {
	return (resp.IsResponse &&
		resp.CSeq == msg.CSeq &&
		resp.CSeqMethod == msg.Method &&
		resp.Via.Last().Compare(msg.Via))
}

// Returns true if `ack` can be considered an appropriate response to `msg`.
// we don't enforce a matching Via because some VoIP software will generate a
// new branch for ACKs.
func AckMatch(msg, ack *Msg) bool {
	return (!ack.IsResponse &&
		ack.Method == "ACK" &&
		ack.CSeq == msg.CSeq &&
		ack.CSeqMethod == "ACK" &&
		ack.Via.Last().CompareAddr(msg.Via))
}

func AttachSDP(msg *Msg, sdp *sdp.SDP) {
	msg.Headers["Content-Type"] = "application/sdp"
	msg.Payload = sdp.String()
}

func DefaultHeaders() Headers {
	return Headers{
		"User-Agent": GosipUserAgent,
		"Allow":      GosipAllow,
	}
}
