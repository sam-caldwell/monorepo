package huffman

// traverseTree - assigns binary codes to each symbol by traversing the Huffman tree
//
//	(c) 2023 Sam Caldwell.  MIT License
//
// traverseTree - assigns binary codes to each symbol by traversing the Huffman tree
func (node *Node) traverseTree(prefix []byte, codes CodeMap) {
	//ansi.Green().Printf("codes: %v\n", codes).Reset()
	if node == nil {
		//ansi.Red().Println("debug node nil").Reset()
		return
	}
	if node.frequency == 0 {
		//ansi.Red().Println("frequency is zero").Reset()
		return
	}
	if node.left == nil && node.right == nil {
		// Convert prefix to a copy to avoid mutating it for subsequent nodes
		prefixCopy := make([]byte, len(prefix))
		copy(prefixCopy, prefix)
		codes[node.symbol] = prefixCopy
		return
	}
	// Append '0' for left child and '1' for right child
	node.left.traverseTree(append(prefix, 0), codes)
	node.right.traverseTree(append(prefix, 1), codes)
}
