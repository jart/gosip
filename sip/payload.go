package sip

type Payload interface {
	ContentType() string
	Data() []byte
}

type MiscPayload struct {
	T string
	D []byte
}

func (p *MiscPayload) ContentType() string {
	return p.T
}

func (p *MiscPayload) Data() []byte {
	return p.D
}
