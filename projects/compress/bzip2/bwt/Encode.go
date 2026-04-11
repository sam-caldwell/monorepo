package bwt

import (
	"bytes"
	"sort"
)

// Encode - Burrows-Wheeler Transformation Encode
//
//	(c) 2023 Sam Caldwell.  MIT License
func Encode(input []byte) (output []byte, index int) {
	length := len(input)
	rotations := make([][]byte, length)
	for i := range rotations {
		rotations[i] = make([]byte, length)
		copy(rotations[i], input[i:])
		copy(rotations[i][length-i:], input[:i])
	}

	// Sort the rotations
	sorted := make([][]byte, length)
	copy(sorted, rotations)
	sort.Slice(sorted, func(i, j int) bool {
		return bytes.Compare(sorted[i], sorted[j]) < 0
	})

	output = make([]byte, length)
	for i := range sorted {
		output[i] = sorted[i][length-1]
		if bytes.Equal(sorted[i], input) {
			index = i
		}
	}
	return output, index
}
