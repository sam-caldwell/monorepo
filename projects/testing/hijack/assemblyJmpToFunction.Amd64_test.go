//go:build amd64
// +build amd64

package hijack

import (
	"encoding/binary"
	"reflect"
	"testing"
	"unsafe"
)

func TestAssemblyJmpToFunction(t *testing.T) {
	const MOV = 0xBA // Assembly instruction MOV
	const JMP = 0xFF // Assembly instruction JMP
	const RDX = 0x22 // CPU RDX Register
	// Create our fake target and its address (destination)
	targetFunc := func() {}
	// Determine the address of our destination address (32/64-bit)
	destination := reflect.ValueOf(targetFunc).Pointer()
	// Get the size of our destination
	sz := unsafe.Sizeof(destination)
	// Convert destination to a []byte slice
	byteAddress := make([]byte, sz)
	switch sz {
	case 4:
		binary.LittleEndian.PutUint32(byteAddress, uint32(destination))
	case 8:
		binary.LittleEndian.PutUint64(byteAddress, uint64(destination))
	default:
		t.Fatalf("TestAssemblyJmpToFunction does not support size %d", sz)
	}
	expectedInstructions := make([]byte, sz+3)
	expectedInstructions[0] = MOV
	copy(expectedInstructions[1:], byteAddress)
	expectedInstructions[sz+1] = JMP
	expectedInstructions[sz+2] = RDX

	actualInstructions := assemblyJmpToFunction(destination)

	// Verify that the generated instructions match the expected instructions
	if !reflect.DeepEqual(actualInstructions, expectedInstructions) {
		t.Errorf("Generated instructions do not match expected instructions.\n"+
			"         destination: %0x\n"+
			"expectedInstructions: %0x\n"+
			"  actualInstructions: %0x\n",
			destination, expectedInstructions, actualInstructions)
	}

	// Test case: Destination is 0 (null address)
	//destination = uintptr(0)
	//expectedInstructions = []byte{0xBA, 0x00, 0x00, 0x00, 0x00, 0xFF, 0x22}
	//instructions = assemblyJmpToFunction(destination)

	// Verify that the generated instructions match the expected instructions
	//if !reflect.DeepEqual(instructions, expectedInstructions) {
	//	t.Errorf("Generated instructions do not match expected instructions")
	//}
}
