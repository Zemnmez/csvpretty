// +build none

package apdu

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"io"
)

var _ = []interface {
	encoding.BinaryMarshaler
	io.WriterTo
}{
	&RawApdu{},
	&Header{},
	&ApduCommand{},
}

/



// MarshalBinary returns this raw APDU as binary bytes
func (r RawApdu) MarshalBinary() (binary []byte, err error) {
	binary = []byte{
		r.CLA, r.INS, r.P1, r.P2,
	}

	binary = append(binary, r.Lc...)
	binary = append(binary, r.Command...)
	binary = append(binary, r.Le...)

	return
}

var endianness = binary.BigEndian

// Apdu represents the Application Protocol Data Unit structure
// of the smart card protocol.
//
// see: https://en.wikipedia.org/wiki/Smart_card_application_protocol_data_unit
type Apdu struct {
	Class       byte
	Instruction byte
	P1          byte
	P2          byte
	Data        []byte
	Expected    byte
}

var _ interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
} = Apdu{}

// MarshalBinary returns the binary representation of this APDU message
func (a Apdu) MarshalBinary() (data []byte, err error) {
	req := []byte{
		a.Class,
		a.Instruction,
		a.P1,
		a.P2,
	}
}
