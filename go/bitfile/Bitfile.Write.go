package bitfile

/*
 * CRSCE bitfile writer
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer
 */

// Write - Write file bits to a target file in 4K chunks.
func (o *BitFile) Write(bit bool) (err error) {
	var bitValue byte
	if bit {
		bitValue = 1
	}
	o.buffer[o.bitPos] = bitValue
	if o.bitPos >= 7 {
		return o.flushBits()
	}
	o.bitPos++
	return nil
}
