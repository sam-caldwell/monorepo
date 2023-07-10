//go:build windows
// +build windows

package systemrecon

import (
	"syscall"
	"unsafe"
)

// ChangeFlags - Change our memory protection flags (Darwin/MacOS version)
func ChangeFlags(address uintptr, length int, newFlags uint32, oldFlags unsafe.Pointer) (err error) {
	/*
	 * Warning:
	 * This is about as unsafe as you can get.  We're playing with memory here (and I intentionally avoid Windows
	 * when possible...so this probably is not the best way to do this.
	 *
	 * As with all things, there's how the world works, and then there's what those special people at Microsoft
	 * do...which is always different.  (insert Seattle joke here.)
	 */
	/*
	 * This function will call Windows procVirtualProtect and ask the operating system to change our permissions
	 * the original permissions will be written to oldFlags, which we can use later to revert the change.
	 *
	 * See https://learn.microsoft.com/en-us/windows/win32/memory/memory-protection-constants
	 */
	var procVirtualProtect = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")

	exitCode, _, _ := procVirtualProtect.Call(address, uintptr(length), uintptr(newFlags), uintptr(oldFlags))

	/*
	 * If it's not backwards, it's not Microsoft...
	 *
	 * In Windows exitCode 0 is an error state but a non-zero value is ok.
	 * EVERYWHERE else, exitCode 0 is ok and non-zero is an error state.
	 */
	if exitCode == 0 {
		err = syscall.GetLastError()
	}
	return err
}
