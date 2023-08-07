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

// SIP Address Library
//
// For example:
//
//   "Justine Tunney" <sip:jart@example.test;isup-oli=29>;tag=feedabee
//
// Roughly equates to:
//
//   {Display: "Justine Tunney",
//    Params: {"tag": "feedabee"},
//    Uri: {Scheme: "sip",
//          User: "jart",
//          Pass: "",
//          Host: "example.test",
//          Port: "",
//          Params: {"isup-oli": "29"}}}
//

package sip

import (
	"bytes"

	"github.com/cresta/gosip/util"
)

// Represents a SIP Address Linked List
type Addr struct {
	Uri     *URI   // never nil
	Display string // blank if not specified
	Param   *Param // these look like ;key=lol;rport;key=wut
	Next    *Addr  // for comma separated lists of addresses
}

func (addr *Addr) String() string {
	var b bytes.Buffer
	addr.Append(&b)
	return b.String()
}

// Returns self if non-nil, otherwise other.
func (addr *Addr) Or(other *Addr) *Addr {
	if addr == nil {
		return other
	}
	return addr
}

// Sets newly generated tag ID and returns self.
func (addr *Addr) Tag() *Addr {
	addr.Param = &Param{"tag", util.GenerateTag(), addr.Param}
	return addr
}

// Reassembles a SIP address into a buffer.
func (addr *Addr) Append(b *bytes.Buffer) {
	if addr == nil {
		return
	}
	if addr.Display != "" {
		appendQuoted(b, []byte(addr.Display))
		b.WriteByte(' ')
	}
	b.WriteByte('<')
	addr.Uri.Append(b)
	b.WriteByte('>')
	addr.Param.Append(b)
	if addr.Next != nil {
		b.WriteByte(',')
		b.WriteByte(' ')
		addr.Next.Append(b)
	}
}

// Deep copies a new Addr object.
func (addr *Addr) Copy() *Addr {
	if addr == nil {
		return nil
	}
	res := new(Addr)
	res.Uri = addr.Uri.Copy()
	res.Param = addr.Param
	res.Next = addr.Next.Copy()
	return res
}

func (addr *Addr) CompareHostPort(other *Addr) bool {
	if addr != nil && other != nil {
		if addr.Uri.CompareHostPort(other.Uri) {
			return true
		}
	}
	return false
}

// Returns pointer to last addr in linked list.
func (addr *Addr) Last() *Addr {
	if addr != nil {
		for ; addr.Next != nil; addr = addr.Next {
		}
	}
	return addr
}

// Returns number of items in the linked list.
func (addr *Addr) Len() int {
	count := 0
	for ; addr != nil; addr = addr.Next {
		count++
	}
	return count
}

// Returns self with linked list reversed.
func (addr *Addr) Reversed() *Addr {
	var res *Addr
	for ; addr != nil; addr = addr.Next {
		a := new(Addr)
		*a = *addr
		a.Next = res
		res = a
	}
	return res
}
