package convert

import "fmt"

type NumericType interface {
	byte | int | int8 | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | float32 | float64
}

// NumericErrToStringErr - given a numeric type and error, return its string, error.  (useful as a wrapper)
func NumericErrToStringErr[V NumericType](n V, err error) (string, error) {
	return fmt.Sprintf("%v", n), err
}
