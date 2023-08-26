package moremath

import "math"

// RoundFloat64 - round a float to a given decimal position
func RoundFloat64(value float64, position int) float64 {

	c := math.Pow10(position)

	return math.Round(value*c) / c

}
