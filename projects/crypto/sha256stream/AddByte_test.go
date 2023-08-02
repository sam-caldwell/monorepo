package crypto

import (
	"reflect"
	"testing"
)

func TestSha256StreamAddByte(t *testing.T) {
	// Create a new Sha256Stream instance
	stream := Sha256Stream{}

	// Add some bytes to the buffer
	bytesToAdd := []byte{0x61, 0x62, 0x63, 0x64} // 'a', 'b', 'c', 'd'
	for _, b := range bytesToAdd {
		stream.AddByte(b)
	}

	// Calculate the expected buffer state after adding the bytes
	expectedBufferState := [64]byte{0x61, 0x62, 0x63, 0x64}
	expectedByteNdx := int8(len(bytesToAdd))
	expectedBitNdx := int8(0)

	// Compare the actual buffer state with the expected buffer state
	if !reflect.DeepEqual(stream.buffer, expectedBufferState) {
		t.Errorf("AddByte() method didn't set the buffer as expected.\n"+
			"Expected Buffer: %v\n"+
			"Actual Buffer: %v",
			expectedBufferState, stream.buffer)
	}

	// Compare the actual byteNdx with the expected byteNdx
	if stream.byteNdx != expectedByteNdx {
		t.Errorf("AddByte() method didn't set the byteNdx as expected.\n"+
			"Expected byteNdx: %d\n"+
			"Actual byteNdx: %d",
			expectedByteNdx, stream.byteNdx)
	}

	// Compare the actual bitNdx with the expected bitNdx
	if stream.bitNdx != expectedBitNdx {
		t.Errorf("AddByte() method didn't set the bitNdx as expected.\n"+
			"Expected bitNdx: %d\n"+
			"Actual bitNdx: %d",
			expectedBitNdx, stream.bitNdx)
	}
}
