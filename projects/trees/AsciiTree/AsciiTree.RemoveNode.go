package asciitree

import "fmt"

// RemoveNode removes the target node and its children from the tree.
func (node *Node) RemoveNode(targetNode *Node) error {
	if targetNode == nil {
		return fmt.Errorf("cannot remove nil node")
	}

	// Traverse the tree to find the target node
	var found bool
	var parent *Node
	findNode(node, targetNode, &found, &parent)
	if !found {
		return fmt.Errorf("target node not found")
	}

	// Remove target node from parent's children
	for i, child := range parent.Children {
		if child == targetNode {
			parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
			break
		}
	}

	// Recursively remove child nodes of the target node
	for _, child := range targetNode.Children {
		if err := child.RemoveNode(child); err != nil {
			return err
		}
	}

	return nil
}
