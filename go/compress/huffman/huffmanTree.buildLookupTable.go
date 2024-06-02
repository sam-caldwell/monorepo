package huffman

import "math"

// buildLookupTable - traverses the Huffman tree to build a lookup table for each symbol.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (t *huffmanTree[CodeType, SymbolType]) buildLookupTable() map[CodeType]huffmanCode[CodeType, SymbolType] {
	var buildTable func(nodeIndex SymbolType, code CodeType, codeLen uint8)

	buildTable = func(nodeIndex SymbolType, code CodeType, codeLen uint8) {
		node := &t.nodes[nodeIndex]

		if node.left == CodeType(math.MaxUint64) {
			// Leaf node on the left
			t.lookupTable[node.leftValue] = huffmanCode[CodeType, SymbolType]{code: code, codeLen: codeLen}
		} else {
			// Traverse left
			buildTable(node.left, code<<1|1, codeLen+1)
		}

		if node.right == CodeType(math.MaxUint64) {
			// Leaf node on the right
			t.lookupTable[node.rightValue] = huffmanCode[CodeType, SymbolType]{code: code << 1, codeLen: codeLen}
		} else {
			// Traverse right
			buildTable(node.right, code<<1, codeLen+1)
		}
	}

	// Start building the table from the root node
	buildTable(0, 0, 0)
	return t.lookupTable
}
