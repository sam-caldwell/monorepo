package convert

func BoolToByte(b bool) byte {
	if b {
		return 0xFF
	}
	return 0x00
}
