package bzip2mtf

import (
	"bytes"
	"strings"
	"testing"
)

func TestMoveToFront_Mtf(t *testing.T) {
	fail := func(curr *testing.T, input, output, expectedOutput []byte) {
		t.Fatalf("Mtf(%v) [%s]\n"+
			"     Got: %v (%s)\n"+
			"Expected: %v (%s)\n",
			input, curr.Name(), output, output,
			expectedOutput, expectedOutput)
	}

	t.Run("nil input", func(t *testing.T) {
		input := []byte(nil)
		expectedOutput := []byte("") // Assuming this is the expected output for a sad path
		output := Mtf(input)
		if !bytes.Equal(output, expectedOutput) {
			fail(t, input, output, expectedOutput)
		}
	})
	t.Run("Empty string", func(t *testing.T) {
		input := []byte("")
		expectedOutput := []byte("") // Assuming this is the expected output for a sad path
		output := Mtf(input)
		if !bytes.Equal(output, expectedOutput) {
			fail(t, input, output, expectedOutput)
		}
	})
	t.Run("Non-Repeating", func(t *testing.T) {
		input := []byte("abcdef")
		expectedOutput := []byte("abcdef") // Assuming this is the expected output for a sad path
		output := Mtf(input)
		if !bytes.Equal(output, expectedOutput) {
			fail(t, input, output, expectedOutput)
		}
	})
	t.Run("Simple Repeating Sequence", func(t *testing.T) {
		input := []byte("abacabad")
		expectedOutput := []byte{97, 98, 1, 99, 1, 2, 1, 100}
		output := Mtf(input)
		if !bytes.Equal(output, expectedOutput) {
			fail(t, input, output, expectedOutput)
		}
	})
	t.Run("Repeating 4A's", func(t *testing.T) {
		input := []byte("aaaa")
		expectedOutput := []byte{97, 0, 0, 0}
		output := Mtf(input)
		if !bytes.Equal(output, expectedOutput) {
			fail(t, input, output, expectedOutput)
		}
	})
	t.Run("Repeating 10A's", func(t *testing.T) {
		input := []byte(strings.Repeat("a", 10))
		expectedOutput := make([]byte, 10)
		expectedOutput[0] = 97
		output := Mtf(input)
		if !bytes.Equal(output, expectedOutput) {
			fail(t, input, output, expectedOutput)
		}
	})
	t.Run("Repeating 4A's and a b", func(t *testing.T) {
		input := []byte("aaaab")
		expectedOutput := []byte{97, 0, 0, 0, 98}
		output := Mtf(input)
		if !bytes.Equal(output, expectedOutput) {
			fail(t, input, output, expectedOutput)
		}
	})
	t.Run("Repeating aaaabbb", func(t *testing.T) {
		input := []byte("aaaabbb")
		expectedOutput := []byte{97, 0, 0, 0, 98, 0, 0}
		output := Mtf(input)
		if !bytes.Equal(output, expectedOutput) {
			fail(t, input, output, expectedOutput)
		}
	})
	t.Run("Repeating ababababababab", func(t *testing.T) {
		input := []byte("ababababababab")
		expectedOutput := []byte{97, 98, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		output := Mtf(input)
		if !bytes.Equal(output, expectedOutput) {
			fail(t, input, output, expectedOutput)
		}
	})

}
