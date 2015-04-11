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
			b.Write(escapeParam([]byte(k)))
			v := params[k]
			if v != "" {
				b.WriteByte('=')
				b.Write(escapeParam([]byte(v)))
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
			b.Write(tokencify([]byte(k)))
			v := params[k]
			if v != "" {
				b.WriteByte('=')
				b.Write(quote([]byte(v)))
			}
		}
	}
}

func (params Params) Has(k string) bool {
	_, ok := params["lr"]
	return ok
}
