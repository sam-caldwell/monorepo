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
func Bwt(data []byte) ([]byte, int) {
	n := len(data)
	if n == 0 {
		return nil, 0
	}

	// Create a slice of rotations
	rotations := make([][]byte, n)
	for i := range rotations {
		rotations[i] = append(data[i:], data[:i]...)
	}

	// Sort the rotations lexicographically
	sort.Slice(rotations, func(i, j int) bool {
		return bytes.Compare(rotations[i], rotations[j]) < 0
	})

	// Extract the last column and find the original index
	lastColumn := make([]byte, n)
	var originalIndex int
	for i, rotation := range rotations {
		lastColumn[i] = rotation[n-1]
		if bytes.Equal(rotation, data) {
			originalIndex = i
		}
	}

	return lastColumn, originalIndex
}
