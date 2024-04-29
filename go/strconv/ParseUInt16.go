package strconv

import "strconv"

// ParseUInt16 - Wrapper for golang strconv.ParseUInt because I am tired of refactoring deprecated stuff golang changes
func ParseUInt16(s string) (uint16, error) {
	value, err := strconv.ParseUint(s, 10, 16)
	return uint16(value), err
}
