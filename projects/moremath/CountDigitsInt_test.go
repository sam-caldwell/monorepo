package moremath

import "testing"

func TestCountDigitsInt(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{input: 0, expected: 1},
		{input: 123, expected: 3},
		{input: -4567, expected: 4},
		{input: 9876543210, expected: 10},
	}

	for _, tc := range testCases {
		result := CountDigitsInt(tc.input)
		if result != tc.expected {
			t.Errorf("For input %d, expected %d digits but got %d", tc.input, tc.expected, result)
		}
	}
}
