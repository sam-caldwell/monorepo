package convert

import (
	"testing"
)

func TestDecodeHexChar(t *testing.T) {
	tests := []struct {
		hexByte     byte
		expected    byte
		shouldError bool
	}{
		{hexByte: '0', expected: 0, shouldError: false},
		{hexByte: '5', expected: 5, shouldError: false},
		{hexByte: 'A', expected: 10, shouldError: false},
		{hexByte: 'f', expected: 15, shouldError: false},
		{hexByte: 'x', shouldError: true}, // Invalid hexadecimal digit
		{hexByte: '@', shouldError: true}, // Invalid hexadecimal digit
	}

	for _, test := range tests {
		var result byte
		ok := DecodeHexChar(&result, test.hexByte)

		if test.shouldError {
			if ok {
				t.Errorf("Expected failure on input '%c', but got ok", test.hexByte)
			}
		} else {
			if !ok {
				t.Errorf("Expected ok but failed on '%c'", test.hexByte)
			} else if result != test.expected {
				t.Errorf("Unexpected result for input '%c'. Expected: %d, Got: %d",
					test.hexByte, test.expected, result)
			}
		}
	}
}
