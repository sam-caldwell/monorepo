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
func Bwt(input []byte) ([]byte, int) {
	n := len(input)
	// Create a slice to store all rotations
	rot := rotations(input)

	// Sort the rotations
	sort.Slice(rot, func(i, j int) bool {
		return bytes.Compare(rot[i], rot[j]) < 0
	})

	// Create the BWT result
	bwt := make([]byte, n)
	originalIndex := -1
	for i, r := range rot {
		// Find the index of the original input
		if bytes.Equal(r, input) {
			originalIndex = i
			break
		}
	}

	// Construct the BWT result
	for i, r := range rot {
		bwt[i] = r[n-1]
	}

	return bwt, originalIndex
}
