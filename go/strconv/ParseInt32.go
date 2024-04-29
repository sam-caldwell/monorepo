package strconv

import "strconv"

// ParseInt32 - Wrapper for golang strconv.ParseInt because I am tired of refactoring deprecated stuff golang changes
func ParseInt32(s string) (int32, error) {
	value, err := strconv.ParseInt(s, 10, 32)
	return int32(value), err
}
