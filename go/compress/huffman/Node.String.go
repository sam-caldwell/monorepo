package huffman

import "fmt"

// String returns a string representation of a node
//
//	(c) 2023 Sam Caldwell.  MIT License
func (node *Node) String() string {
	if node.symbol == 0 {
		return fmt.Sprintf("(%d)", node.frequency)
	}
	return fmt.Sprintf("(%c:%d)", node.symbol, node.frequency)
}
