package sip

import (
	"strings"
	"sync"

	"github.com/jart/gosip/sdp"
)

// BodyParser is the type of functions parsing message bodies.
type BodyParser func([]byte) (Payload, error)

// BodyParserRegistry is a registry of body parsing functions.
type BodyParserRegistry struct {
	registry sync.Map
}

// Register registers a boty parser for a given content type.
func (r *BodyParserRegistry) Register(ctype string, parser BodyParser) {
	r.registry.Store(strings.ToLower(ctype), parser)
}

func (r *BodyParserRegistry) get(ctype string) (BodyParser, bool) {
	v, ok := r.registry.Load(strings.ToLower(ctype))
	if !ok {
		return nil, false
	}
	vv, ok := v.(BodyParser)
	if !ok {
		// this cannot happen
		return nil, false
	}
	return vv, true
}

// DefaultBodyParserRegistry is the default BodyParserRegistry used by
// ParseMsg.
var DefaultBodyParserRegistry BodyParserRegistry

// RegisterBodyParser registers a body parser on DefaultBodyParserRegistry.
func RegisterBodyParser(ctype string, parser BodyParser) {
	DefaultBodyParserRegistry.Register(ctype, parser)
}

func init() {
	RegisterBodyParser(
		sdp.ContentType,
		func(body []byte) (Payload, error) {
			return sdp.Parse(string(body))
		},
	)
}
