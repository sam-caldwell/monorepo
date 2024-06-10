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
func Ibwt(input []byte) []byte {
	if input == nil || len(input) == 0 {
		return []byte{}
	}
	eof := input[0]
	encodedText := input[1:]

	// Initialize table with empty rows
	table := make([][]byte, len(encodedText))
	for i := range table {
		table[i] = make([]byte, len(encodedText))
	}

	// Reconstruct the table iteratively
	for i := 0; i < len(encodedText); i++ {
		// Prepend the input characters to each row
		for j := range table {
			table[j][len(encodedText)-i-1] = encodedText[j]
		}

		// Sort the rows lexicographically
		sort.Slice(table, func(a, b int) bool {
			return bytes.Compare(table[a], table[b]) < 0
		})
	}

	// Find the row that ends with the EOF character
	var result []byte
	for _, row := range table {
		if row[len(encodedText)-1] == eof {
			result = row
			break
		}
	}

	// Remove the EOF character from the result
	if len(result) > 0 && result[len(encodedText)-1] == eof {
		result = result[:len(encodedText)-1]
	}

	return result
}
