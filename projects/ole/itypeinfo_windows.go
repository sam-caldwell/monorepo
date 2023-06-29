//go:build windows
// +build windows

//lint:file-ignore SA1019 skipping because of deprecated syscall

package ole

import (
	"syscall"
	"unsafe"
)

func (v *ITypeInfo) GetTypeAttr() (tattr *TYPEATTR, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().GetTypeAttr),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&tattr)),
		0)
	if hr != 0 {
		err = NewError(hr)
	}
	return
}
