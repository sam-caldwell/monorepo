package huffman

// huffmanTree - A bitwise binary tree
//
//	(c) 2023 Sam Caldwell.  MIT License
type huffmanTree[CodeType, SymbolType IndexTypes] struct {
	//nodes - a list of all non-leaf nodes in the tree, where nodes[0] is tree root.
	nodes []huffmanNode[CodeType]
	//nextNode - index of the next element of nodes to use when building the tree.
	nextNode CodeType

	//lookupTable - provides a quick look up table for each symbol in the tree.
	lookupTable map[CodeType]huffmanCode[CodeType, SymbolType]
}
