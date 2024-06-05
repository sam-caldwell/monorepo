package huffman

import (
	"bytes"
	"fmt"
)

// Decode - Given a CodeMap, return the decoded []byte.
//
//	 Decodes the input encoded using Huffman encoding and returns
//	 the original byte slice and an error if any
//
//		(c) 2023 Sam Caldwell.  MIT License
func Decode(encodedInput []byte, codeMap CodeMap) ([]byte, error) {
	// Create a reverse map from the CodeMap
	reverseMap := make(map[string]byte)
	for symbol, code := range codeMap {
		reverseMap[string(code)] = symbol
	}

	var decodedOutput bytes.Buffer
	var currentCode bytes.Buffer

	// Traverse the encoded input and use the reverseMap to decode it
	for _, bit := range encodedInput {
		currentCode.WriteByte(bit)
		if symbol, exists := reverseMap[currentCode.String()]; exists {
			decodedOutput.WriteByte(symbol)
			currentCode.Reset()
		}
	}

	// If there's leftover bits in currentCode, it means the input was not valid Huffman encoded data
	if currentCode.Len() > 0 {
		return nil, fmt.Errorf("invalid Huffman encoded input")
	}

	return decodedOutput.Bytes(), nil
}
