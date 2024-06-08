package bzip2mtf

import (
	"testing"
)

func TestInverseMoveToFront_Imtf(t *testing.T) {
	fail := func(curr *testing.T, input, output, expectedOutput string) {
		t.Fatalf("Imtf(%v) [%s]\n"+
			"     Got: %v (%s)\n"+
			"Expected: %v (%s)\n",
			input, curr.Name(), output, output,
			expectedOutput, expectedOutput)
	}

	t.Run("Empty compressed text", func(t *testing.T) {
		input := []byte(nil)
		expectedOutput := "" // Assuming this is the expected output for a sad path
		output := Imtf(input)
		if output != expectedOutput {
			fail(t, string(input), output, expectedOutput)
		}
	})

	t.Run("Decoding abacabad", func(t *testing.T) {
		compressedText := []byte{97, 98, 1, 99, 1, 2, 1, 100}
		expectedOutput := "abacabad"
		output := Imtf(compressedText)
		if output != expectedOutput {
			fail(t, string(compressedText), output, expectedOutput)
		}
	})

	t.Run("Decoding aaaa", func(t *testing.T) {
		compressedText := []byte{97, 0, 0, 0}
		expectedOutput := "aaaa"
		output := Imtf(compressedText)
		if output != expectedOutput {
			fail(t, string(compressedText), output, expectedOutput)
		}
	})

	t.Run("Decoding a repeating sequence", func(t *testing.T) {
		compressedText := []byte{97, 0, 0, 0, 98}
		expectedOutput := "aaaab"
		output := Imtf(compressedText)
		if output != expectedOutput {
			fail(t, string(compressedText), output, expectedOutput)
		}
	})
}
