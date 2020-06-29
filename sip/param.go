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

// Addr / Via Parameter Library

package sip

import (
	"bytes"
	"strings"
)

// Param is a linked list of ;key="values" for Addr/Via parameters.
type Param struct {
	Name  string
	Value string
	Next  *Param
}

// Get returns an entry in O(n) time.
func (p *Param) Get(name string) *Param {
	if p == nil {
		return nil
	}
	if strings.EqualFold(p.Name, name) {
		return p
	}
	return p.Next.Get(name)
}

// Append serializes parameters in insertion order.
func (p *Param) Append(b *bytes.Buffer) {
	if p == nil {
		return
	}
	p.Next.Append(b)
	b.WriteByte(';')
	appendSanitized(b, []byte(p.Name), tokenc)
	if p.Value != "" {
		b.WriteByte('=')
		appendQuoted(b, []byte(p.Value))
	}
}
