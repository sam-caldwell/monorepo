package convert

import "fmt"

// IntToStringFuncWrapper - convert an integer to a string
func IntToStringFuncWrapper(fn func() (int, error)) (string, error) {
	raw, err := fn()
	return fmt.Sprintf("%d", raw), err
}
