package convert

// DecodeHexUint32 - decode a 32-bit hexadecimal byte string as uint32
//
//	(c) 2023 Sam Caldwell.  MIT License
func DecodeHexUint32(value *uint32, src []byte) (ok bool) {
	/*
	 * decodeHexUint32()
	 *      - Assumptions:
	 *           - Given 8 bytes of hexadecimal
	 *           - Each byte represents a single hex digit
	 *           - thus we need two bytes (src[i], src[i+1] to represent a hex number
	 *           - We pass our value by reference (address) so we can change the caller's
	 *             data.
	 *      - Expectations:
	 *           - value will have a numeric 32-bit unsigned value
	 *           - a boolean true says the operation went well.
	 *           - a boolean false says you really should look at your inputs.
	 *           - Unlike when I commented decodeHexBytes64() I decided on wine rather than scotch
	 */
	if len(src) < 8 {
		panic(len(src))
	}
	var b1, b2, b3, b4 byte
	var temp uint64
	ok = DecodeHexByte(&b1, src[0], src[1]) &&
		DecodeHexByte(&b2, src[2], src[3]) &&
		DecodeHexByte(&b3, src[4], src[5]) &&
		DecodeHexByte(&b4, src[6], src[7])

	temp = (uint64(b1) << 24) | (uint64(b2) << 16) | (uint64(b3) << 8) | uint64(b4)
	*value = uint32(temp)
	return
}
