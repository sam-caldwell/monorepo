package huffman

import (
	"github.com/sam-caldwell/monorepo/go/compress/bitreader"
	"math"
)

// Decode reads bits from the given bitReader and navigates the tree until a
// symbol is found.
func (t *huffmanTree[CodeType, SymbolType]) Decode(br *bitreader.BitReader) (v CodeType) {
	nodeIndex := CodeType(0) // node 0 is the root of the tree.

	for {
		var bit CodeType
		node := &t.nodes[nodeIndex]

		if err := br.Error(); err == nil {
			if b := br.ReadBit(); b {
				bit = 1
			} else {
				bit = 0
			}
		} else {
			break //Keep reading until we hit an error or EOF.
		}

		if l, r := node.left, node.right; bit == 1 {
			nodeIndex = l
		} else {
			nodeIndex = r
		}

		if nodeIndex == CodeType(math.MaxUint64) {
			// We found a leaf. Use the value of bit to decide
			// whether is a left or a right value.
			l, r := node.leftValue, node.rightValue
			if bit == 1 {
				v = l
			} else {
				v = r
			}
			return
		}
	}

	return

}
