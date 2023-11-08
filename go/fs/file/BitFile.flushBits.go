package file

/*
 * CRSCE bitfile writer
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer
 */

// flushBits - Assume we have reached a byte boundary and flush the current bits (byte) to the buffer and disk
func (o *BitFile) flushBits() (err error) {

	o.bitPos = 0

	o.bufferPos++

	if o.bufferPos >= o.bufferSize {
		err = o.flushBuffer()
	}

	return err

}
