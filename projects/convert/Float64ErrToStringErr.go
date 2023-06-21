package convert

import "fmt"

// Float64ErrToStringErr - given a float64 and error, return its string, error.  (useful as a wrapper)
func Float64ErrToStringErr(n float64, err error) (string, error) {
	return fmt.Sprintf("%f", n), err
}
