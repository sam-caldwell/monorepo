package convert

import "fmt"

// IntToStringFuncWrapper - convert an integer to a string
//
//	(c) 2023 Sam Caldwell.  MIT License
func IntToStringFuncWrapper(fn func() (int, error)) (string, error) {
	raw, err := fn()
	return fmt.Sprintf("%d", raw), err
}
