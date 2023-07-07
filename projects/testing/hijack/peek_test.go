package hijack

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestPeek(t *testing.T) {
	// Test case: Peek into a memory address with valid data
	data := []byte{0x12, 0x34, 0x56, 0x78}
	addr := uintptr(unsafe.Pointer(&data[0]))
	length := len(data)
	result := peek(addr, length)

	// Verify that the peeked data matches the original data
	if !reflect.DeepEqual(result, data) {
		t.Errorf("Peeked data does not match the original data")
	}

	// Test case: Peek into a memory address with zero length
	var emptyData []byte
	emptyAddr := uintptr(unsafe.Pointer(&emptyData))
	emptyLength := len(emptyData)
	emptyResult := peek(emptyAddr, emptyLength)

	// Verify that peeking into an empty memory address returns an empty slice
	if len(emptyResult) != 0 {
		t.Errorf("Peeked data is not empty for an empty memory address")
	}
}
