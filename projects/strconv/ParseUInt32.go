package strconv

import "strconv"

// ParseUInt32 - Wrapper for golang strconv.ParseUInt because I am tired of refactoring deprecated stuff golang changes
func ParseUInt32(s string) (uint32, error) {
	value, err := strconv.ParseUint(s, 10, 32)
	return uint32(value), err
}
