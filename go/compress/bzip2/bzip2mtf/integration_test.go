package bzip2mtf

import (
    "bytes"
    "math/rand"
    "testing"
)

func TestMtfIntegration(t *testing.T) {

	// Define the number of test cases
	numTestCases := 10

	for i := 0; i < numTestCases; i++ {
		// Generate a random byte slice of random length
		input := generateRandomBytes(rand.Intn(100) + 1)

		// Perform MTF encoding
		compressedText := Mtf(input)

		// Perform iMTF decoding
		decodedText := Imtf(compressedText)

		// Check if the decoded text matches the original input
		if !bytes.Equal([]byte(decodedText), input) {
			t.Errorf("Integration test failed for input: %v\n"+
				"Decoded text: %v\n"+
				"Original input: %v\n",
				input, decodedText, input)
		}
	}
}

// Function to generate a random byte slice of given length
func generateRandomBytes(length int) []byte {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(rand.Intn(256)) // Generate a random byte value
	}
	return bytes
}
