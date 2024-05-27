package convert

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"math"
	"strconv"
)

// StringToUint - Convert the string value to an unsigned-integer.
//
//	(c) 2023 Sam Caldwell.  MIT License
func StringToUint(value string) (n uint) {
	number, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		panic(err)
	}
	if number <= math.MaxUint32 {
		return uint(number)
	}
	panic(errors.BoundsCheckError)
}
