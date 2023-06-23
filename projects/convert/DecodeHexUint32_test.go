package convert

import (
	"testing"
)

func TestDecodeHexUint32(t *testing.T) {
	tests := []struct {
		src         []byte
		expected    uint32
		shouldError bool
	}{
		{src: []byte{'0', '0', '0', '0', '0', '0', '0', '0'}, expected: 0, shouldError: false},
		{src: []byte{'F', 'F', '0', '0', 'A', '5', 'C', 'D'}, expected: 4278232525, shouldError: false},
		{src: []byte{'1', '2', '3', '4', '5', '6', '7', '8'}, expected: 305419896, shouldError: false},
		{src: []byte{'X', '5', 'A', 'F', '2', 'B', 'C', 'D'}, shouldError: true}, // Invalid hexadecimal digit
		{src: []byte{'9', '@', 'A', 'B', 'C', 'D', 'E', 'F'}, shouldError: true}, // Invalid hexadecimal digit
	}

	for _, test := range tests {
		var result uint32
		ok := DecodeHexUint32(&result, test.src)

		if test.shouldError {
			if ok {
				t.Fatalf("Expected failure for input '%s', but got ok", string(test.src))
			}
		} else {
			if !ok {
				t.Fatalf("Expected ok but failed on input '%s'", string(test.src))
			} else if result != test.expected {
				t.Fatalf("Unexpected result for input '%s'.\n"+
					"\tExpected: %d,\n"+
					"\tGot: %d\n",
					string(test.src), test.expected, result)
			}
		}
	}
}
