package strconv

import "strconv"

// ParseInt64 - Wrapper for golang strconv.ParseInt because I am tired of refactoring deprecated stuff golang changes
func ParseInt64(s string) (int64, error) {
	value, err := strconv.ParseInt(s, 10, 64)
	return value, err
}
