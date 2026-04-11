package asciitree

// Node represents a node in the tree.
type Node struct {
	Key      string
	Value    any
	Children []*Node
	Parent   *Node
}
