package rle

import (
	"bytes"
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{"Empty input", []byte(""), []byte{}},
		{"Nil input", []byte{}, []byte{}},
		{"simple repeated sequences1", []byte("AAABBBCCC"), []byte("A3B3C3")},
		{"simple repeated sequences2", []byte("AAAABBBCCDAA"), []byte("A4B3C2D1A2")},
		{"complex repeated sequences",
			[]byte(
				strings.Repeat("A", 256) +
					strings.Repeat("B", 64) +
					strings.Repeat("C", 1)),
			[]byte("A256B64C1"),
		},
		{"no repeated characters", []byte("XYZ"), []byte("X1Y1Z1")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Encode(test.input)
			if !bytes.Equal(result, test.expected) {
				t.Fatalf("Encode('%s') not equal\n"+
					"returned '%v'\n"+
					"expected '%v'", test.input, result, test.expected)
			}
		})
	}
}
