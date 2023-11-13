package convert

import "fmt"

// NumericErrToStringErr - given a numeric type and error, return its string, error.  (useful as a wrapper)
func NumericErrToStringErr[V Number](n V, err error) (string, error) {
	return fmt.Sprintf("%v", n), err
}
