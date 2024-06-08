package bzip2mtf

import (
	"bytes"
	"github.com/sam-caldwell/monorepo/go/random"
	"testing"
)

func TestMtfIntegration(t *testing.T) {

	// Define the number of test cases
	numTestCases := 10

	for i := 0; i < numTestCases; i++ {
		// Generate a random byte slice of random length
		input, err := random.GenerateRandomBytes(random.Int(1, 1024))
		if err != nil {
			t.Fatal(err)
		}

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
