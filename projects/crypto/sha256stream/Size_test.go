package crypto

import (
	"testing"
)

func TestSha256StreamSize(t *testing.T) {
	// Create a new Sha256Stream instance
	stream := Sha256Stream{}

	// Test Case 1: Add 10 bytes to the buffer. 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'
	// Size should not increase for input less than 64 bytes
	bytesToAdd1 := []byte{0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A}
	expectedSize1 := 0
	for _, b := range bytesToAdd1 {
		stream.AddByte(b)
	}
	actualSize1 := stream.Size()
	if actualSize1 != expectedSize1 {
		t.Errorf("Size() method didn't return the correct number of bytes processed (Test Case 1).\n"+
			"Expected size: %d\n"+
			"Actual size: %d", expectedSize1, actualSize1)
	}

	// Test Case 2: Add 23 bytes to the buffer
	bytesToAdd2 := []byte{
		0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64,
		0x20, 0x54, 0x68, 0x72, 0x65, 0x65, 0x21, 0x21, 0x21, 0x21, 0x21, 0x21,
	} // 'Hello World Three!!!!!!'
	expectedSize2 := 0 // Size should not increase for input less than 64 bytes
	for _, b := range bytesToAdd2 {
		stream.AddByte(b)
	}
	actualSize2 := stream.Size()
	if actualSize2 != expectedSize2 {
		t.Errorf("Size() method didn't return the correct number of bytes processed (Test Case 2).\n"+
			"Expected size: %d\n"+
			"Actual size: %d", expectedSize2, actualSize2)
	}

	// Test Case 3: Add 64 bytes to the buffer
	bytesToAdd3 := make([]byte, 64)
	expectedSize3 := len(bytesToAdd3) // Size should increase by 64 for 64 bytes input
	for i := 0; i < len(bytesToAdd3); i++ {
		bytesToAdd3[i] = byte(i)
	}
	for _, b := range bytesToAdd3 {
		stream.AddByte(b)
	}
	actualSize3 := stream.Size()
	if actualSize3 != expectedSize3 {
		t.Errorf("Size() method didn't return the correct number of bytes processed (Test Case 3).\n"+
			"Expected size: %d\n"+
			"Actual size: %d", expectedSize3, actualSize3)
	}

	// Test Case 4: Add 128 bytes to the buffer
	bytesToAdd4 := make([]byte, 128)
	expectedSize4 := len(bytesToAdd4) // Size should increase by 64 for 64 bytes input, and by another 64 for the next 64 bytes input
	for i := 0; i < len(bytesToAdd4); i++ {
		bytesToAdd4[i] = byte(i)
	}
	for _, b := range bytesToAdd4 {
		stream.AddByte(b)
	}
	actualSize4 := stream.Size()
	if actualSize4 != expectedSize4 {
		t.Errorf("Size() method didn't return the correct number of bytes processed (Test Case 4).\n"+
			"Expected size: %d\n"+
			"Actual size: %d", expectedSize4, actualSize4)
	}
}
