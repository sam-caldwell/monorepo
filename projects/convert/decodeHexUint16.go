package convert

// decodeHexUint16 - Decode a hexadecimal 2-byte slice into a 16-bit unsigned-integer
func decodeHexUint16(value *uint16, src []byte) (ok bool) {
	var b1, b2 byte

	ok = decodeHexByte(&b1, src[0], src[1]) &&
		decodeHexByte(&b2, src[2], src[3])

	*value = (uint16(b1) << 8) | uint16(b2)

	return
}
