package types

import "fmt"

// String - Return a string version of an ArgTypes
func (arg *ArgTypes) String() (result string) {
	switch *arg {
	case String:
		result = "String"
	case Integer:
		result = "Integer"
	case Boolean:
		result = "Boolean"
	case Float:
		result = "Float"
	case Flag:
		result = "Flag"
	default:
		result = fmt.Sprintf(errUnknownType, *arg)
	}
	return result
}
