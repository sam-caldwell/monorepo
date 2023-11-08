package moremath

/*
 * sigmoid() function
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A simple sigmoid function.
 */

import "math"

// Sigmoid activation function
func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}
