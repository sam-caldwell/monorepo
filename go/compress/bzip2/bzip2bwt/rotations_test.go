package bzip2bwt

import (
	"reflect"
	"testing"
)

func TestRotations(t *testing.T) {
	tests := []struct {
		input    []byte
		expected [][]byte
	}{
		{
			input: []byte("abc"),
			expected: [][]byte{
				[]byte("abc"),
				[]byte("bca"),
				[]byte("cab"),
			},
		},
		{
			input: []byte("a"),
			expected: [][]byte{
				[]byte("a"),
			},
		},
		{
			input:    []byte{},
			expected: nil,
		},
		{
			input:    []byte(""),
			expected: nil,
		},
		{
			input: []byte("aaa"),
			expected: [][]byte{
				[]byte("aaa"),
				[]byte("aaa"),
				[]byte("aaa"),
			},
		},
	}

	for _, test := range tests {
		result := rotations(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("rotations(%s) = %v; want %v", test.input, result, test.expected)
		}
	}
}
