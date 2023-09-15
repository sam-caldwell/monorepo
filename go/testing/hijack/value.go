package hijack

import "unsafe"

// value - A simple pointer structure
type value struct {
	_   uintptr
	ptr unsafe.Pointer
}
