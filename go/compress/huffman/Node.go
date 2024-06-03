package huffman

// Node represents a leaf node with a symbol and its frequency
//
// We can use this both for the PriorityQueue (heap) and Huffman tree.
// This will mean storing our symbols from the frequency table into the
// PriorityQueue then relinking them (using the left/right pointers) to
// create the binary tree.
//
//	(c) 2023 Sam Caldwell.  MIT License
type Node struct {
	// symbol - an element from the frequency table.
	symbol byte
	// frequency - The number of times the symbol appears in the input.
	frequency uint
	// left - a pointer to the 0-node.
	left *Node
	// right - a pointer to the 1-node
	right *Node
}
