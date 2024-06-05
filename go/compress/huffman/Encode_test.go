package huffman

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	t.Skip("disabled for debugging")
	t.Run("Edge Case 1: Empty Input", func(t *testing.T) {
		var input []byte
		expectedCodes := make(CodeMap)
		codes := Encode(input)
		t.Logf("Encode('%v')\n"+
			"codes    '%v'\n"+
			"expected '%v'",
			input, codes, expectedCodes)
		if !reflect.DeepEqual(codes, expectedCodes) {
			t.Fatalf("Encode(%v)\n"+
				"got      %v\n"+
				"expected %v",
				input, codes, expectedCodes)
		}
	})
	t.Run("Edge Case 2: Input with one symbol", func(t *testing.T) {
		input := []byte{'a'}
		expectedCodes := CodeMap{'a': []byte{}}
		codes := Encode(input)
		if !reflect.DeepEqual(codes, expectedCodes) {
			t.Fatalf("Encode(%v)\n"+
				" actual:  %v\n"+
				" expected %v\n",
				input, codes, expectedCodes)
		}
	})
	t.Run("Normal Case 1: Input with multiple symbols and equal frequency (balanced tree)", func(t *testing.T) {
		input := []byte{'a', 'b', 'c', 'd', 'a', 'b', 'c', 'd'}

		codes := Encode(input)

		var totalLength int
		for _, b := range input {
			totalLength += len(codes[b])
		}
		expectedLength := len(input) * 2 // 8 characters * 2 bits per character

		if totalLength != expectedLength {
			t.Fatalf("Encode('%s' [%v]) length: %d\n"+
				" actual total length: %d\n"+
				" expected total length: %d\n",
				input, input, len(input), totalLength, expectedLength)
		}
	})
	t.Run("Normal Case 1: Input with multiple symbols and equal frequency (unbalanced tree)", func(t *testing.T) {
		input := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'a', 'b', 'c', 'd', 'e', 'f'}

		codes := Encode(input)

		var totalLength int
		for _, symbol := range input {
			totalLength += len(codes[symbol])
		}
		expectedLength := 32 // 12 characters * 3 bits per character

		if totalLength != expectedLength {
			t.Fatalf("Encode('%s' [%v]) length: %d\n"+
				" actual total length:   %d (%v)\n"+
				" expected total length: %d (%v)\n",
				input, input, len(input), totalLength, codes, expectedLength, input)
		}
	})
	t.Run("Normal Case 2: Input with varying frequencies", func(t *testing.T) {
		input := []byte{'a', 'b', 'a', 'c', 'd', 'a', 'b', 'e', 'f', 'f', 'a', 'b', 'd'}
		expectedCodes := CodeMap{
			'a': []byte{1, 1},
			'b': []byte{0, 1},
			'c': []byte{1, 1, 1, 0},
			'd': []byte{0, 0},
			'e': []byte{1, 1, 0, 0, 1},
			'f': []byte{1, 1, 0, 0, 0},
		}
		codes := Encode(input)
		if !reflect.DeepEqual(codes, expectedCodes) {
			t.Fatalf("Encode(%v)\n"+
				"     got %v\n"+
				"expected %v",
				input, codes, expectedCodes)
		}
	})
}
