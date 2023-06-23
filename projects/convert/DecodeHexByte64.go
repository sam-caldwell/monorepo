package convert

// DecodeHexByte64 -  decodes two sets of hexadecimal bytes and returns the resulting byte values
func DecodeHexByte64(src []byte) (value [8]byte, ok bool) {
	if len(src) == 16 {
		/*
		 *   decodeHexByte()
		 *      - Assumptions:
		 *            -Given 16 bytes of information (8-bits)
		 *            -Each byte represents a single hex digit
		 *            -Thus we need two bytes (src[i], src[i+1]) to represent a hex number.
		 *            -We pass our target value into decodeHexByte() by address and expect
		 *             this value to be updated by the function and a boolean result indicating
		 *             success to be returned.
		 *      - Expectations:
		 *            - value will have the numeric hex value (byte)
		 *            - a boolean true will continue execution to the next byte-pair
		 *            - a boolean false will short-circuit our operation.
		 *            - I will be consuming scotch in 30 minutes.
		 */
		ok = DecodeHexByte(&value[0], src[0], src[1]) &&
			DecodeHexByte(&value[1], src[2], src[3]) &&
			DecodeHexByte(&value[2], src[4], src[5]) &&
			DecodeHexByte(&value[3], src[6], src[7]) &&
			DecodeHexByte(&value[4], src[8], src[9]) &&
			DecodeHexByte(&value[5], src[10], src[11]) &&
			DecodeHexByte(&value[6], src[12], src[13]) &&
			DecodeHexByte(&value[7], src[14], src[15])
	}
	return
}
