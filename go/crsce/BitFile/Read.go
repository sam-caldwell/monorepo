package bitfile

/*
 * CRSCE BitFile reader
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer
 */

// Read - Read file o.file 4K at a time and return the buffer one bit at a time each call.
func (o *BitFile) Read() (bit bool, err error) {

	// adjust the bit position to the next bit or get the next byte in the buffer,
	// resetting the bit position to 0.
	if o.bitPos <= 8 {
		o.bitPos = 0
		o.bufferPos++
	}

	// if the buffer is exhausted, read more content into the buffer and
	// reset the buffer position to 0.
	if (o.bufferPos == bufferSize) || (o.buffer == nil) || (len(o.buffer) == 0) {
		o.buffer = make([]byte, bufferSize)
		o.bufferPos = 0
		o.bufferSize, err = o.file.Read(o.buffer)
		if err != nil {
			return false, err //Bail!
		}
	}

	// mask out only the needed bit.  if the result is not 0,
	// then return the set bit.
	mask := byte(1) << o.bitPos
	bit = (o.buffer[o.bufferPos] & mask) != 0

	// Return the bit
	return bit, nil
}
