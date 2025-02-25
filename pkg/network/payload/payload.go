package payload

import "github.com/ixje/neo-go-legacy/pkg/io"

// Payload is anything that can be binary encoded/decoded.
type Payload interface {
	io.Serializable
}

// NullPayload is a dummy payload with no fields.
type NullPayload struct {
}

// NewNullPayload returns zero-sized stub payload.
func NewNullPayload() *NullPayload {
	return &NullPayload{}
}

// DecodeBinary implements Serializable interface.
func (p *NullPayload) DecodeBinary(r *io.BinReader) {}

// EncodeBinary implements Serializable interface.
func (p *NullPayload) EncodeBinary(w *io.BinWriter) {}
