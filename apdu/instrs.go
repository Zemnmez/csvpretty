
package apdu

// Code generated DO NOT EDIT.

//go:generate go run github.com/zemnmez/cardauth/apdu/gen/instrs 
//go:generate gofmt -w -s $GOFILE

var(
// RawInstrDeactivateFile represents the "DEACTIVATE FILE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
RawInstrDeactivateFile = [...]byte{0x04}

// RawInstrEraseRecord(s) represents the "ERASE RECORD (S)" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.8
RawInstrEraseRecord(s) = [...]byte{0x0C}

// RawInstrEraseBinary represents the "ERASE BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.7
RawInstrEraseBinary = [...]byte{0x0E, 0x0F}

// RawInstrPerformScqlOperation represents the "PERFORM SCQL OPERATION" instruction as defined in ISO/IEC 7816-4:2005(E) Part 7
RawInstrPerformScqlOperation = [...]byte{0x10}

// RawInstrPerformTransactionOperation represents the "PERFORM TRANSACTION OPERATION" instruction as defined in ISO/IEC 7816-4:2005(E) Part 7
RawInstrPerformTransactionOperation = [...]byte{0x12}

// RawInstrPerformUserOperation represents the "PERFORM USER OPERATION" instruction as defined in ISO/IEC 7816-4:2005(E) Part 7
RawInstrPerformUserOperation = [...]byte{0x14}

// RawInstrVerify represents the "VERIFY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.6
RawInstrVerify = [...]byte{0x20, 0x21}

// RawInstrManageSecurityEnvironment represents the "MANAGE SECURITY ENVIRONMENT" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.11
RawInstrManageSecurityEnvironment = [...]byte{0x22}

// RawInstrChangeReferenceData represents the "CHANGE REFERENCE DATA" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.7
RawInstrChangeReferenceData = [...]byte{0x24}

// RawInstrDisableVerificationRequirement represents the "DISABLE VERIFICATION REQUIREMENT" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.9
RawInstrDisableVerificationRequirement = [...]byte{0x26}

// RawInstrEnableVerificationRequirement represents the "ENABLE VERIFICATION REQUIREMENT" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.8
RawInstrEnableVerificationRequirement = [...]byte{0x28}

// RawInstrPerformSecurityOperation represents the "PERFORM SECURITY OPERATION" instruction as defined in ISO/IEC 7816-4:2005(E) Part 8
RawInstrPerformSecurityOperation = [...]byte{0x2A}

// RawInstrResetRetryCounter represents the "RESET RETRY COUNTER" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.10
RawInstrResetRetryCounter = [...]byte{0x2C}

// RawInstrActivateFile represents the "ACTIVATE FILE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
RawInstrActivateFile = [...]byte{0x44}

// RawInstrGenerateAsymmetricKeyPair represents the "GENERATE ASYMMETRIC KEY PAIR" instruction as defined in ISO/IEC 7816-4:2005(E) Part 8
RawInstrGenerateAsymmetricKeyPair = [...]byte{0x46}

// RawInstrManageChannel represents the "MANAGE CHANNEL" instruction as defined in ISO/IEC 7816-4:2005(E) 7.1.2
RawInstrManageChannel = [...]byte{0x70}

// RawInstrExternal(/Mutual)Authenticate represents the "EXTERNAL (/ MUTUAL) AUTHENTICATE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.4
RawInstrExternal(/Mutual)Authenticate = [...]byte{0x82}

// RawInstrGetChallenge represents the "GET CHALLENGE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.3
RawInstrGetChallenge = [...]byte{0x84}

// RawInstrGeneralAuthenticate represents the "GENERAL AUTHENTICATE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.5
RawInstrGeneralAuthenticate = [...]byte{0x86, 0x87}

// RawInstrInternalAuthenticate represents the "INTERNAL AUTHENTICATE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.2
RawInstrInternalAuthenticate = [...]byte{0x88}

// RawInstrSearchBinary represents the "SEARCH BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.6
RawInstrSearchBinary = [...]byte{0xA0, 0xA1}

// RawInstrSearchRecord represents the "SEARCH RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.7
RawInstrSearchRecord = [...]byte{0xA2}

// RawInstrSelect represents the "SELECT" instruction as defined in ISO/IEC 7816-4:2005(E) 7.1.1
RawInstrSelect = [...]byte{0xA4}

// RawInstrReadBinary represents the "READ BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.3
RawInstrReadBinary = [...]byte{0xB0, 0xB1}

// RawInstrReadRecord represents the "READ RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) (S) 7.3.3
RawInstrReadRecord = [...]byte{0xB2, 0xB3}

// RawInstrGetResponse represents the "GET RESPONSE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.6.1
RawInstrGetResponse = [...]byte{0xC0}

// RawInstrEnvelope represents the "ENVELOPE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.6.2
RawInstrEnvelope = [...]byte{0xC2, 0xC3}

// RawInstrGetData represents the "GET DATA" instruction as defined in ISO/IEC 7816-4:2005(E) 7.4.2
RawInstrGetData = [...]byte{0xCA, 0xCB}

// RawInstrWriteBinary represents the "WRITE BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.6
RawInstrWriteBinary = [...]byte{0xD0, 0xD1}

// RawInstrWriteRecord represents the "WRITE RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.4
RawInstrWriteRecord = [...]byte{0xD2}

// RawInstrUpdateBinary represents the "UPDATE BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.5
RawInstrUpdateBinary = [...]byte{0xD6, 0xD7}

// RawInstrPutData represents the "PUT DATA" instruction as defined in ISO/IEC 7816-4:2005(E) 7.4.3
RawInstrPutData = [...]byte{0xDA, 0xDB}

// RawInstrUpdateRecord represents the "UPDATE RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.5
RawInstrUpdateRecord = [...]byte{0xDC, 0xDD}

// RawInstrCreateFile represents the "CREATE FILE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
RawInstrCreateFile = [...]byte{0xE0}

// RawInstrAppendRecord represents the "APPEND RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.6
RawInstrAppendRecord = [...]byte{0xE2}

// RawInstrDeleteFile represents the "DELETE FILE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
RawInstrDeleteFile = [...]byte{0xE4}

// RawInstrTerminateDf represents the "TERMINATE DF" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
RawInstrTerminateDf = [...]byte{0xE6}

// RawInstrTerminateEf represents the "TERMINATE EF" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
RawInstrTerminateEf = [...]byte{0xE8}

// RawInstrTerminateCardUsage represents the "TERMINATE CARD USAGE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
RawInstrTerminateCardUsage = [...]byte{0xFE}

)

