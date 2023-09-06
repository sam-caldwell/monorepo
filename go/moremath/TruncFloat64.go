package moremath

import "math"

// TruncFloat64 - truncate a float to a given position
func TruncFloat64(value float64, position int) float64 {

	shift := math.Pow10(position)

	return math.Trunc(value*shift) / shift

}
