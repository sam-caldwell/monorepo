package huffman

// node - Huffman tree node
//
//	(c) 2023 Sam Caldwell.  MIT License
type node struct {
	Character byte
	Frequency int
	Left      *node
	Right     *node
}
