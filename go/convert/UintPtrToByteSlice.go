package convert

import (
	"encoding/binary"
	"fmt"
	endianness "github.com/sam-caldwell/monorepo/go/convert/Endianness"
	"unsafe"
)

// UintPtrToByteSlice - Convert a uintptr of arbitrary size to a []byte slice
//
//	     This file creates a function which will convert an uintptr address to a []byte slice.
//
//	     In golang uintptr could vary in size depending on the system.  It could be a 32-bit
//	     system with 32-bit addresses or a 64-bit system with 64-bit addresses.  We cannot
//	     make an assumption here.  We need to determine the appropriate size then both
//	     create our final byte slice and encode 'destination' (uintptr) into that resulting
//	     byte slice.  We also need to future-proof for the day when 128-bit CPUs are a thing
//
//	     How big is destination?  ("Size does matter...ask a programmer" --Monica <redacted>)
//
//		    (c) 2023 Sam Caldwell.  MIT License
func UintPtrToByteSlice(addr uintptr) (result []byte) {
	/*
	 * In golang uintptr could vary in size depending on the system.  It could be a 32-bit
	 * system with 32-bit addresses or a 64-bit system with 64-bit addresses.  We cannot
	 * make an assumption here.  We need to determine the appropriate size then both
	 * create our final byte slice and encode 'destination' (uintptr) into that resulting
	 * byte slice.
	 *
	 * How big is destination?  ("Size does matter...ask a programmer" --Monica <redacted>)
	 */
	const (
		uintptr32Sz = 4 // 8 bits x 4 bytes = 32 bits
		uintptr64Sz = 8 // 8 bits x 8 bytes = 64 bits
	)

	size := unsafe.Sizeof(addr)

	result = make([]byte, size)

	switch size {

	case uintptr32Sz:
		if endianness.LittleEndian {
			binary.LittleEndian.PutUint32(result, uint32(addr))
		} else {
			binary.BigEndian.PutUint32(result, uint32(addr))
		}

	case uintptr64Sz:
		if endianness.LittleEndian {
			binary.LittleEndian.PutUint64(result, uint64(addr))
		} else {
			binary.BigEndian.PutUint64(result, uint64(addr))
		}

	default:
		panic(fmt.Sprintf("assemblyJmpToFunction does not support size %d", size))
	}

	return result
}
