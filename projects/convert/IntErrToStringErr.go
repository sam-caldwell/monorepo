package convert

import "fmt"

// IntErrToStringErr - given an int and error, return its string, error.  (useful as a wrapper)
//
//	(c) 2023 Sam Caldwell.  MIT License
func IntErrToStringErr(n int, err error) (string, error) {
	return fmt.Sprintf("%d", n), err
}
