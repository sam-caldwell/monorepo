package huffman

import "github.com/sam-caldwell/monorepo/go/ansi"

// printNode - Recursively print tree structure
//
//	(c) 2023 Sam Caldwell.  MIT License
func printNode(n *node, prefix string, isTail bool) {
	if n == nil {
		return
	}

	// Print the current node
	ansi.Blue().
		Printf("%s%s- %c:%d\n",
			prefix, getBranch(isTail),
			n.Character, n.Frequency).
		Reset()

	// Determine the new prefix for the child nodes
	newPrefix := prefix + getNextPrefix(isTail)

	// Print the left and right children
	printNode(n.Left, newPrefix, n.Right == nil)
	printNode(n.Right, newPrefix, true)
}