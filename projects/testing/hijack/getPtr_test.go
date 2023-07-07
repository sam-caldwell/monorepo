package hijack

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestGetPtr(t *testing.T) {
	// Happy path test case
	value := 42
	rv := reflect.ValueOf(&value)
	ptr := getPtr(rv)

	// Verify that the obtained pointer matches the expected value
	expectedPtr := unsafe.Pointer(&value)
	if ptr != expectedPtr {
		t.Errorf("Expected pointer: %p, Got: %p", expectedPtr, ptr)
	}

	// Sad path test case
	var nilValue *int
	rv = reflect.ValueOf(nilValue)
	ptr = getPtr(rv)

	// Verify that the obtained pointer is nil
	if ptr != nil {
		t.Errorf("Expected nil pointer, Got: %p", ptr)
	}
}
