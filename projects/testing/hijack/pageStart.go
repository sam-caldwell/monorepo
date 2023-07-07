package hijack

import "syscall"

// pageStart - make syscall to get the address where our page starts
func pageStart(ptr uintptr) uintptr {
	return ptr & ^(uintptr(syscall.Getpagesize() - 1))
}
