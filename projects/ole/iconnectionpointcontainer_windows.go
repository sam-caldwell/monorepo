//go:build windows
// +build windows

//lint:file-ignore SA1019 skipping because of deprecated syscall

package ole

import (
	"syscall"
	"unsafe"
)

func (v *IConnectionPointContainer) EnumConnectionPoints(points interface{}) error {
	return NewError(E_NOTIMPL)
}

func (v *IConnectionPointContainer) FindConnectionPoint(iid *GUID, point **IConnectionPoint) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().FindConnectionPoint,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(point)))
	if hr != 0 {
		err = NewError(hr)
	}
	return
}
