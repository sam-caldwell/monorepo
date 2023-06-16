package types

import "fmt"

// Valid - Return true if ArgTypes is valid or false if ArgTypes is not valid.
func (arg *ArgTypes) Valid() (result error) {
	switch *arg {
	case String:
		fallthrough
	case Integer:
		fallthrough
	case Boolean:
		fallthrough
	case Float:
		fallthrough
	case Flag:
		result = nil
	default:
		result = fmt.Errorf(errInvalidArgType, arg.String())
	}
	return result
}
