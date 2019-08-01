package apdu

//go:generate go run github.com/zemnmez/cardauth/apdu/gen constants --data gen/resources/P1.csv --go $GOFILE --out $GOFILE --comment "{{.OriginalName}}; {{.Info}}" --infoMap p1Info --prefix P1 --reverseStringMap p1ReverseStringMap --secondariesMap p1SecondariesMap --stringMap p1String --type P1
var _ = 0

const (
	P1NormalProcessing  P1 = 0x61 // Normal Processing; SW2 encodes the number of data bytes still available
	P1WarningProcessing P1 = 0x62 // Warning Processing; State of non-volatile memory is unchanged
	P1WarningProcessing P1 = 0x63 // Warning Processing; State of non-volatile memory has changed
	P1ExecutionError    P1 = 0x64 // Execution Error; State of non-volatile memory is unchanged
	P1ExecutionError    P1 = 0x65 // Execution Error; State of non-volatile memory has changed
	P1ExecutionError    P1 = 0x66 // Execution Error; Security-related issues
	P1CheckingError     P1 = 0x67 // Checking Error; Wrong length; no further indication
	P1CheckingError     P1 = 0x68 // Checking Error; Functions in CLA not supported Functions in CLA not supported
	P1CheckingError     P1 = 0x69 // Checking Error; Command not allowed
	P1CheckingError     P1 = 0x6a // Checking Error; Wrong parameters P1-P2
	P1CheckingError     P1 = 0x6b // Checking Error; Wrong parameters P1-P2
	P1CheckingError     P1 = 0x6c // Checking Error; Wrong Le field; SW2 encodes the exact number of available data bytes (see text below)
	P1CheckingError     P1 = 0x6d // Checking Error; Instruction code not supported or invalid
	P1CheckingError     P1 = 0x6e // Checking Error; Class not supported
	P1CheckingError     P1 = 0x6f // Checking Error; No precise diagnosis
	_                      = 0
)

var (
	p1Info             = map[P1]string{P1NormalProcessing: "Normal Processing; SW2 encodes the number of data bytes still available", P1WarningProcessing: "Warning Processing; State of non-volatile memory is unchanged", P1WarningProcessing: "Warning Processing; State of non-volatile memory has changed", P1ExecutionError: "Execution Error; State of non-volatile memory is unchanged", P1ExecutionError: "Execution Error; State of non-volatile memory has changed", P1ExecutionError: "Execution Error; Security-related issues", P1CheckingError: "Checking Error; Wrong length; no further indication", P1CheckingError: "Checking Error; Functions in CLA not supported Functions in CLA not supported", P1CheckingError: "Checking Error; Command not allowed", P1CheckingError: "Checking Error; Wrong parameters P1-P2", P1CheckingError: "Checking Error; Wrong parameters P1-P2", P1CheckingError: "Checking Error; Wrong Le field; SW2 encodes the exact number of available data bytes (see text below)", P1CheckingError: "Checking Error; Instruction code not supported or invalid", P1CheckingError: "Checking Error; Class not supported", P1CheckingError: "Checking Error; No precise diagnosis"}
	p1ReverseStringMap = map[string]P1{"P1NormalProcessing": P1NormalProcessing, "P1WarningProcessing": P1WarningProcessing, "P1WarningProcessing": P1WarningProcessing, "P1ExecutionError": P1ExecutionError, "P1ExecutionError": P1ExecutionError, "P1ExecutionError": P1ExecutionError, "P1CheckingError": P1CheckingError, "P1CheckingError": P1CheckingError, "P1CheckingError": P1CheckingError, "P1CheckingError": P1CheckingError, "P1CheckingError": P1CheckingError, "P1CheckingError": P1CheckingError, "P1CheckingError": P1CheckingError, "P1CheckingError": P1CheckingError, "P1CheckingError": P1CheckingError}
	p1String           = map[P1]string{P1NormalProcessing: "P1NormalProcessing", P1WarningProcessing: "P1WarningProcessing", P1WarningProcessing: "P1WarningProcessing", P1ExecutionError: "P1ExecutionError", P1ExecutionError: "P1ExecutionError", P1ExecutionError: "P1ExecutionError", P1CheckingError: "P1CheckingError", P1CheckingError: "P1CheckingError", P1CheckingError: "P1CheckingError", P1CheckingError: "P1CheckingError", P1CheckingError: "P1CheckingError", P1CheckingError: "P1CheckingError", P1CheckingError: "P1CheckingError", P1CheckingError: "P1CheckingError", P1CheckingError: "P1CheckingError"}
	p1SecondariesMap   = map[P1]P1{}
)
