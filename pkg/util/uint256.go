package util

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ixje/neo-go-legacy/pkg/io"
)

// Uint256Size is the size of Uint256 in bytes.
const Uint256Size = 32

// Uint256 is a 32 byte long unsigned integer.
type Uint256 [Uint256Size]uint8

// Uint256DecodeStringLE attempts to decode the given string (in LE representation) into an Uint256.
func Uint256DecodeStringLE(s string) (u Uint256, err error) {
	if len(s) != Uint256Size*2 {
		return u, fmt.Errorf("expected string size of %d got %d", Uint256Size*2, len(s))
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return u, err
	}
	return Uint256DecodeBytesLE(b)
}

// Uint256DecodeStringBE attempts to decode the given string (in BE representation)
// into an Uint256.
func Uint256DecodeStringBE(s string) (u Uint256, err error) {
	if len(s) != Uint256Size*2 {
		return u, fmt.Errorf("expected string size of %d got %d", Uint256Size*2, len(s))
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return u, err
	}

	return Uint256DecodeBytesBE(b)
}

// Uint256DecodeBytesBE attempts to decode the given string (in BE representation) into an Uint256.
func Uint256DecodeBytesBE(b []byte) (u Uint256, err error) {
	if len(b) != Uint256Size {
		return u, fmt.Errorf("expected []byte of size %d got %d", Uint256Size, len(b))
	}
	copy(u[:], b)
	return u, nil
}

// Uint256DecodeBytesLE attempts to decode the given string (in LE representation) into an Uint256.
func Uint256DecodeBytesLE(b []byte) (u Uint256, err error) {
	b = ArrayReverse(b)
	return Uint256DecodeBytesBE(b)
}

// BytesBE returns a byte slice representation of u.
func (u Uint256) BytesBE() []byte {
	return u[:]
}

// Reverse reverses the Uint256 object
func (u Uint256) Reverse() Uint256 {
	res, _ := Uint256DecodeBytesLE(u.BytesBE())
	return res
}

// BytesLE return a little-endian byte representation of u.
func (u Uint256) BytesLE() []byte {
	return ArrayReverse(u.BytesBE())
}

// Equals returns true if both Uint256 values are the same.
func (u Uint256) Equals(other Uint256) bool {
	return u == other
}

// String implements the stringer interface.
func (u Uint256) String() string {
	return u.StringBE()
}

// StringBE produces string representation of Uint256 with BE byte order.
func (u Uint256) StringBE() string {
	return hex.EncodeToString(u.BytesBE())
}

// StringLE produces string representation of Uint256 with LE byte order.
func (u Uint256) StringLE() string {
	return hex.EncodeToString(u.BytesLE())
}

// UnmarshalJSON implements the json unmarshaller interface.
func (u *Uint256) UnmarshalJSON(data []byte) (err error) {
	var js string
	if err = json.Unmarshal(data, &js); err != nil {
		return err
	}
	js = strings.TrimPrefix(js, "0x")
	*u, err = Uint256DecodeStringLE(js)
	return err
}

// MarshalJSON implements the json marshaller interface.
func (u Uint256) MarshalJSON() ([]byte, error) {
	return []byte(`"0x` + u.StringLE() + `"`), nil
}

// CompareTo compares two Uint256 with each other. Possible output: 1, -1, 0
//
//	1 implies u > other.
//
// -1 implies u < other.
//
//	0 implies  u = other.
func (u Uint256) CompareTo(other Uint256) int { return bytes.Compare(u[:], other[:]) }

// EncodeBinary implements io.Serializable interface.
func (u *Uint256) EncodeBinary(w *io.BinWriter) {
	w.WriteBytes(u[:])
}

// DecodeBinary implements io.Serializable interface.
func (u *Uint256) DecodeBinary(r *io.BinReader) {
	r.ReadBytes(u[:])
}