type Instruction []byte
var (
// InstrDeactivateFile represents the "DEACTIVATE FILE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
InstrDeactivateFile = Instruction(RawInstrDeactivateFile[:])

// InstrEraseRecord(s) represents the "ERASE RECORD (S)" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.8
InstrEraseRecord(s) = Instruction(RawInstrEraseRecord(s)[:])

// InstrEraseBinary represents the "ERASE BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.7
InstrEraseBinary = Instruction(RawInstrEraseBinary[:])

// InstrPerformScqlOperation represents the "PERFORM SCQL OPERATION" instruction as defined in ISO/IEC 7816-4:2005(E) Part 7
InstrPerformScqlOperation = Instruction(RawInstrPerformScqlOperation[:])

// InstrPerformTransactionOperation represents the "PERFORM TRANSACTION OPERATION" instruction as defined in ISO/IEC 7816-4:2005(E) Part 7
InstrPerformTransactionOperation = Instruction(RawInstrPerformTransactionOperation[:])

// InstrPerformUserOperation represents the "PERFORM USER OPERATION" instruction as defined in ISO/IEC 7816-4:2005(E) Part 7
InstrPerformUserOperation = Instruction(RawInstrPerformUserOperation[:])

// InstrVerify represents the "VERIFY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.6
InstrVerify = Instruction(RawInstrVerify[:])

// InstrManageSecurityEnvironment represents the "MANAGE SECURITY ENVIRONMENT" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.11
InstrManageSecurityEnvironment = Instruction(RawInstrManageSecurityEnvironment[:])

// InstrChangeReferenceData represents the "CHANGE REFERENCE DATA" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.7
InstrChangeReferenceData = Instruction(RawInstrChangeReferenceData[:])

// InstrDisableVerificationRequirement represents the "DISABLE VERIFICATION REQUIREMENT" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.9
InstrDisableVerificationRequirement = Instruction(RawInstrDisableVerificationRequirement[:])

// InstrEnableVerificationRequirement represents the "ENABLE VERIFICATION REQUIREMENT" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.8
InstrEnableVerificationRequirement = Instruction(RawInstrEnableVerificationRequirement[:])

// InstrPerformSecurityOperation represents the "PERFORM SECURITY OPERATION" instruction as defined in ISO/IEC 7816-4:2005(E) Part 8
InstrPerformSecurityOperation = Instruction(RawInstrPerformSecurityOperation[:])

// InstrResetRetryCounter represents the "RESET RETRY COUNTER" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.10
InstrResetRetryCounter = Instruction(RawInstrResetRetryCounter[:])

// InstrActivateFile represents the "ACTIVATE FILE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
InstrActivateFile = Instruction(RawInstrActivateFile[:])

// InstrGenerateAsymmetricKeyPair represents the "GENERATE ASYMMETRIC KEY PAIR" instruction as defined in ISO/IEC 7816-4:2005(E) Part 8
InstrGenerateAsymmetricKeyPair = Instruction(RawInstrGenerateAsymmetricKeyPair[:])

// InstrManageChannel represents the "MANAGE CHANNEL" instruction as defined in ISO/IEC 7816-4:2005(E) 7.1.2
InstrManageChannel = Instruction(RawInstrManageChannel[:])

// InstrExternal(/Mutual)Authenticate represents the "EXTERNAL (/ MUTUAL) AUTHENTICATE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.4
InstrExternal(/Mutual)Authenticate = Instruction(RawInstrExternal(/Mutual)Authenticate[:])

// InstrGetChallenge represents the "GET CHALLENGE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.3
InstrGetChallenge = Instruction(RawInstrGetChallenge[:])

// InstrGeneralAuthenticate represents the "GENERAL AUTHENTICATE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.5
InstrGeneralAuthenticate = Instruction(RawInstrGeneralAuthenticate[:])

// InstrInternalAuthenticate represents the "INTERNAL AUTHENTICATE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.5.2
InstrInternalAuthenticate = Instruction(RawInstrInternalAuthenticate[:])

// InstrSearchBinary represents the "SEARCH BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.6
InstrSearchBinary = Instruction(RawInstrSearchBinary[:])

// InstrSearchRecord represents the "SEARCH RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.7
InstrSearchRecord = Instruction(RawInstrSearchRecord[:])

// InstrSelect represents the "SELECT" instruction as defined in ISO/IEC 7816-4:2005(E) 7.1.1
InstrSelect = Instruction(RawInstrSelect[:])

// InstrReadBinary represents the "READ BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.3
InstrReadBinary = Instruction(RawInstrReadBinary[:])

// InstrReadRecord represents the "READ RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) (S) 7.3.3
InstrReadRecord = Instruction(RawInstrReadRecord[:])

// InstrGetResponse represents the "GET RESPONSE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.6.1
InstrGetResponse = Instruction(RawInstrGetResponse[:])

// InstrEnvelope represents the "ENVELOPE" instruction as defined in ISO/IEC 7816-4:2005(E) 7.6.2
InstrEnvelope = Instruction(RawInstrEnvelope[:])

// InstrGetData represents the "GET DATA" instruction as defined in ISO/IEC 7816-4:2005(E) 7.4.2
InstrGetData = Instruction(RawInstrGetData[:])

// InstrWriteBinary represents the "WRITE BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.6
InstrWriteBinary = Instruction(RawInstrWriteBinary[:])

// InstrWriteRecord represents the "WRITE RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.4
InstrWriteRecord = Instruction(RawInstrWriteRecord[:])

// InstrUpdateBinary represents the "UPDATE BINARY" instruction as defined in ISO/IEC 7816-4:2005(E) 7.2.5
InstrUpdateBinary = Instruction(RawInstrUpdateBinary[:])

// InstrPutData represents the "PUT DATA" instruction as defined in ISO/IEC 7816-4:2005(E) 7.4.3
InstrPutData = Instruction(RawInstrPutData[:])

// InstrUpdateRecord represents the "UPDATE RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.5
InstrUpdateRecord = Instruction(RawInstrUpdateRecord[:])

// InstrCreateFile represents the "CREATE FILE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
InstrCreateFile = Instruction(RawInstrCreateFile[:])

// InstrAppendRecord represents the "APPEND RECORD" instruction as defined in ISO/IEC 7816-4:2005(E) 7.3.6
InstrAppendRecord = Instruction(RawInstrAppendRecord[:])

// InstrDeleteFile represents the "DELETE FILE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
InstrDeleteFile = Instruction(RawInstrDeleteFile[:])

// InstrTerminateDf represents the "TERMINATE DF" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
InstrTerminateDf = Instruction(RawInstrTerminateDf[:])

// InstrTerminateEf represents the "TERMINATE EF" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
InstrTerminateEf = Instruction(RawInstrTerminateEf[:])

// InstrTerminateCardUsage represents the "TERMINATE CARD USAGE" instruction as defined in ISO/IEC 7816-4:2005(E) Part 9
InstrTerminateCardUsage = Instruction(RawInstrTerminateCardUsage[:])

)
