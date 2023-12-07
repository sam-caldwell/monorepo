package asciitree

// findNode recursively searches for the target node.
func findNode(node, targetNode *Node, found *bool, parent **Node) {
	if node == nil {
		return
	}

	if node == targetNode {
		*found = true
		*parent = node.Parent
		return
	}

	for _, child := range node.Children {
		findNode(child, targetNode, found, parent)
		if *found {
			break
		}
	}
}
