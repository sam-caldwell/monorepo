package huffman

// getNextPrefix returns prefix to be used for the next level of the tree
//
//	(c) 2023 Sam Caldwell.  MIT License
func getNextPrefix(isTail bool) string {
	if isTail {
		return "    "
	}
	return "│   "
}
