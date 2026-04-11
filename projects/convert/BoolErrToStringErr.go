package convert

import "fmt"

// BoolErrToStringErr - given a bool and error, return its string, error.  (useful as a wrapper)
//
//	(c) 2023 Sam Caldwell.  MIT License
func BoolErrToStringErr(n bool, err error) (string, error) {
	return fmt.Sprintf("%v", n), err
}
