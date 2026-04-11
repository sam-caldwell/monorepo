package asciitree

// MoveToRoot moves the current node to the root of the tree.
func (node *Node) MoveToRoot() error {
	for node.Parent != nil {
		if err := node.MoveParent(); err != nil {
			return err
		}
	}
	return nil
}
