package bzip2bwt

import (
	"testing"
)

func TestFindUniqueEOF(t *testing.T) {
	tests := []struct {
		input    []byte
		expected byte
	}{
		{[]byte("BANANA"), byte(0x04)},       // 0x04 should not be in "BANANA" (return default)
		{[]byte("PANAMA"), byte(0x04)},       // 0x04 should not be in "PANAMA" (return default)
		{[]byte("APPLE"), byte(0x04)},        // 0x04 should not be in "APPLE" (return default)
		{[]byte("BOOKKEEPER"), byte(0x04)},   // 0x04 should not be in "BOOKKEEPER" (return default)
		{[]byte("HE$LO"), byte(0x04)},        // 0x04 should not be in "HE$LO" (return default)
		{[]byte{65, 66, 04, 67}, byte(0x05)}, // find 0x04 and expect 0x05
	}

	for _, test := range tests {
		result := findUniqueEOF(test.input)
		if result != test.expected {
			t.Errorf("findUniqueEOF(%s) = %q; expected %q", test.input, result, test.expected)
		}

		// Verify that the selected EOF character is not present in the input
		for _, char := range test.input {
			if char == result {
				t.Errorf("EOF character %q is present in the input %s", result, test.input)
			}
		}
	}
}
