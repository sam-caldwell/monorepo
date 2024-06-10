package huffman

import (
	"bytes"
	"testing"
)

func TestCodeMap_Bytes(t *testing.T) {
	testCases := []struct {
		input          CodeMap
		expectedOutput []byte
	}{
		{
			input: CodeMap{
				'a': []byte("hello"),
				'b': []byte("world"),
			},
			expectedOutput: []byte{'a', 5, 'h', 'e', 'l', 'l', 'o', 'b', 5, 'w', 'o', 'r', 'l', 'd'},
		},
		{
			input: CodeMap{
				'x': []byte("golang"),
			},
			expectedOutput: []byte{'x', 6, 'g', 'o', 'l', 'a', 'n', 'g'},
		},
		{
			input: CodeMap{
				'k': []byte{},
			},
			expectedOutput: []byte{'k', 0},
		},
		{
			input:          CodeMap{},
			expectedOutput: []byte{},
		},
	}

	for _, tc := range testCases {
		t.Run(string(tc.expectedOutput), func(t *testing.T) {
			output := tc.input.Bytes()
			if !bytes.Equal(output, tc.expectedOutput) {
				t.Fatalf("CodeMap.Bytes() = %v; want %v", output, tc.expectedOutput)
			}
		})
	}
}
