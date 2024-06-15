package bwt

import (
	"bytes"
	"sort"
)

// Decode - Burrows-Wheeler Transformation Decode
//
//	(c) 2023 Sam Caldwell.  MIT License
func Decode(output []byte, index int) []byte {
	length := len(output)
	table := make([][]byte, length)
	for i := range table {
		table[i] = make([]byte, length)
	}

	for i := 0; i < length; i++ {
		for j := range table {
			table[j][0] = output[j]
		}
		sort.Slice(table, func(i, j int) bool {
			return bytes.Compare(table[i], table[j]) < 0
		})
	}

	result := make([]byte, length)
	for i, row := range table {
		result[i] = row[length-1]
	}
	return result
}
