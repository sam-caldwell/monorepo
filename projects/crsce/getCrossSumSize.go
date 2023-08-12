package crsce

import "math"

// getCrossSumSize Calculate the size of a cross sums matrix.
func getCrossSumSize(sz int) int {
	return int(math.Ceil(math.Sqrt(float64(sz))))
}
