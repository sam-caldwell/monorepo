package misc

import (
	"testing"
)

func TestInt64ToIntSafe(t *testing.T) {
	// Test cases
	testCases := []struct {
		input    int64
		expected int
		isError  bool
	}{
		{1234567890, 1234567890, false},                      // Within int range
		{9223372036854775807, int(^uint(0) >> 1), false},     // Max int64 value
		{-9223372036854775808, -int(^uint(0)>>1) - 1, false}, // Min int64 value
		{9223372036854775807, int(^uint(0) >> 1), false},     // Outside int range (overflow)
	}

	// Iterate over test cases
	for _, tc := range testCases {
		result, err := Int64ToIntSafe(tc.input)

		if tc.isError {
			// Expected error
			if err == nil {
				t.Errorf("Expected error for input %d, but got no error", tc.input)
			}
		} else {
			// Expected result
			if err != nil {
				t.Errorf("Unexpected error for input %d: %v", tc.input, err)
			} else if result != tc.expected {
				t.Errorf("Expected %d for input %d, but got %d", tc.expected, tc.input, result)
			}
		}
	}
}
