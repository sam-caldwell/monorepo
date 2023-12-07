package asciitree

import (
	"fmt"
)

// DrawTree recursively draws the tree structure using ASCII extended characters.
func DrawTree(node *Node, depth int, lastChild bool) {
	if node == nil {
		return
	}

	// Print indentation
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}

	// Print branch character
	var branch string
	if depth == 0 {
		branch = " "
	} else if lastChild {
		branch = "└"
	} else {
		branch = "├"
	}
	fmt.Print(branch, " ")

	// Print node key
	fmt.Print(node.Key)

	// Print node value if it is a string
	if str, ok := node.Value.(string); ok {
		fmt.Print(":", str)
	}

	fmt.Println()

	// Draw children recursively with updated depth and last child information
	lastChildFlag := true
	for i, child := range node.Children {
		isLast := i == len(node.Children)-1
		if !isLast {
			lastChildFlag = false
		}
		DrawTree(child, depth+1, lastChildFlag)
	}
}
