package bzip2bwt

import (
	"bytes"
	"sort"
)

// Bwt - Burrows-Wheeler Transform
//
//	 This function performs a Burrows-Wheeler transformation
//	 It's opposite is Ibwt()
//
//		(c) 2023 Sam Caldwell.  MIT License
func Bwt(input []byte) (codedText []byte, eof byte) {
	// Scan the input to find a unique EOF character
	eof = findUniqueEOF(input)

	// Append the EOF character to the input
	input = append(input, eof)

	length := len(input)

	// Create a slice to hold all rotations
	rotations := make([][]byte, length)

	// Generate all rotations of the input string
	for i := 0; i < length; i++ {
		rotations[i] = append(input[i:], input[:i]...)
	}

	// Sort the rotations lexicographically
	sort.Slice(rotations, func(i, j int) bool {
		return bytes.Compare(rotations[i], rotations[j]) < 0
	})

	// Extract the last column
	lastColumn := make([]byte, length)
	for i := 0; i < length; i++ {
		lastColumn[i] = rotations[i][length-1]
	}

	return lastColumn, eof
}
