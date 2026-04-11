package hijack

import (
	"reflect"
	"unsafe"
)

// getPtr - return a pointer to the given object
func getPtr(object reflect.Value) unsafe.Pointer {
	return (*value)(unsafe.Pointer(&object)).ptr
}
