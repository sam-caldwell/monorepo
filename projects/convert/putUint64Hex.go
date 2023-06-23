package convert

// putUint64Hex - write a 64-bit value into a hexadecimal byte array
func putUint64Hex(b []byte, v uint64) {
	/*
	 * To convert an Uint64 to a byte slice in hexadecimal
	 * we perform a series of mask-and-shift operations.
	 */
	b[0] = hexTable[byte(v>>56)>>4]
	b[1] = hexTable[byte(v>>56)&0x0f]

	b[2] = hexTable[byte(v>>48)>>4]
	b[3] = hexTable[byte(v>>48)&0x0f]

	b[4] = hexTable[byte(v>>40)>>4]
	b[5] = hexTable[byte(v>>40)&0x0f]

	b[6] = hexTable[byte(v>>32)>>4]
	b[7] = hexTable[byte(v>>32)&0x0f]

	b[8] = hexTable[byte(v>>24)>>4]
	b[9] = hexTable[byte(v>>24)&0x0f]

	b[10] = hexTable[byte(v>>16)>>4]
	b[11] = hexTable[byte(v>>16)&0x0f]

	b[12] = hexTable[byte(v>>8)>>4]
	b[13] = hexTable[byte(v>>8)&0x0f]

	b[14] = hexTable[byte(v)>>4]
	b[15] = hexTable[byte(v)&0x0f]
}
