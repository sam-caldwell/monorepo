package strconv

import "strconv"

// ParseInt16 - Wrapper for golang strconv.ParseInt because I am tired of refactoring deprecated stuff golang changes
func ParseInt16(s string) (int16, error) {
	value, err := strconv.ParseInt(s, 10, 16)
	return int16(value), err
}
