package rle

import (
	"bytes"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"testing"
)

func TestDecode(t *testing.T) {
	malformed := fmt.Errorf(errors.MalformedInput)
	tests := []struct {
		name     string
		input    []byte
		expected []byte
		err      error
	}{
		// Happy path
		{"happy: test1", []byte("A3B3C3"), []byte("AAABBBCCC"), nil},
		{"happy: test2", []byte("A4B3C2D1A2"), []byte("AAAABBBCCDAA"), nil},
		{"happy: No repeated characters", []byte("X1Y1Z1"), []byte("XYZ"), nil},

		// Sad path (malformed input)
		{"sad: Missing count for last character", []byte("A3B3C"), []byte{}, malformed},
		{"sad: Missing count after 2nd character", []byte("A3BCC"), []byte{}, malformed},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Decode(test.input)
			if err != nil {
				if test.err == nil {
					t.Fatalf("Decode(%q): got unexpected error: %v", test.name, err)
				}
				if err.Error() != test.err.Error() {
					t.Fatalf("Decode(%q):\n"+
						"got error %v,\n"+
						"wanted error %v",
						test.name, err, test.err)
				}
			}
			if !bytes.Equal(result, test.expected) {
				t.Fatalf("Decode(%s)\n"+
					"returned '%v'\n"+
					"expected '%v'",
					string(test.input), result, test.expected)
			}
		})
	}
}
