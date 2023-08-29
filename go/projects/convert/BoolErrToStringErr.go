package convert

import "fmt"

// BoolErrToStringErr - given a bool and error, return its string, error.  (useful as a wrapper)
func BoolErrToStringErr(n bool, err error) (string, error) {
	return fmt.Sprintf("%v", n), err
}
