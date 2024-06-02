package huffman

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"sort"
)

// newHuffmanTree -  builds a Huffman tree from a slice
//
// Given a slice containing the code lengths of each symbol.
// The maximum code length is determined by IndexTypes (e.g. 32 bits).
//
//	(c) 2023 Sam Caldwell.  MIT License
func newHuffmanTree[CodeType, SymbolType IndexTypes](lengths []uint8) (huffmanTree[CodeType, SymbolType], error) {
	var t huffmanTree[CodeType, SymbolType]

	if len(lengths) < 2 {
		return t, fmt.Errorf(errors.NotEnoughSymbols)
	}

	// First we sort the code length assignments by ascending code length,
	// using the symbol value to break ties.
	pairs := make([]huffmanSymbolLengthPair, len(lengths))
	for i, length := range lengths {
		pairs[i].value = uint16(i)
		pairs[i].length = length
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].length < pairs[j].length {
			return true
		}
		if pairs[i].length > pairs[j].length {
			return false
		}
		if pairs[i].value < pairs[j].value {
			return true
		}
		return false
	})

	// Now we assign codes to the symbols, starting with the longest code.
	// We keep the codes packed into T, at the most-significant end.
	// So branches are taken from the MSB downwards. This makes it easy to
	// sort them later.
	code := CodeType(0)
	length := uint8(32)

	codes := make([]huffmanCode[CodeType, SymbolType], len(lengths))
	for i := len(pairs) - 1; i >= 0; i-- {
		if length > pairs[i].length {
			length = pairs[i].length
		}
		codes[i].code = code
		codes[i].codeLen = length
		codes[i].symbol = pairs[i].value
		// We need to 'increment' the code, which means treating |code|
		// like a |length| bit number.
		code += 1 << (32 - length)
	}

	// Now we can sort by the code so that the left half of each branch are
	// grouped together, recursively.
	sort.Slice(codes, func(i, j int) bool {
		return codes[i].code < codes[j].code
	})

	t.nodes = make([]huffmanNode[CodeType], len(codes))
	_, err := t.buildHuffmanNode(codes, 0)
	return t, err
}
