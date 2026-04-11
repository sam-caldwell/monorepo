package convert

import (
	"testing"
)

func TestDecodeHexByte(t *testing.T) {
	tests := []struct {
		c1, c2   byte
		expected byte
		ok       bool
	}{
		{c1: '0', c2: '0', expected: 0x00, ok: true},
		{c1: '5', c2: 'a', expected: 0x5A, ok: true},
		{c1: '5', c2: 'A', expected: 0x5A, ok: true},
		{c1: 'f', c2: 'F', expected: 0xFF, ok: true},
		{c1: 'x', c2: '5', expected: 0x00, ok: false}, // Invalid hexadecimal digit
		{c1: '9', c2: '@', expected: 0x00, ok: false}, // Invalid hexadecimal digit
	}

	for _, test := range tests {
		var result byte
		ok := DecodeHexByte(&result, test.c1, test.c2)
		if test.ok {
			if !ok {
				t.Fatalf("Expected Ok but got failure on '%c%c'", test.c1, test.c2)
			}
			if result != test.expected {
				t.Fatalf("Expected ok but failed on input on '%c%c'\n"+
					"\tExpected: '%d'\n"+
					"\t     Got: '%d'\n", test.c1, test.c2, test.expected, result)
			}
		} else { // ShouldError == false

			if ok {
				t.Fatalf("Test is ok but it should not be on %c%c", test.c1, test.c2)
			}
		}
	}
}
