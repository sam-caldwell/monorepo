package bzip2mtf

// Mtf - perform the move-to-front transformation on the input byte slice.
//
// (c) 2023 Sam Caldwell.  MIT License
func Mtf(input []byte) []byte {

	// Make a copy of the common dictionary
	dictionary := make([]byte, 256)
	copy(dictionary, commonDictionary[:])

	var compressedText []byte

	// Iterate over each character in the plain text
	for _, c := range input {
		// Find the rank of the character in the dictionary
		var rank byte
		for i, value := range dictionary {
			if c == value {
				rank = byte(i)
				break
			}
		}

		// Update the compressed text with the rank
		compressedText = append(compressedText, rank)

		// Move the character to the front of the dictionary
		copy(dictionary[1:], dictionary[:rank])
		dictionary[0] = c
	}

	return compressedText
}
