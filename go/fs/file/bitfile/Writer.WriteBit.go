package bitfile

/*
 * CRSCE bitfile writer
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer
 */

// WriteBit - Write file bits to a target file in 4K chunks.
func (o *BitFile) WriteBit(pos, value byte) (err error) {
	mask := byte(1 << pos)
	if bit := (value & mask) != 0; bit {
		o.buffer[o.bufferPos] = o.buffer[o.bufferPos] | (1 << pos)
	}

	return nil

}
