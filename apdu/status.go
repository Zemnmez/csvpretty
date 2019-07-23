package apdu

//go:generate go run golang.org/x/tools/cmd/stringer -type=Status

// Status represents a processing state as defined in ISO/IEC 7816-4:2005(E) 5.1.3
// true statuses are actually 1-2 bytes, but we're taking some liberties here
// for the sake of expression -- statuses use the two byte fields SW1 and SW2, and
// often leave SW2 blank for extra data
type Status uint16

// StatusBlankedBytesMask represents a mask showing when bits are blanked (XX in the spec)
// these bits are used to indicate additional information
const StatusBlankedBytesMask = 0x00FF

// "normal processing"
const (
	// No further qualification
	StatusNoFurtherQualification Status = 0x9000

	// SW2 encodes the number of data bytes still available
	StatusBytesAvailable Status = 0x61FF
)

// "warning processing"
const (
	// Warning: state of non-volatile memory is unchanged (further qualification in SW2)
	StatusWarningNonVolatileMemoryUnchanged Status = 0x62FF

	// Warning: state of non-volatile memory has changed (further qualifications in SW2)
	StatusWarningNonVolatileMemoryChanged Status = 0x63FF
)

// "execution error"
const (
	// Error: state of non-volatile memory is unchanged (further qualification in SW2)
	StatusErrorNonVolatileMemoryUnchanged Status = 0x64FF

	// Error: state of non-volatile memory has changed (further qualifications in SW2)
	StatusErrorNonVolatileMemoryChanged Status = 0x65FF

	// Error: "security related issues" (really)
	StatusErrorSecurityIssue Status = 0x66FF
)

// "checking error"
const (
	// Error: "wrong length; no further indication"
	StatusErrorWrongLength Status = 0x6700

	// Error: "Functions in CLA not supported (further qualification in SW2)"
	StatusErrorFunctionUnsupported Status = 0x68FF

	// Error: "Command not allowed (further qualification in SW2)"
	StatusErrorCommandNotAllowed Status = 0x69FF

	// Error: "Wrong parameters P1-P2 (further qualification in SW2)"
	StatusErrorWrongParametersWithInfo Status = 0x6aFF

	// Error: "Wrong parameters P1-P2"
	StatusErrorWrongParameters Status = 0x6b00

	// Error: "Wrong Le field; SW2 encodes the exact number of available data bytes"
	StatusErrorWrongLe Status = 0x6CFF

	// Error: "Instruction code not supported or invalid"
	StatusErrorUnsupportedInstructionCode Status = 0x6d00

	// Error: "Class not supported"
	StatusErrorUnsupportedClass Status = 0x6e00

	//Error: "No precise diagnosis"
	StatusErrorUnknown Status = 0x6f00
)
