package huffman

// getBranch - returns a branch character based on tree location
//
//	(c) 2023 Sam Caldwell.  MIT License
func getBranch(isTail bool) string {
	if isTail {
		return "└"
	}
	return "├"
}
