package huffman

// Decode - Given a CodeMap, return the decoded []byte.
//
//	Decodes the input encoded using Huffman encoding and returns
//	the original byte slice and an error if any
//
//	(c) 2023 Sam Caldwell.  MIT License
func Decode(freqs map[byte]int, encodedBits []bool) (out []byte) {
	tree := buildHuffmanTree(freqs)
	node := tree.Root
	for _, bit := range encodedBits {
		if bit {
			node = node.Right
		} else {
			node = node.Left
		}
		if node.Left == nil && node.Right == nil {
			out = append(out, node.Symbol)
			node = tree.Root
		}
	}
	return out
}
