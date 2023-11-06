package bitfile

/*
 * NewBlock()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Block constructor function, allocates memory for a Block object.
 */

// NewBlock creates a new Block with a buffer of specified size
func NewBlock(sz int) *Block {
	if sz <= 0 {
		sz = 0
	}
	return &Block{
		buffer: make([]byte, sz),
	}
}
