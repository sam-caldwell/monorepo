package asciitree

import "fmt"

// MoveParent moves the current node to its parent.
func (node *Node) MoveParent() error {
	if node.Parent == nil {
		return fmt.Errorf("cannot move parent of root node")
	}

	// Extract current node from parent's children
	for i, child := range node.Parent.Children {
		if child == node {
			node.Parent.Children = append(node.Parent.Children[:i], node.Parent.Children[i+1:]...)
			break
		}
	}

	// Add current node as a child of its parent's parent
	if node.Parent.Parent != nil {
		node.Parent.Parent.Children = append(node.Parent.Parent.Children, node)
		node.Parent = node.Parent.Parent
	} else {
		node.Parent = nil
	}

	return nil
}
