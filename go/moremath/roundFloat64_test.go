package moremath

import (
	"math"
	"testing"
)

func TestRoundFloat64(t *testing.T) {
	tests := []struct {
		value    float64
		position int
		expected float64
	}{
		{3.14159, 2, 3.14},
		{2.71828, 3, 2.718},
		{0.12345, 4, 0.1235},
		{-1.234567, 3, -1.235},
	}

	for _, test := range tests {
		result := RoundFloat64(test.value, test.position)
		if math.Abs(result-test.expected) > 1e-6 {
			t.Fatalf("RoundFloat64(%f, %d) = %f, expected %f", test.value, test.position, result, test.expected)
		}
	}
}
