//go:build amd64
// +build amd64

package hijack

import (
	"testing"
)

func TestAssemblyJmpToFunction(t *testing.T) {
	//const (
	//	ptr32bit = 4 // 32bits / 8 bits per byte = 4
	//	ptr64bit = 8 // 64bits / 8 bits per byte = 8
	//
	//	MOV = 0xBA // Assembly instruction MOV
	//	JMP = 0xFF // Assembly instruction JMP
	//	RDX = 0x22 // CPU RDX Register
	//)
	//
	//// Create our fake target and its address (destination)
	//targetFunc := func() {}
	//// Determine the address of our destination address (32/64-bit)
	//destination := reflect.ValueOf(targetFunc).Pointer()
	//// Get the size of our destination
	//AddressSize := unsafe.Sizeof(destination)
	//// Convert destination to a []byte slice
	//byteAddress := make([]byte, AddressSize)
	//switch AddressSize {
	//case ptr32bit: //We are dealing with 32-bits...cool
	//	if endianness.LittleEndian {
	//		binary.LittleEndian.PutUint32(byteAddress, uint32(destination))
	//	} else {
	//		binary.BigEndian.PutUint32(byteAddress, uint32(destination))
	//	}
	//case ptr64bit: //We are dealing with 64-bits...how modern of ya?
	//	if endianness.LittleEndian {
	//		binary.LittleEndian.PutUint64(byteAddress, uint64(destination))
	//	} else {
	//		binary.BigEndian.PutUint64(byteAddress, uint64(destination))
	//	}
	//default:
	//	// We're in the future!  I told you not to take the delorean past 88mph...but did you listen? No...
	//	t.Fatalf("TestAssemblyJmpToFunction does not support size %d", AddressSize)
	//}
	//
	//expectedInstructions := make([]byte, AddressSize+3)
	//
	//expectedInstructions[0] = MOV
	//
	//copy(expectedInstructions[1:], byteAddress)
	//
	//expectedInstructions[AddressSize+1] = JMP
	//
	//expectedInstructions[AddressSize+2] = RDX
	//
	//actualInstructions := AssemblyJmpToFunction(destination)
	//
	//// Verify that the generated instructions match the expected instructions
	//if !reflect.DeepEqual(actualInstructions, expectedInstructions) {
	//	t.Fatalf("Generated instructions do not match expected instructions.\n"+
	//		"         destination: %0x\n"+
	//		"expectedInstructions: %0x\n"+
	//		"  actualInstructions: %0x\n",
	//		destination, expectedInstructions, actualInstructions)
	//}
}
