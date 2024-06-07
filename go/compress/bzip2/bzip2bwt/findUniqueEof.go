package bzip2bwt

// findUniqueEOF scans the input and returns a unique EOF character that does not exist in the input
//
//	(c) 2023 Sam Caldwell.  MIT License
func findUniqueEOF(input []byte) byte {
	// Define a set to track characters in the input
	charSet := make(map[byte]bool)

	// Scan the input and add characters to the set
	for _, char := range input {
		charSet[char] = true
	}

	// Check for a unique EOF character that does not exist in the input
	for eof := byte(0x04); ; eof++ {
		if _, exists := charSet[eof]; !exists {
			return eof
		}
	}
}
