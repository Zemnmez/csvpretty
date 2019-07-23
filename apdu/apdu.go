package apdu

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"io"
)

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