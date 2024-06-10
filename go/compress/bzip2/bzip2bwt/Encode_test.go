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
			input:    []byte("a"),
			expected: []byte("\x04a\x04"),
			eof:      byte(file.EOF),
		},
		{
			input:    []byte("a\x04"),
			expected: []byte("\x05a\x04\x05"),
			eof:      byte(0x05),
		},
		{
			input:    []byte{},
			expected: []byte("\x04\x04"),
			eof:      byte(file.EOF),
		},
		{
			input:    []byte(""),
			expected: []byte("\x04\x04"),
			eof:      byte(file.EOF),
		},
		{
			input:    []byte("banana"),
			expected: []byte("\x04annb\x04aa"),
			eof:      byte(file.EOF),
		},
		{
			input:    []byte("abracadabra"),
			expected: []byte("\x04ard\x04rcaaaabb"),
			eof:      byte(file.EOF),
		},
		{
			input:    []byte("mississippi"),
			expected: []byte("\x04ipssm\x04pissii"),
			eof:      byte(file.EOF),
		},
		{
			input:    []byte("test"),
			expected: []byte("\x04ttes\x04"),
			eof:      byte(file.EOF),
		},
	}

	for i, tt := range tests {
		t.Run(string(tt.input), func(t *testing.T) {
			result := Bwt(tt.input)
			if !bytes.Equal(result, tt.expected) {
				t.Fatalf("test %d:"+
					"Bwt(%q)\n"+
					"   got %q\n"+
					"expected %q", i, tt.input, result, tt.expected)
			}
			if eof := result[0]; eof != tt.eof {
				t.Fatalf("test %d:"+
					"Bwt(%q)\n"+
					"   eof = %q\n"+
					"expected %q\n", i, tt.input, eof, tt.eof)
			}
		})
	}
}
