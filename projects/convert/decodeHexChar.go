package convert

// decodeHexChar - decode a single byte representing a hexadecimal digit and return the numeric equivalent as a byte
func decodeHexChar(value *byte, c byte) (ok bool) {
	if value == nil {
		return false
	}
	switch {
	case '0' <= c && c <= '9':
		*value = c - '0'
		return true
	case 'a' <= c && c <= 'f':
		*value = c - 'a' + 10
		return true
	case 'A' <= c && c <= 'F':
		*value = c - 'A' + 10
		ok = true
	default:
		ok = false
	}
	return
}
