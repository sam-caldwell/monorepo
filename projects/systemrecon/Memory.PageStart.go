package systemrecon

import "syscall"

// PageStart - make syscall to get the address where our page starts
func PageStart(ptr uintptr) uintptr {
	return ptr & ^(uintptr(syscall.Getpagesize() - 1))
}
