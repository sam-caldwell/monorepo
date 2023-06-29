//go:build windows
// +build windows

//lint:file-ignore SA1019 skipping because of deprecated syscall

package ole

import (
	"syscall"
	"unsafe"
)

func getClassInfo(disp *IProvideClassInfo) (tinfo *ITypeInfo, err error) {
	hr, _, _ := syscall.Syscall(
		disp.VTable().GetClassInfo,
		2,
		uintptr(unsafe.Pointer(disp)),
		uintptr(unsafe.Pointer(&tinfo)),
		0)
	if hr != 0 {
		err = NewError(hr)
	}
	return
}
