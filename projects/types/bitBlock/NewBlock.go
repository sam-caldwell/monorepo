package bitBlock

/*
 * NewBlock()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Block package initializer function
 */

// NewBlock - initialize the bitBlock package with a given size buffer
func NewBlock(sz uint) *Block {
	var block Block
	block.lock.Lock()
	defer block.lock.Unlock()

	if sz > 0 {
		block.buffer = make([]byte, sz)
	}
	return &block
}
