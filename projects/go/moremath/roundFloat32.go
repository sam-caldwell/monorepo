package moremath

import "math"

// RoundFloat32 - round a float to a given decimal position
func RoundFloat32(value float32, position int) float32 {
	c := math.Pow10(position)
	rounded := math.Round(float64(value)*c) / c
	return float32(rounded)
}
