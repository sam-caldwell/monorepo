package asciitree

// AddChildNode - Add a child node to the current tree
func (node *Node) AddChildNode(key string, value interface{}) {
	newNode := &Node{Key: key, Value: value}
	node.Children = append(node.Children, newNode)
	newNode.Parent = node
}
