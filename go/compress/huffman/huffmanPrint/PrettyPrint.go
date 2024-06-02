package huffmanPrint

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
)

// PrettyPrintTree - Use ASCII Art to display the Huffman tree.
//
//	(c) 2023 Sam Caldwell.  MIT License
func PrettyPrintTree(title string, root *node) {
	ansi.Blue().LF().
		Line(words.Hyphen, 40).
		Println(title)
	printNode(root, words.Space, true)
	ansi.Blue().
		Line(words.Hyphen, 40).LF().
		Reset()
}
