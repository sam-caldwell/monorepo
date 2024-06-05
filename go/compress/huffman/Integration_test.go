package huffman

import (
	"bytes"
	"testing"
)

func TestHuffman_End_To_End(t *testing.T) {
	// Example input
	input := []byte("this is an example for huffman encoding")

	// Encode the input
	encodedData := Encode(input)

	// Create a byte slice to simulate the encoded input
	var encodedInput []byte
	for _, b := range input {
		encodedInput = append(encodedInput, encodedData[b]...)
	}

	// Decode the encoded input
	decodedOutput, err := Decode(encodedInput, encodedData)
	if err != nil {
		t.Errorf("Unexpected error decoding: %v", err)
	}

	// Check if the decoded output matches the original input
	if !bytes.Equal(decodedOutput, input) {
		t.Errorf("Decoded output does not match original input.\n"+
			"Expected: %s\n"+
			"Got: %s", input, decodedOutput)
	}
}
