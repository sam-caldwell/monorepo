package moremath

import (
	"math"
	"testing"
)

func TestRoundFloat32(t *testing.T) {
	tests := []struct {
		value    float32
		position int
		expected float32
	}{
		{float32(3.14159), 2, float32(3.14)},
		{float32(2.71828), 3, float32(2.718)},
		{float32(0.12345), 4, float32(0.1235)},
		{float32(-1.234567), 3, float32(-1.235)},
	}

	for _, test := range tests {
		const tolerance = 1e-6
		result := RoundFloat32(test.value, test.position)
		if math.Abs(float64(result-test.expected)) > tolerance {
			t.Fatalf("RoundFloat64(%f, %d) = %f, expected %f", test.value, test.position, result, test.expected)
		}
	}
}
