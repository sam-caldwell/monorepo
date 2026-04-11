package convert

import "encoding/binary"

// ByteSliceToUint64WithLittleEndian - Convert 8-byte slice to uint64
//
//	 This function converts it to uint64.
//	 This assumes LittleEndian.
//
//		(c) 2023 Sam Caldwell.  MIT License
func ByteSliceToUint64WithLittleEndian(b []byte) uint64 {
	if len(b) != 8 {
		panic("byte slice must be exactly 8 bytes long")
	}
	return binary.LittleEndian.Uint64(b)
}
