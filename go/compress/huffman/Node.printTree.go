package huffman

import "github.com/sam-caldwell/monorepo/go/ansi"

// printTree is a helper function to recursively print the tree
//
//	(c) 2023 Sam Caldwell.  MIT License
func (node *Node) printTree(prefix string, isTail bool) {
	if node.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		node.right.printTree(newPrefix, false)
	}

	if isTail {
		ansi.Blue().Println(prefix + "└── " + node.String()).Flush().Reset()
	} else {
		ansi.Blue().Println(prefix + "┌── " + node.String()).Flush().Reset()
	}

	if node.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		node.left.printTree(newPrefix, true)
	}
}
