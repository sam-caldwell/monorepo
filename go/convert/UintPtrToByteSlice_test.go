package convert

import (
	"encoding/binary"
	"reflect"
	"testing"
	"unsafe"
)

func TestUintPtrToByteSlice(t *testing.T) {
	/*
	 * Test plan...
	 *    For our given architecture, uintptr may be varying sizes
	 *    so we need to test 0 to max and some points between.
	 */

	// Determine the size of uintptr on this system
	sizeOfUintPtr := func() uintptr {
		var n = 42
		return unsafe.Sizeof(n)
	}()

	test32bit := func(ptr32 uintptr) {
		// Test uintptr with size 4 (32-bit)
		expected32 := make([]byte, 4)
		binary.LittleEndian.PutUint32(expected32, uint32(ptr32))
		result32 := UintPtrToByteSlice(ptr32)
		if !reflect.DeepEqual(result32, expected32) {
			t.Fatalf("UintPtrToByteSlice failed: expected %v, got %v", expected32, result32)
		}
	}
	test64bit := func(ptr64 uintptr) {
		// Test uintptr with size 8 (64-bit)
		expected64 := make([]byte, 8)
		binary.LittleEndian.PutUint64(expected64, uint64(ptr64))
		result64 := UintPtrToByteSlice(ptr64)
		if !reflect.DeepEqual(result64, expected64) {
			t.Fatalf("UintPtrToByteSlice failed: expected %v, got %v", expected64, result64)
		}
	}

	// Based on the size of uintptr, execute a maxtest
	switch size := sizeOfUintPtr; size {
	case 4:
		test32bit(0)
		test32bit(4294967295)
		test32bit((4294967295 - 1) / 2)
		test32bit(4294967295)
	case 8:
		test64bit(0)
		test64bit((18446744073709551615 - 1) / 2)
		test64bit(18446744073709551615)
	default:
		t.Fatalf("UintPtrToByteSlice cannot be tested for size: %d", size)
	}
}
