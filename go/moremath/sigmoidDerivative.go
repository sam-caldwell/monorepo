package moremath

/*
 * sigmoidDerivative() function
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A simple sigmoid derivative function.
 */

// SigmoidDerivative - derivative of the sigmoid function
func SigmoidDerivative(x float64) float64 {
	return x * (1.0 - x)
}
