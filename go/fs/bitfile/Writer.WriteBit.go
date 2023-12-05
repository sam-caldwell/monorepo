package bitfile

/*
 * bitfile Writer.WriteBit() method
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 */

// WriteBit - Write a given bit to the buffer and when all bits are buffered, push to file.
// Given a value (byte), extract the bit at position (pos) and push to the current byte at
// buffer[bufferPos].
func (o *Writer) WriteBit(pos, inputValue byte) (err error) {
	// readMask decodes the bit from the given value
	var readMask = byte(1 << pos)
	// determine the bit value and store it.
	if bit := readMask&inputValue != 0; bit {
		o.bitPos = pos
		o.setBit()
	} else {
		o.bitPos = pos
		o.clearBit()
	}
	o.bitPos++
	return nil
}
