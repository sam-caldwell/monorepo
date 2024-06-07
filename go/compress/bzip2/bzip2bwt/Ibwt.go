package bzip2bwt

import (
	"bytes"
	"sort"
)

// Ibwt - Inverse Burrows-Wheeler Transform
//
//		This function performs the Inverse Burrows-Wheeler transformation
//		It's opposite is Bwt()
//
//	 1. Initialize a slice to store the sorted characters of the input and a count array
//	    to count occurrences of each character in the input.
//
//	 2. Construct the sortedChars array by counting occurrences of each character in the input.
//
//	 3. Iterate through the sortedChars array to perform the inverse transformation and
//	    reconstruct the original input.
//
//	 4. Update the original index based on the occurrences of the character in the input.
//
//	    (c) 2023 Sam Caldwell.  MIT License
func Ibwt(r []byte, eof byte) []byte {
	// Initialize table with empty rows
	table := make([][]byte, len(r))
	for i := range table {
		table[i] = make([]byte, len(r))
	}

	// Reconstruct the table iteratively
	for i := 0; i < len(r); i++ {
		// Prepend the input characters to each row
		for j := range table {
			table[j][len(r)-i-1] = r[j]
		}

		// Sort the rows lexicographically
		sort.Slice(table, func(a, b int) bool {
			return bytes.Compare(table[a], table[b]) < 0
		})
	}

	// Find the row that ends with the EOF character
	var result []byte
	for _, row := range table {
		if row[len(r)-1] == eof {
			result = row
			break
		}
	}

	// Remove the EOF character from the result
	if len(result) > 0 && result[len(r)-1] == eof {
		result = result[:len(r)-1]
	}

	return result
}
