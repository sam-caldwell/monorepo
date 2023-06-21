package convert

import "fmt"

// IntToString - convert an integer to a string
func IntToString(fn func() (int, error)) (string, error) {
	raw, err := fn()
	return fmt.Sprintf("%d", raw), err
}
