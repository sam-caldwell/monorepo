package huffman

import "github.com/sam-caldwell/monorepo/go/ansi"

// Encode - Perform huffman encoding
//
//	(c) 2023 Sam Caldwell.  MIT License
func Encode(input []byte) *CodeMap {

	frequencyTable := NewFrequencyTable()

	frequencyTable.Calculate(input)

	root := frequencyTable.BuildHuffmanTree()

	codes := make(CodeMap)

	root.traverseTree([]byte{}, codes)
	ansi.Red().
		Printf("%v\n", codes).
		Printf("%v\n", codes.Bytes()).
		Reset()
	return &codes
}
