package huffman

// huffmanNode is a node in the huffmanTree.
//
// left and right contain indexes into the nodes slice of the tree.
// If left or right is invalidNodeValue then the child is a left node and its
// value is in leftValue/rightValue.
//
// The symbols are uint16s because bzip2 encodes not only MTF indexes in the
// tree, but also two magic values for run-length encoding and an EOF symbol.
// Thus, there are more than 256 possible symbols.
//
//	(c) 2023 Sam Caldwell.  MIT License
type huffmanNode[CodeType IndexTypes] struct {
	left, right           CodeType
	leftValue, rightValue CodeType
}
