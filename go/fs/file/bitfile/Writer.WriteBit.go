package bitfile

/*
 * bitfile Writer.WriteBit() method
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Write a single bit to the buffer and ultimately to the target file.
 */

// WriteBit - Write a given bit to the buffer and when all bits are buffered, push to file.
// Given a value (byte), extract the bit at position (pos) and push to the current byte at
// buffer[bufferPos].
func (o *Writer) WriteBit(pos, inputValue byte) (err error) {
	// readMask decodes the bit from the given value
	var readMask = byte(1 << pos)
	// determine the bit value and store it.
	if bit := readMask&inputValue != 0; bit {
		//ansi.Blue().Printf("Set bit %d (mask %08b): %08b", pos, readMask, o.buffer[o.bufferPos]).LF().Reset()
		o.bitPos = pos
		o.setBit()
		//ansi.Green().Printf("Set bit %d (mask %08b): %08b", pos, readMask, o.buffer[o.bufferPos]).LF().Reset()
	} else {
		//ansi.Blue().Printf("Clear bit %d (mask %08b): %08b", pos, readMask, o.buffer[o.bufferPos]).LF().Reset()
		o.bitPos = pos
		o.clearBit()
		//ansi.Red().Printf("Clear bit %d (mask %08b): %08b", pos, readMask, o.buffer[o.bufferPos]).LF().Reset()
	}
	o.bitPos++
	return nil
}
