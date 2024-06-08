package bzip2mtf

import "bytes"

// Imtf - perform the inverse move-to-front transformation on the input byte slice.
//
// (c) 2023 Sam Caldwell.  MIT License
func Imtf(compressedText []byte) string {
	// Make a copy of the common dictionary
	dictionary := make([]byte, 256)
	copy(dictionary, commonDictionary[:])

	var decodedText bytes.Buffer

	// Iterate over each rank in the compressed text
	for _, rank := range compressedText {
		// Retrieve the character corresponding to the rank
		character := dictionary[rank]

		// Append the character to the decoded text
		decodedText.WriteByte(character)

		// Move the character to the front of the dictionary
		copy(dictionary[1:], dictionary[:rank])
		dictionary[0] = character
	}

	return decodedText.String()
}
