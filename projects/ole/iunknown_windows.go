//go:build windows
// +build windows

//lint:file-ignore SA1019 skipping because of deprecated syscall

package ole

import (
	"fmt"
	"reflect"
	"syscall"
	"unsafe"
)

func reflectQueryInterface(self interface{}, method uintptr, interfaceID *GUID, obj interface{}) (err error) {
	selfValue := reflect.ValueOf(self).Elem()
	objValue := reflect.ValueOf(obj).Elem()

	hr, _, _ := syscall.Syscall(
		method,
		3,
		selfValue.UnsafeAddr(),
		uintptr(unsafe.Pointer(interfaceID)),
		objValue.Addr().Pointer())
	if hr != 0 {
		err = NewError(hr)
	}
	return
}

func queryInterface(unk *IUnknown, iid *GUID) (disp *IDispatch, err error) {
	hr, _, _ := syscall.Syscall(
		unk.VTable().QueryInterface,
		3,
		uintptr(unsafe.Pointer(unk)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&disp)))
	if hr != 0 {
		err = NewError(hr)
	}
	return
}

func addRef(unk *IUnknown) int32 {
	ret, _, _ := syscall.Syscall(
		unk.VTable().AddRef,
		1,
		uintptr(unsafe.Pointer(unk)),
		0,
		0)
	return int32(ret)
}

func release(unk *IUnknown) int32 {
	if unk == nil {
		// Return an appropriate error code, or handle the error as required.
		return 0
	}
	ret, _, err := syscall.Syscall(
		unk.VTable().Release,
		1,
		uintptr(unsafe.Pointer(unk)),
		0, 0)

	if err != 0 {
		panic(fmt.Sprintf("Error: %v", err))
	}

	return int32(ret)
}
