package huffman

// Encode - Perform huffman encoding
//
//	(c) 2023 Sam Caldwell.  MIT License
func Encode(input []byte) CodeMap {

	frequencyTable := NewFrequencyTable()

	frequencyTable.Calculate(input)

	root := frequencyTable.BuildHuffmanTree()

	codes := make(CodeMap)

	root.traverseTree([]byte{}, codes)

	return codes
}
