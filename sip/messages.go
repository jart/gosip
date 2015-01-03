package sip

import (
	"github.com/jart/gosip/sdp"
	"github.com/jart/gosip/util"
	"log"
)

const (
	GosipUA    = "gosip/0.1"
	GosipAllow = "INVITE, ACK, CANCEL, BYE, OPTIONS"
)

func NewRequest(tp *Transport, method string, to, from *Addr) *Msg {
	return &Msg{
		Method:     method,
		Request:    to.Uri.Copy(),
		Via:        tp.Via.Copy().Branch(),
		From:       from.Or(tp.Contact).Copy().Tag(),
		To:         to.Copy(),
		CallID:     util.GenerateCallID(),
		CSeq:       util.GenerateCSeq(),
		CSeqMethod: method,
		UserAgent:  GosipUA,
	}
}

func NewResponse(msg *Msg, status int) *Msg {
	return &Msg{
		Status:      status,
		Phrase:      Phrase(status),
		Via:         msg.Via,
		From:        msg.From,
		To:          msg.To,
		CallID:      msg.CallID,
		CSeq:        msg.CSeq,
		CSeqMethod:  msg.CSeqMethod,
		RecordRoute: msg.RecordRoute,
		UserAgent:   GosipUA,
		Allow:       GosipAllow,
	}
}

// http://tools.ietf.org/html/rfc3261#section-17.1.1.3
func NewAck(msg, invite *Msg) *Msg {
	return &Msg{
		Method:             MethodAck,
		Request:            msg.Contact.Uri,
		From:               msg.From,
		To:                 msg.To,
		Via:                msg.Via.Detach(),
		CallID:             msg.CallID,
		CSeq:               msg.CSeq,
		CSeqMethod:         "ACK",
		Route:              msg.RecordRoute.Reversed(),
		Authorization:      invite.Authorization,
		ProxyAuthorization: invite.ProxyAuthorization,
		UserAgent:          GosipUA,
	}
}

func NewCancel(invite *Msg) *Msg {
	if invite.IsResponse() || invite.Method != MethodInvite {
		log.Printf("Can't CANCEL anything non-INVITE:\n%s", invite)
	}
	return &Msg{
		Method:     MethodCancel,
		Request:    invite.Request,
		Via:        invite.Via,
		From:       invite.From,
		To:         invite.To,
		CallID:     invite.CallID,
		CSeq:       invite.CSeq,
		CSeqMethod: MethodCancel,
		Route:      invite.Route,
	}
}

func NewBye(invite, remote *Msg, lseq *int) *Msg {
	if lseq == nil {
		lseq = new(int)
		*lseq = invite.CSeq
	}
	*lseq++
	return &Msg{
		Method:     MethodBye,
		Request:    remote.Contact.Uri,
		From:       invite.From,
		To:         remote.To,
		CallID:     invite.CallID,
		CSeq:       *lseq,
		CSeqMethod: MethodBye,
		Route:      remote.RecordRoute.Reversed(),
	}
}

// Returns true if `resp` can be considered an appropriate response to `msg`.
// Do not use for ACKs.
func ResponseMatch(req, rsp *Msg) bool {
	return (rsp.IsResponse() &&
		rsp.CSeq == req.CSeq &&
		rsp.CSeqMethod == req.Method &&
		rsp.Via.Last().CompareHostPort(req.Via) &&
		rsp.Via.Last().CompareBranch(req.Via))
}

// Returns true if `ack` can be considered an appropriate response to `msg`.
// We don't enforce a matching Via because some VoIP software will generate a
// new branch for ACKs.
func AckMatch(msg, ack *Msg) bool {
	return (!ack.IsResponse() &&
		ack.Method == MethodAck &&
		ack.CSeq == msg.CSeq &&
		ack.CSeqMethod == MethodAck &&
		ack.Via.Last().CompareHostPort(msg.Via))
}

func AttachSDP(msg *Msg, ms *sdp.SDP) {
	if msg.Headers == nil {
		msg.Headers = Headers{}
	}
	msg.ContentType = sdp.ContentType
	msg.Payload = ms.String()
}
