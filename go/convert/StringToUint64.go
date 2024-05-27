package convert

import (
	"strconv"
)

// StringToUint64 - Convert the string value to a 64-bit unsigned-integer.
func StringToUint64(value string) (n uint64) {
	number, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return number
}
