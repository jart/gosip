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

package dialog

import (
	"log"

	"github.com/cresta/gosip/sip"
	"github.com/cresta/gosip/util"
)

const (
	GosipUA    = "slytherin/1.o"
	GosipAllow = "INVITE, ACK, CANCEL, BYE, OPTIONS"
)

func NewRequest(tp *Transport, method string, to, from *sip.Addr) *sip.Msg {
	return &sip.Msg{
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

func NewResponse(msg *sip.Msg, status int) *sip.Msg {
	return &sip.Msg{
		Status:      status,
		Phrase:      sip.Phrase(status),
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
func NewAck(msg, invite *sip.Msg) *sip.Msg {
	return &sip.Msg{
		Method:             sip.MethodAck,
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

func NewCancel(invite *sip.Msg) *sip.Msg {
	if invite.IsResponse() || invite.Method != sip.MethodInvite {
		log.Printf("Can't CANCEL anything non-INVITE:\r\n%s", invite)
	}
	return &sip.Msg{
		Method:     sip.MethodCancel,
		Request:    invite.Request,
		Via:        invite.Via,
		From:       invite.From,
		To:         invite.To,
		CallID:     invite.CallID,
		CSeq:       invite.CSeq,
		CSeqMethod: sip.MethodCancel,
		Route:      invite.Route,
	}
}

func NewBye(invite, remote *sip.Msg, lseq *int) *sip.Msg {
	if lseq == nil {
		lseq = new(int)
		*lseq = invite.CSeq
	}
	*lseq++
	return &sip.Msg{
		Method:     sip.MethodBye,
		Request:    remote.Contact.Uri,
		From:       invite.From,
		To:         remote.To,
		CallID:     invite.CallID,
		CSeq:       *lseq,
		CSeqMethod: sip.MethodBye,
		Route:      remote.RecordRoute.Reversed(),
	}
}

// Returns true if `resp` can be considered an appropriate response to `msg`.
// Do not use for ACKs.
func ResponseMatch(req, rsp *sip.Msg) bool {
	return (rsp.IsResponse() &&
		rsp.CSeq == req.CSeq &&
		rsp.CSeqMethod == req.Method &&
		rsp.Via.Last().CompareHostPort(req.Via) &&
		rsp.Via.Last().CompareBranch(req.Via))
}

// Returns true if `ack` can be considered an appropriate response to `msg`.
// We don't enforce a matching Via because some VoIP software will generate a
// new branch for ACKs.
func AckMatch(msg, ack *sip.Msg) bool {
	return (!ack.IsResponse() &&
		ack.Method == sip.MethodAck &&
		ack.CSeq == msg.CSeq &&
		ack.CSeqMethod == sip.MethodAck &&
		ack.Via.Last().CompareHostPort(msg.Via))
}
