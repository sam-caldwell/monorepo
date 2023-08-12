package main

import (
	"fmt"
	"math"
)

func equation(n float64) float64 {
	return n - 2*math.Ceil(math.Log2(n)) - 32
}

func bisectionMethod(a, b, epsilon float64, maxIterations int) (float64, error) {
	if equation(a)*equation(b) >= 0 {
		return 0, fmt.Errorf("initial values do not bracket the root")
	}

	var c float64

	for i := 0; i < maxIterations; i++ {
		c = (a + b) / 2
		if equation(c) == 0 || (b-a)/2 < epsilon {
			break
		}
		if equation(c)*equation(a) < 0 {
			b = c
		} else {
			a = c
		}
	}

	return c, nil
}

func main() {
	a := float64(1)      // Lower bound of the interval
	b := float64(1000)   // Upper bound of the interval
	epsilon := 1e-6      // Tolerance for convergence
	maxIterations := 100 // Maximum number of iterations

	root, err := bisectionMethod(a, b, epsilon, maxIterations)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	minimumSize := math.Ceil(root)
	fmt.Printf("minimum size of n: %0.f\n", minimumSize)
}
