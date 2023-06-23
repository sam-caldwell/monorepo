package convert

// PutByteHex - store a byte array as a hex byte array
func PutByteHex(hexOutput, byteSource *[]byte) {

	/*
	 * To express a byte in hexadecimal, we need two bytes.
	 *   -We take the numeric byte value (0x00-0xFF)
	 *   -Then we calculate the two Hex values representing each hex digit.
	 */
	for i := 0; i < len(*byteSource); i++ {

		// Hex Digit 1 calculated and stored as a digit
		(*hexOutput)[i*2] = hexTable[(*byteSource)[i]>>4]

		(*hexOutput)[i*2+1] = hexTable[(*byteSource)[i]&0x0f]

	}

}
