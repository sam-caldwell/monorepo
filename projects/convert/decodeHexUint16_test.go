package convert

import (
	"testing"
)

func TestDecodeHexUint16(t *testing.T) {
	tests := []struct {
		src         []byte
		expected    uint16
		shouldError bool
	}{
		{src: []byte{'0', '0', '0', '0'}, expected: 0, shouldError: false},
		{src: []byte{'F', 'F', '0', '0'}, expected: 65280, shouldError: false},
		{src: []byte{'1', '2', '3', '4'}, expected: 4660, shouldError: false},
		{src: []byte{'X', '5', 'A', 'F'}, shouldError: true}, // Invalid hexadecimal digit
		{src: []byte{'9', '@', 'A', 'B'}, shouldError: true}, // Invalid hexadecimal digit
	}

	for _, test := range tests {
		var result uint16
		ok := decodeHexUint16(&result, test.src)

		if test.shouldError {
			if ok {
				t.Fatalf("Expected failure for input '%s', but got ok", string(test.src))
			}
		} else {
			if !ok {
				t.Fatalf("Expected ok but failed on input '%s'", string(test.src))
			} else if result != test.expected {
				t.Fatalf("Unexpected result for input '%s'. Expected: %d, Got: %d",
					string(test.src), test.expected, result)
			}
		}
	}
}
