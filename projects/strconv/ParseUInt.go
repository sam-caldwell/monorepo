package strconv

import "strconv"

// ParseUInt - Wrapper for golang strconv.ParseUInt because I am tired of refactoring deprecated stuff golang changes
func ParseUInt(s string) (uint, error) {
	value, err := strconv.ParseUint(s, 10, 32)
	return uint(value), err
}
