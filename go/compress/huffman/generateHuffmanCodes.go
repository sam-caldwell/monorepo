package huffman

// generateHuffmanCodes - recursively generate the output codes map
//
//	(c) 2023 Sam Caldwell.  MIT License
func generateHuffmanCodes(n *node, code string, codes map[byte]string) {
	if n.Left == nil && n.Right == nil {
		codes[n.Character] = code
		return
	}
	if n.Left != nil {
		generateHuffmanCodes(n.Left, code+"0", codes)
	}
	if n.Right != nil {
		generateHuffmanCodes(n.Right, code+"1", codes)
	}
}
