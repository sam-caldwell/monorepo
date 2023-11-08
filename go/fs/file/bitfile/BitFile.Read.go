package file

/*
 * CRSCE BitFile.Read()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Read a file bit by bit (loading a buffer as needed)
 */

// Read - Read file o.file 4K at a time and return the buffer one bit at a time each call.
func (o *BitFile) Read() (bit bool, err error) {
	//log.Printf("Read() start (bit: %v, bitPos: %v, bufferPos: %v, filePos: %v)\n", bit, o.bitPos, o.bufferPos, o.filePos)
	// adjust the bit position to the next bit or get the next byte in the buffer,
	// resetting the bit position to 0.
	if o.bitPos >= 8 {
		o.bitPos = 0
		o.bufferPos++
		//log.Println("bit position reset")
	}

	// if the buffer is exhausted, read more content into the buffer and
	// reset the buffer position to 0.
	if o.bufferPos >= bufferSize {
		o.buffer = make([]byte, bufferSize)
		o.bufferPos = 0
		o.bufferSize, err = o.file.Read(o.buffer)
		//log.Printf("contents. buffer:'%s'", (string)(o.buffer))
		if err != nil {
			//log.Println("...file read failed")
			return false, err //Bail!
		}
	}

	// mask out only the needed bit.  if the result is not 0,
	// then return the set bit.
	v := o.buffer[o.bufferPos]
	mask := byte(128) >> o.bitPos
	bit = (v & mask) != 0
	//log.Printf("Read() done (bit: %5v (%08b|%c), mask:%8b bitPos: %v, bufferPos: %v, filePos: %v)\n",
	//	bit, v, (rune)(v), mask, o.bitPos, o.bufferPos, o.filePos)
	o.bitPos++
	// Return the bit
	return bit, nil
}
