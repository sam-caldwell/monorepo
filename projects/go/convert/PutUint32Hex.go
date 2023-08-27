package convert

// PutUint32Hex - write a 32-bit value into a hexidecimal byte array
func PutUint32Hex(b []byte, v uint32) {
	/*
	 * To convert an Uint32 to a byte slice in hexadecimal
	 * we perform a series of mask-and-shift operations.
	 */
	b[0] = hexTable[byte(v>>24)>>4]
	b[1] = hexTable[byte(v>>24)&0x0f]
	b[2] = hexTable[byte(v>>16)>>4]
	b[3] = hexTable[byte(v>>16)&0x0f]
	b[4] = hexTable[byte(v>>8)>>4]
	b[5] = hexTable[byte(v>>8)&0x0f]
	b[6] = hexTable[byte(v)>>4]
	b[7] = hexTable[byte(v)&0x0f]
}
