package convert

// decodeHexByte - Decode a hexadecimal byte
func decodeHexByte(value *byte, c1, c2 byte) (ok bool) {
	/*
	 * A hex number is two characters (0x00 - 0xFF)
	 * thus from a string to []byte this is two bytes
	 *
	 * our inputs are c1 and c2 (hexidecimal characters)
	 * the output is the 8-bit numeric value.
	 */
	var n1, n2 byte
	ok = decodeHexChar(&n1, c1) && decodeHexChar(&n2, c2)
	*value = (n1 << 4) | n2
	return
}
