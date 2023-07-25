//go:build arm64
// +build arm64

package hijack

import (
	"encoding/binary"
	endianness "github.com/sam-caldwell/go/v2/projects/convert/Endianness"
	"reflect"
	"testing"
	"unsafe"
)

func TestAssemblyJmpToFunction(t *testing.T) {
	const (
		ptr32bit = 4 // 32bits / 8 bits per byte = 4
		ptr64bit = 8 // 64bits / 8 bits per byte = 8
	)

	// Create our fake target and its address (destination)
	targetFunc := func() {}
	// Determine the address of our destination address (32/64-bit)
	destination := reflect.ValueOf(targetFunc).Pointer()
	// Get the size of our destination
	AddressSize := unsafe.Sizeof(destination)
	// Allocate memory so we can convert destination to a []byte slice
	byteAddress := make([]byte, AddressSize)
	// Determine if we are dealing with 32 or 64 bits?
	switch AddressSize {
	case ptr32bit: //We are dealing with 32-bits...cool
		if endianness.LittleEndian {
			binary.LittleEndian.PutUint32(byteAddress, uint32(destination))
		} else {
			binary.BigEndian.PutUint32(byteAddress, uint32(destination))
		}
	case ptr64bit: //We are dealing with 64-bits...how modern of ya?
		if endianness.LittleEndian {
			binary.LittleEndian.PutUint64(byteAddress, uint64(destination))
		} else {
			binary.BigEndian.PutUint64(byteAddress, uint64(destination))
		}
	default:
		// We're in the future!  I told you not to take the delorean past 88mph...but did you listen? No...
		t.Fatalf("TestAssemblyJmpToFunction does not support size %d", AddressSize)
	}

	expectedInstructions := make([]byte, AddressSize+3)
	//                                       0x00  0x01  0x02  0x03                  | 04 byte - instruction: ldr x17
	copy(expectedInstructions[0x00:], []byte{0xF8, 0x1F, 0xFF, 0x10}) //             |
	//                                       0x04 0x05 0x06 0x07 0x08 0x09 0x0A 0x0B | 08 byte - 64-bit address
	copy(expectedInstructions[0x04:], byteAddress) //                                |
	//                                       0x0C  0x0D  0x0E  0x0F                  | 16 byte - instruction: ldr x17
	copy(expectedInstructions[12:], []byte{0xE0, 0x03, 0x3F, 0xD6}) //               | --------------
	//                                                                               | 16-bytes total

	//instructions := AssemblyJmpToFunction(destination)

	// Verify that the generated instructions match the expected instructions
	//if !reflect.DeepEqual(instructions, expectedInstructions) {
	//	t.Errorf("Generated instructions do not match expected instructions")
	//}

	// Test case: Destination is 0 (null address)
	destination = uintptr(0)
	//expectedInstructions = []byte{
	//	0xF8, 0x1F, 0xFF, 0x10,
	//	0x00, 0x00, 0x00, 0x00,
	//	0xE0, 0x03, 0x3F, 0xD6}
	//instructions = AssemblyJmpToFunction(destination)

	// Verify that the generated instructions match the expected instructions
	//if !reflect.DeepEqual(instructions, expectedInstructions) {
	//	t.Errorf("Generated instructions do not match expected instructions")
	//}
}
