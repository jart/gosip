package sip

import (
	"github.com/jart/gosip/sdp"
	"github.com/jart/gosip/util"
	"log"
)

func NewRequest(tp Transport, method string, to, from *Addr) *Msg {
	return &Msg{
		Method:     method,
		Request:    to.Uri.Copy(),
		Via:        tp.Via().Branch(),
		From:       from.Or(tp.Contact()).Tag(),
		To:         to.Copy(),
		Contact:    tp.Contact(),
		CallID:     util.GenerateCallID(),
		CSeq:       util.GenerateCSeq(),
		CSeqMethod: method,
		Headers:    Headers{"User-Agent": "gosip/1.o"},
	}
}

func NewResponse(msg *Msg, status int) *Msg {
	return &Msg{
		IsResponse: true,
		Status:     status,
		Phrase:     Phrases[status],
		Via:        msg.Via,
		From:       msg.From,
		To:         msg.To,
		CallID:     msg.CallID,
		CSeq:       msg.CSeq,
		CSeqMethod: msg.CSeqMethod,
		Headers:    Headers{"User-Agent": "gosip/1.o"},
	}
}

// http://tools.ietf.org/html/rfc3261#section-17.1.1.3
func NewAck(invite *Msg) *Msg {
	return &Msg{
		Method:     "ACK",
		Request:    invite.Request,
		Via:        invite.Via,
		From:       invite.From,
		To:         invite.To,
		CallID:     invite.CallID,
		CSeq:       invite.CSeq,
		CSeqMethod: "ACK",
		Route:      invite.Route,
		Headers:    Headers{"User-Agent": "gosip/1.o"},
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
		Route:      invite.Route,
		Headers:    Headers{"User-Agent": "gosip/1.o"},
	}
}

func NewBye(last, invite, ok200 *Msg) *Msg {
	return &Msg{
		Request:    ok200.Contact.Uri,
		Via:        invite.Via.Branch(),
		From:       last.From,
		To:         last.To,
		CallID:     last.CallID,
		Method:     "BYE",
		CSeq:       last.CSeq + 1,
		CSeqMethod: "BYE",
		Route:      ok200.RecordRoute,
		Headers:    make(map[string]string),
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
