package apdu

import (
	"encoding/binary"
	"fmt"
)

var endianness = binary.BigEndian

// Apdu represents the Application Protocol Data Unit structure
// of the smart card protocol.
//
// see: https://en.wikipedia.org/wiki/Smart_card_application_protocol_data_unit
type Apdu struct {
	Class       byte
	Instruction Instruction
	P1          byte
	P2          byte
	Data        []byte
	Expected    byte
}

type Instruction byte

func (i Instruction) String() string {
	if v, ok := instrStringMap[i]; ok {
		return v
	}

	return fmt.Sprintf("Instruction(%d)", int(i))
}

func (i Instrution) Info() string {
	if v, ok := instrInfoMap[i]; ok {
		return v
	}

	return fmt.Sprintf("no info!", int(i))
}

func (i Instruction) IsAlias() (isAlias bool) {
	_, isAlias = instrSecondariesMap[i]
	return
}

func (i Instruction) Resolve() Instruction {
	if v, ok := instrSecondariesMap[i]; ok {
		return v
	}
	return i
}
