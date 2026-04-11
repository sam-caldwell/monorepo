package asciitree

import "fmt"

// MoveToChild - Move the current node pointer to the child node using its index.
func (node *Node) MoveToChild(childIndex int) error {
	if childIndex < 0 || childIndex >= len(node.Children) {
		return fmt.Errorf("invalid child index: %d", childIndex)
	}

	// Extract the target child
	targetChild := node.Children[childIndex]

	// Remove current node from parent's children
	for i, child := range node.Parent.Children {
		if child == node {
			node.Parent.Children = append(node.Parent.Children[:i], node.Parent.Children[i+1:]...)
			break
		}
	}

	// Replace target child with current node
	targetChild.Parent.Children[childIndex] = node
	node.Parent = targetChild.Parent

	// Update child's parent and children
	targetChild.Children = append(targetChild.Children, node.Children...)
	for _, child := range node.Children {
		child.Parent = targetChild
	}
	node.Children = nil

	return nil
}
