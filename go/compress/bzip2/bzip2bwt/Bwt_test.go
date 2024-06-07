package bzip2bwt

import (
	"bytes"
	"testing"
)

func TestBwt(t *testing.T) {
	tests := []struct {
		input    []byte
		expected []byte
	}{
		{
			input:    []byte{},
			expected: []byte("\x04"),
		},
		{
			input:    []byte(""),
			expected: []byte("\x04"),
		},
		{
			input:    []byte("banana"),
			expected: []byte("annb\x04aa"),
		},
		{
			input:    []byte("abracadabra"),
			expected: []byte("ard\x04rcaaaabb"),
		},
		{
			input:    []byte("mississippi"),
			expected: []byte("ipssm\x04pissii"),
		},
		{
			input:    []byte("test"),
			expected: []byte("ttes\x04"),
		},
		{
			input:    []byte("a"),
			expected: []byte("a\x04"),
		},
	}

	for _, tt := range tests {
		t.Run(string(tt.input), func(t *testing.T) {
			result := Bwt(tt.input)
			if !bytes.Equal(result, tt.expected) {
				t.Errorf("Bwt(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}
