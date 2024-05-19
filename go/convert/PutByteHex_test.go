package convert

import (
	"reflect"
	"testing"
)

func TestPutByteHex(t *testing.T) {
	byteSource := []byte{0x12, 0x34, 0xAB, 0xCD}
	expectedHexOutput := []byte{'1', '2', '3', '4', 'A', 'B', 'C', 'D'}

	hexOutput := make([]byte, len(byteSource)*2)

	PutByteHex(hexOutput, byteSource)

	if !reflect.DeepEqual(hexOutput, expectedHexOutput) {
		t.Fatalf("Unexpected hex output. Expected: %v, Got: %v", expectedHexOutput, hexOutput)
	}
}
