package moremath

import "math"

// TruncFloat32 - truncate a float to a given position
func TruncFloat32(value float32, position int) float32 {

	c := math.Pow10(position)

	return float32(math.Trunc(float64(value)*c) / c)

}
