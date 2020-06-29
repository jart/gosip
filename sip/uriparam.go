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

// URI Parameter Library

package sip

import (
	"bytes"
	"strings"
)

// URIParam is a linked list of ;key=values for URI parameters.
type URIParam struct {
	Name  string
	Value string
	Next  *URIParam
}

// Get returns an entry in O(n) time.
func (p *URIParam) Get(name string) *URIParam {
	if p == nil {
		return nil
	}
	if strings.EqualFold(p.Name, name) {
		return p
	}
	return p.Next.Get(name)
}

// Append serializes URI parameters in insertion order.
func (p *URIParam) Append(b *bytes.Buffer) {
	if p == nil {
		return
	}
	p.Next.Append(b)
	b.WriteByte(';')
	appendEscaped(b, []byte(p.Name), paramc)
	if p.Value != "" {
		b.WriteByte('=')
		appendEscaped(b, []byte(p.Value), paramc)
	}
}
