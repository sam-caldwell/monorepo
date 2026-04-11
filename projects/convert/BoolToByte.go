package convert

// BoolToByte - Convert Boolean true/false to a byte value
//
//	(c) 2023 Sam Caldwell.  MIT License
func BoolToByte(b bool) byte {
	if b {
		return 0xFF
	}
	return 0x00
}
