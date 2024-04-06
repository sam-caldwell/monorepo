package simpleArgs

func GetOptionUint16Value(name string, required bool) (value uint16, err error) {
	n, err := GetOptionIntValue(name, required)
	return uint16(n), err
}
