package bzip2bwt

import (
	"bytes"
	"testing"
)

func TestBwt(t *testing.T) {
	tests := []struct {
		input         []byte
		expectedBwt   []byte
		expectedIndex int
	}{
		{
			input:         []byte("banana"),
			expectedBwt:   []byte("nnbaaa"),
			expectedIndex: 3,
		},
		{
			input:         []byte("example"),
			expectedBwt:   []byte("xlepame"), // Corrected expected BWT output
			expectedIndex: 2,                 // Corrected expected index
		},
		{
			input:         []byte("abcde"),
			expectedBwt:   []byte("eabcd"),
			expectedIndex: 0,
		},
		{
			input:         []byte("a"),
			expectedBwt:   []byte("a"),
			expectedIndex: 0,
		},
		{
			input:         []byte(""),
			expectedBwt:   []byte{},
			expectedIndex: -1,
		},
	}

	for _, test := range tests {
		bwtResult, originalIndex := Bwt(test.input)
		if !bytes.Equal(bwtResult, test.expectedBwt) || originalIndex != test.expectedIndex {
			t.Errorf("Bwt(%s)\n"+
				"Got:  (%s, %d)\n"+
				"want: (%s, %d)",
				test.input,
				bwtResult, originalIndex,
				test.expectedBwt, test.expectedIndex)
		}
	}
}
