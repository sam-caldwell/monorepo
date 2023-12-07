package asciitree

// Node represents a node in the tree.
type Node struct {
	Key      string
	Value    interface{}
	Children []*Node
}
