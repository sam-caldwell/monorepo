package huffman

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	t.Skip("debugging")
	t.Run("Edge Case 1: Empty Input", func(t *testing.T) {
		var input []byte
		expectedCodes := make(CodeMap)
		codes := Encode(input)
		if !reflect.DeepEqual(codes, expectedCodes) {
			t.Errorf("Encode(%v) = %v, expected %v",
				input, codes, expectedCodes)
		}
	})
	t.Run("Edge Case 2: Input with one symbol", func(t *testing.T) {
		input := []byte{'a'}
		expectedCodes := CodeMap{'a': []byte{}}
		codes := Encode(input)
		if !reflect.DeepEqual(codes, expectedCodes) {
			t.Errorf("Encode(%v)\n"+
				" actual:  %v\n"+
				" expected %v\n",
				input, codes, expectedCodes)
		}
	})
	t.Run("Normal Case 1: Input with multiple symbols", func(t *testing.T) {
		input := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'a', 'b', 'c', 'd', 'e', 'f'}
		expectedCodes := CodeMap{
			'a': []byte{1, 1, 0},
			'b': []byte{1, 0, 0},
			'c': []byte{0, 1},
			'd': []byte{0, 0},
			'e': []byte{1, 0, 1},
			'f': []byte{1, 1, 1},
		}
		codes := Encode(input)
		if !reflect.DeepEqual(codes, expectedCodes) {
			t.Errorf("Encode(%v)\n"+
				" actual:  %v\n"+
				" expected %v\n",
				input, codes, expectedCodes)
		}
	})
	//t.Run("Normal Case 2: Input with varying frequencies", func(t *testing.T) {
	//input: []byte{'a', 'b', 'a', 'c', 'd', 'a', 'b', 'e', 'f', 'f', 'a', 'b', 'd'},
	//    expectedCodes: CodeMap{
	//        'a': []byte("0"),
	//        'b': []byte("10"),
	//        'c': []byte("1101"),
	//        'd': []byte("111"),
	//        'e': []byte("11001"),
	//        'f': []byte("11000"),
	//    },
	//		codes := Encode(test.input)
	//		if !reflect.DeepEqual(codes, test.expectedCodes) {
	//			t.Errorf("Encode(%v) = %v, expected %v", test.input, codes, test.expectedCodes)
	//		}
	//})
}
