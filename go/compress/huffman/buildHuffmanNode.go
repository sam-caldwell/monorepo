package huffman

import (
	"fmt"
	"math"
)

// buildHuffmanNode takes a slice of sorted huffmanCodes and builds a node in
// the Huffman tree at the given level. It returns the index of the newly
// constructed node.
func (t *huffmanTree[T]) buildHuffmanNode(codes []huffmanCode[T], level uint32) (nodeIndex uint16, err error) {
	test := uint32(1) << (31 - level)

	// We have to search the list of codes to find the divide between the left and right sides.
	firstRightIndex := len(codes)
	for i, code := range codes {
		if code.code&test != 0 {
			firstRightIndex = i
			break
		}
	}

	left := codes[:firstRightIndex]
	right := codes[firstRightIndex:]

	if len(left) == 0 || len(right) == 0 {
		// There is a superfluous level in the Huffman tree indicating
		// a bug in the encoder. However, this bug has been observed in
		// the wild, so we handle it.

		// If this function was called recursively then we know that
		// len(codes) >= 2 because, otherwise, we would have hit the
		// "leaf node" case, below, and not recurred.
		//
		// However, for the initial call it's possible that len(codes)
		// is zero or one. Both cases are invalid because a zero length
		// tree cannot encode anything and a length-1 tree can only
		// encode EOF and so is superfluous. We reject both.
		if len(codes) < 2 {
			return 0, fmt.Errorf("empty Huffman tree")
		}

		// In this case the recursion doesn't always reduce the length
		// of codes so we need to ensure termination via another
		// mechanism.
		if level == 31 {
			// Since len(codes) >= 2 the only way that the values
			// can match at all 32 bits is if they are equal, which
			// is invalid. This ensures that we never enter
			// infinite recursion.
			return 0, fmt.Errorf("equal symbols in Huffman tree")
		}

		if len(left) == 0 {
			return buildHuffmanNode(t, right, level+1)
		}
		return buildHuffmanNode(t, left, level+1)
	}

	nodeIndex = uint16(t.nextNode)
	node := &t.nodes[t.nextNode]
	t.nextNode++

	if len(left) == 1 {
		// leaf node
		node.left = T(math.MaxUint64)
		node.leftValue = left[0].value
	} else {
		node.left, err = buildHuffmanNode(t, left, level+1)
	}

	if err != nil {
		return
	}

	if len(right) == 1 {
		// leaf node
		node.right = T(math.MaxUint64)
		node.rightValue = right[0].value
	} else {
		node.right, err = buildHuffmanNode(t, right, level+1)
	}

	return
}
