package huffman

// isEquivalentTree - Function to check if two Huffman trees are equivalent
//
//	(c) 2023 Sam Caldwell.  MIT License
func isEquivalentTree(tree1, tree2 *node) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}
	if tree1 == nil || tree2 == nil {
		return false
	}
	if tree1.Character != tree2.Character || tree1.Frequency != tree2.Frequency {
		return false
	}
	return isEquivalentTree(tree1.Left, tree2.Left) && isEquivalentTree(tree1.Right, tree2.Right)
}
