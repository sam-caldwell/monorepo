package huffman

import (
	"bytes"
	"testing"
)

func TestDecode(t *testing.T) {
	t.Run("Test: Happy Path", func(t *testing.T) {
		var encodedInput []byte

		// Example input
		input := []byte("this is an example for huffman encoding")

		// Encode the input
		encodedData := Encode(input)

		// Create a byte slice to simulate the encoded input
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
				"Got:      %s",
				input, decodedOutput)
		}
	})
	t.Run("Test: Invalid Input", func(t *testing.T) {
		// Create an invalid encoded input by appending an extra bit
		invalidInput := []byte{0, 1, 0, 1, 1, 1, 0, 0, 0, 1, 1} // Extra '1' at the end

		// Test decoding the invalid input
		_, err := Decode(invalidInput, CodeMap{})
		if err == nil {
			t.Error("Expected error for decoding invalid input, but got none")
		}
	})
}
