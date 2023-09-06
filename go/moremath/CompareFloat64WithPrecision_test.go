package moremath

import (
	"testing"
)

func TestCompareFloat64WithPrecision(t *testing.T) {
	tests := []struct {
		a, b                float64
		numberDecimalPlaces int
		expectedResult      bool
	}{
		{3.14159, 3.1415, 3, true},
		{2.71828, 2.71838, 4, false},
		{0.12345, 0.123456, 5, true},
		{-1.2345675, -1.2345674, 6, true},
		{1.23e10, 1.23e10, 8, true},
		{0.001, 0.001, 3, true},
		{0.001, 0.002, 3, false},
		{0, 0, 2, true},
	}

	for _, test := range tests {
		result := CompareFloat64WithPrecision(test.a, test.b, test.numberDecimalPlaces)
		if result != test.expectedResult {
			t.Errorf("CompareFloat64WithPrecision(%f, %f, %d) = %t, expected %t", test.a, test.b, test.numberDecimalPlaces, result, test.expectedResult)
		}
	}
}
