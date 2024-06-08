package bzip2bwt

import (
	"bytes"
	"testing"
)

func TestIbwt(t *testing.T) {
	tests := []struct {
		input    []byte
		eof      byte
		expected []byte
	}{
		{
			input:    []byte("\x04annb\x04aa"),
			eof:      byte(0x04),
			expected: []byte("banana"),
		},
		{
			input:    []byte("\x04ard\x04rcaaaabb"),
			eof:      byte(0x04),
			expected: []byte("abracadabra"),
		},
		{
			input:    []byte("\x04ipssm\x04pissii"),
			eof:      byte(0x04),
			expected: []byte("mississippi"),
		},
		{
			input:    []byte("\x04ttes\x04"),
			eof:      byte(0x04),
			expected: []byte("test"),
		},
		{
			input:    []byte("\x04a\x04"),
			eof:      byte(0x04),
			expected: []byte("a"),
		},
	}

	for _, tt := range tests {
		t.Run(string(tt.input), func(t *testing.T) {
			result := Ibwt(tt.input)
			if !bytes.Equal(result, tt.expected) {
				t.Fatalf("Ibwt(%q, %q) = %q, expected %q", tt.input, tt.eof, result, tt.expected)
			}
		})
	}
}
