package convert

// PutUint16Hex - Write uint16 to a byte array as hex
func PutUint16Hex(b []byte, v uint16) {
	/*
	 * To convert an Uint16 to a byte slice in hexadecimal
	 * we perform a series of mask-and-shift operations.
	 */
	b[0] = hexTable[byte(v>>8)>>4]
	b[1] = hexTable[byte(v>>8)&0x0f]
	b[2] = hexTable[byte(v)>>4]
	b[3] = hexTable[byte(v)&0x0f]
}
