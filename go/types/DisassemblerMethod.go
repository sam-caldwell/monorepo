package types

import "fmt"

// DisassemblerMethod - Generic type for the disassembly method to be used (recursive, linear)
type DisassemblerMethod int

const (
	Linear DisassemblerMethod = iota
	Recursive
)

const (
	StrMethodLinear    string = "linear"
	StrMethodRecursive string = "recursive"
)

// String - Return the DisassemblerMethod value
func (m *DisassemblerMethod) String() string {
	switch *m {
	case Linear:

		return StrMethodLinear
	case Recursive:

		return StrMethodRecursive
	default:

		panic("DisassemblerMethod has invalid state")
	}
}

// FromStringPtr - Cast String to numeric value given a string pointer
func (m *DisassemblerMethod) FromStringPtr(methodString *string) error {
	switch *methodString {
	case StrMethodLinear:
		*m = Linear
	case StrMethodRecursive:
		*m = Recursive
	default:
		return fmt.Errorf("invalid method (%s)", *methodString)
	}
	return nil
}
