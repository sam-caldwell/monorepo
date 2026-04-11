package strconv

import "strconv"

// ParseUInt64 - Wrapper for golang strconv.ParseUInt because I am tired of refactoring deprecated stuff golang changes
func ParseUInt64(s string) (uint64, error) {
	value, err := strconv.ParseUint(s, 10, 64)
	return value, err
}
