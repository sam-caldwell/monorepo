package rle

import (
	"bytes"
	"strings"
	"testing"
)

func TestRleIntegration(t *testing.T) {
	tests := []struct {
		input []byte
	}{
		{[]byte("AAABBBCCC")},
		{[]byte("AAAABBBCCDAA")},
		{[]byte("")},    // Empty input
		{[]byte("XYZ")}, // No repeated characters
		{[]byte("A")},
		{[]byte("BBBB")},
		{[]byte("ABCDEABCDE")},
		{[]byte(
			strings.Repeat("A", 256) +
				strings.Repeat("B", 64) +
				strings.Repeat("C", 1))},
	}

	for _, test := range tests {
		encoded := Encode(test.input)
		decoded, err := Decode(encoded)
		if err != nil {
			t.Errorf("Decode returned error for input %s: %v",
				string(test.input), err)
		}
		if !bytes.Equal(decoded, test.input) {
			t.Errorf("Encode/Decode mismatch for input %s: "+
				"got %s, "+
				"expected %s",
				string(test.input),
				string(decoded),
				string(test.input))
		}
	}
}
