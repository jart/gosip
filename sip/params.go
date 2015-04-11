package sip

import (
	"bytes"
	"sort"
)

type Params map[string]string

func (params Params) Copy() Params {
	res := make(Params, len(params))
	for k, v := range params {
		res[k] = v
	}
	return res
}

func (params Params) Append(b *bytes.Buffer) {
	if params != nil && len(params) > 0 {
		keys := make([]string, len(params))
		i := 0
		for k, _ := range params {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			b.WriteByte(';')
			appendEscaped(b, []byte(k), paramc)
			v := params[k]
			if v != "" {
				b.WriteByte('=')
				appendEscaped(b, []byte(v), paramc)
			}
		}
	}
}

func (params Params) AppendQuoted(b *bytes.Buffer) {
	if params != nil && len(params) > 0 {
		keys := make([]string, len(params))
		i := 0
		for k, _ := range params {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			b.WriteByte(';')
			appendSanitized(b, []byte(k), tokenc)
			v := params[k]
			if v != "" {
				b.WriteByte('=')
				appendQuoted(b, []byte(v))
			}
		}
	}
}

func (params Params) Has(k string) bool {
	_, ok := params["lr"]
	return ok
}
