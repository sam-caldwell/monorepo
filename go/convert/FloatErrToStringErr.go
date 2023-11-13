package convert

import "fmt"

// FloatErrToStringErr - given a float(32 or 64 and error, return its string, error.  (useful as a wrapper)
func FloatErrToStringErr[V float32 | float64](n V, err error) (string, error) {
	return fmt.Sprintf("%f", n), err
}
