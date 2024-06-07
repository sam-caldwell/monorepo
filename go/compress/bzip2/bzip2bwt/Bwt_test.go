package bzip2bwt

import (
	"bytes"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"testing"
)

func TestBwt(t *testing.T) {
	tests := []struct {
		input    []byte
		expected []byte
		eof      byte
	}{
		{
			input:    []byte{},
			expected: []byte("\x04"),
			eof:      byte(file.EOF),
		},
		{
			input:    []byte(""),
			expected: []byte("\x04"),
			eof:      byte(file.EOF),
		},
		{
			input:    []byte("banana"),
			expected: []byte("annb\x04aa"),
			eof:      byte(file.EOF),
		},
		{
			input:    []byte("abracadabra"),
			expected: []byte("ard\x04rcaaaabb"),
			eof:      byte(file.EOF),
		},
		{
			input:    []byte("mississippi"),
			expected: []byte("ipssm\x04pissii"),
			eof:      byte(file.EOF),
		},
		{
			input:    []byte("test"),
			expected: []byte("ttes\x04"),
			eof:      byte(file.EOF),
		},
		{
			input:    []byte("a"),
			expected: []byte("a\x04"),
			eof:      byte(file.EOF),
		},
		{
			input:    []byte("a\x04"),
			expected: []byte("a\x04\x05"),
			eof:      byte(0x05),
		},
	}

	for _, tt := range tests {
		t.Run(string(tt.input), func(t *testing.T) {
			result, eof := Bwt(tt.input)
			if !bytes.Equal(result, tt.expected) {
				t.Errorf("Bwt(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
			if eof != tt.eof {
				t.Errorf("Bwt(%q) eof = %q, expected %q", tt.input, eof, tt.eof)
			}
		})
	}
}
