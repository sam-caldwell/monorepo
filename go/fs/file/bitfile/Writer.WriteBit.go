package bitfile

/*
 * bitfile Writer.WriteBit() method
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Write a single bit to the buffer and ultimately to the target file.
 */

// WriteBit - Write file bits to a target file in 4K chunks.
func (o *Writer) WriteBit(pos, value byte) (err error) {

	mask := byte(1 << pos)

	if bit := (value & mask) != 0; bit {
		o.buffer[o.bufferPos] = o.buffer[o.bufferPos] | (1 << pos)
	}

	return nil

}
