package huffman

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
)

// PrettyPrint method to print the Huffman tree using ASCII art
//
//	(c) 2023 Sam Caldwell.  MIT License
func (node *Node) PrettyPrint() {
	ansi.Blue().Println("(root)").Reset()
	printTree(node, "  ", true)
}

// printTree is a helper function to recursively print the tree
//
//	(c) 2023 Sam Caldwell.  MIT License
func printTree(node *Node, prefix string, isTail bool) {
	if node.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		printTree(node.right, newPrefix, false)
	}

	if isTail {
		ansi.Blue().Println(prefix + "└── " + nodeString(node)).Reset()
	} else {
		ansi.Blue().Println(prefix + "┌── " + nodeString(node)).Reset()
	}

	if node.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		printTree(node.left, newPrefix, true)
	}
}

// nodeString returns a string representation of a node
//
//	(c) 2023 Sam Caldwell.  MIT License
func nodeString(node *Node) string {
	if node.symbol == 0 {
		return fmt.Sprintf("(%d)", node.frequency)
	}
	return fmt.Sprintf("(%c:%d)", node.symbol, node.frequency)
}
