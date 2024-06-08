package huffman

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
)

// PrettyPrint method to print the Huffman tree using ASCII art
//
//	(c) 2023 Sam Caldwell.  MIT License
func (node *Node) PrettyPrint() {
	ansi.Blue().
		Line(words.Hyphen, 40).
		Println("tree...").
		Flush().Reset()
	if node == nil {
		ansi.Blue().Println("<nil>").Flush().Reset()
		return
	}
	ansi.Blue().Println("(root)").Flush().Reset()
	node.printTree("  ", true)
	ansi.Blue().
		Line(words.Hyphen, 40).
		Flush().Reset()
}
