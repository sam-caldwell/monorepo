//go:build ppc || ppc64 || s390 || s390x
// +build ppc ppc64 s390 s390x

package endianness

const (
	//BigEndian: most significant byte is stored first (at the smallest address).
	//Memory Address:  0   1   2   3
	//Value (32-bit):  MSB  ..  ..  LSB
	//
	//	(c) 2023 Sam Caldwell.  MIT License
	BigEndian = true

	//LittleEndian: least significant byte is stored first (at the smallest address).
	//Memory Address:  0   1   2   3
	//Value (32-bit):  LSB  ..  ..  MSB
	//
	//	(c) 2023 Sam Caldwell.  MIT License
	LittleEndian = false
)
