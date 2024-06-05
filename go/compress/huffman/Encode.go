package huffman

// Encode - Perform huffman encoding
//
//	(c) 2023 Sam Caldwell.  MIT License
func Encode(input []byte) CodeMap {

	frequencyTable := NewFrequencyTable()

	frequencyTable.Calculate(input)
	frequencyTable.Dump()

	root := frequencyTable.BuildHuffmanTree()

	codes := make(CodeMap)

	//on empty input, root is not nil...
	root.PrettyPrint()

	root.traverseTree([]byte{}, codes)

	return codes
}
