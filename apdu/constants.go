package apdu

// Instruction represents a smart card instruction as defined in 
// ISO/IEC 7816-4:2005(E) 5.1.2
type Instruction uint
const (
	InsActivateFile Instruction = 0x44
	InsAppendRecord Instruction = 0xE2
	InsChangeReferenceData Instruction = 0x24
	InsDeactivateFile Instruction = 0xe0
	InsDeleteFile Instruction = 0x04
	INS_DISABLE_VERIFICATION_REQUIREMENT = 0x26
	INS_ENABLE_VERIFICATION_REQUIREMENT = 0x28
)