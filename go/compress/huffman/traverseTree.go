package huffman

// traverseTree - assigns binary codes to each symbol by traversing the Huffman tree
//
//	(c) 2023 Sam Caldwell.  MIT License
//
// traverseTree - assigns binary codes to each symbol by traversing the Huffman tree
func traverseTree(root *Node, prefix []byte, codes CodeMap) {
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		// Convert prefix to a copy to avoid mutating it for subsequent nodes
		prefixCopy := make([]byte, len(prefix))
		copy(prefixCopy, prefix)
		codes[root.symbol] = prefixCopy
		return
	}
	// Append '0' for left child and '1' for right child
	traverseTree(root.left, append(prefix, '0'), codes)
	traverseTree(root.right, append(prefix, '1'), codes)
}
