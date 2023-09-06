package convert

// DecodeHexUint16 - Decode a hexadecimal 2-byte slice into a 16-bit unsigned-integer
func DecodeHexUint16(value *uint16, src []byte) (ok bool) {
	var b1, b2 byte

	ok = DecodeHexByte(&b1, src[0], src[1]) &&
		DecodeHexByte(&b2, src[2], src[3])

	*value = (uint16(b1) << 8) | uint16(b2)

	return
}
