
//+build none

package apdu

//go:generate go run golang.org/x/tools/cmd/stringer -type=Status,StatusCategory
//go:generate go run golang.org/x/tools/cmd/stringer -type=Status,StatusCategory -output=status_errorstring.go -linecomment
//go:generate gofmt -w -r String->Error ./status_errorstring.go
//go:generate go run github.com/zemnmez/cardauth/apdu/gen/repl -from=_Status -to=_Error_Status -file=status_errorstring.go

// StatusCategory represents SW1 bytes for broad groups of statuses, some of which
// may not have a corresponding Status.
type StatusCategory uint8

// SW1 bytes for broad groups
const (
	StatusBytesAvailable                    StatusCategory = 0x61 // SW2 encodes the number of data bytes still available
	StatusWarningNonVolatileMemoryUnchanged StatusCategory = 0x62 // Warning: state of non-volatile memory is unchanged (further qualification in SW2)
	StatusWarningNonVolatileMemoryChanged   StatusCategory = 0x63 // Warning: state of non-volatile memory has changed (further qualifications in SW2)
	StatusErrorNonVolatileMemoryUnchanged   StatusCategory = 0x64 // Error: state of non-volatile memory is unchanged (further qualification in SW2)
	StatusErrorNonVolatileMemoryChanged     StatusCategory = 0x65 // Error: state of non-volatile memory has changed (further qualifications in SW2)
	StatusErrorSecurityIssue                StatusCategory = 0x66 // Error: "security related issues"
	StatusErrorFunctionUnsupported          StatusCategory = 0x68 // Error: "Functions in CLA not supported (further qualification in SW2)"
	StatusErrorCommandNotAllowed            StatusCategory = 0x69 // Error: "Command not allowed (further qualification in SW2)"
	StatusErrorWrongParametersWithInfo      StatusCategory = 0x6a // Error: "Wrong parameters P1-P2 (further qualification in SW2)"
	StatusErrorWrongLe                      StatusCategory = 0x6C // Error: "Wrong Le field: SW2 encodes the exact number of available data bytes"
)

// Status represents a processing state as defined in ISO/IEC 7816-4:2005(E) 5.1.3
// true statuses are actually 1-2 bytes, but we're taking some liberties here
// for the sake of expression -- statuses use the two byte fields SW1 and SW2, and
// often leave SW2 blank for extra data
type Status uint16

// statuses
const (
	StatusNoFurtherQualification                                              Status = 0x9000 // No further qualification
	StatusWarningNonVolatileMemoryUnchangedNoInformation                      Status = 0x6202 // Warning: state of non-volatile memory has not changed: No reason given
	StatusWarningNonVolatileMemoryUnchangePartOfReturnedDataCorrupted         Status = 0x6281 // Warning: state of non-volatile memory has not changed: Part of returned data may be corrupted
	StatusWarningNonVolatileMemoryUnchangeEndOfFileOrRecordReachedPrematurely Status = 0x6282 // Warning: state of non-volatile memory has not changed: End of file or record reached before reading Ne bytes
	StatusWarningNonVolatileMemoryUnchangeSelectedFileDeactivated             Status = 0x6283 // Warning: state of non-volatile memory has not changed: Selected file deactivated
	StatusWarningNonVolatileMemoryUnchangeFileControlInvalid                  Status = 0x6284 // Warning: state of non-volatile memory has not changed: File control information not formatted correctly
	StatusWarningNonVolatileMemoryUnchangeFileTerminated                      Status = 0x6285 // Warning: state of non-volatile memory has changed: Selected file in termination state
	StatusWarningNonVolatileMemoryUnchangeNoInputDataFromSensor               Status = 0x6286 // Warning: state of non-volatile memory has not changed: No input data available from a sensor on the card
	StatusWarningNonVolatileMemoryChangedNoInformation                        Status = 0x6300 // Warning: state of non-volatile memory has changed: No reason given
	StatusWarningNonVolatileMemoryChangedFileFull                             Status = 0x6381 // Warning: state of non-volatile memory has changed: File filled up by the last write
	StatusErrorWrongLength                                                    Status = 0x6700 // Error: "wrong length; no further indication"
	StatusErrorWrongParameters                                                Status = 0x6b00 // Error: "Wrong parameters P1-P2"
	StatusErrorUnsupportedInstructionCode                                     Status = 0x6d00 // Error: "Instruction code not supported or invalid"
	StatusErrorUnsupportedClass                                               Status = 0x6e00 // Error: "Class not supported"
	StatusErrorUnknown                                                        Status = 0x6f00 // Error: "No precise diagnosis"

	// Counter from 0 to 15 encoded by 'X' (exact meaning depending on the command)
	// ^ bruh what the fuck how do you expect me to even write code for that
)
