package rle

import "strconv"

// Encode - Run-Length Encoding on a []byte slice.
//
//	(c) 2023 Sam Caldwell.  MIT License
func Encode(input []byte) []byte {
	var (
		encoded []byte
		count   int
	)
	length := len(input)
	if length == 0 {
		return encoded
	}

	count = 1
	for i := 1; i < length; i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			encoded = append(encoded, input[i-1])
			encoded = append(encoded, []byte(strconv.Itoa(count))...)
			count = 1
		}
	}

	// Add the last character and its count
	encoded = append(encoded, input[length-1])
	encoded = append(encoded, []byte(strconv.Itoa(count))...)

	return encoded
}
