package bzip2bwt

// rotations - create all rotations of the input string
//
//	(c) 2023 Sam Caldwell.  MIT License
func rotations(s []byte) [][]byte {
	n := len(s)
	if n == 0 {
		return nil
	}
	// Create a slice of byte slices to store all rotations
	table := make([][]byte, n)
	for i := 0; i < n; i++ {
		// Concatenate the rotated slices correctly
		rot := append(s[i:], s[:i]...)
		table[i] = rot
	}
	return table
}
