//go:build !windows
// +build !windows

package geometry

import (
	"os"
	"syscall"
	"unsafe"
)

func (ws *TextWindow) GetProperties() {
	fd := int(os.Stdin.Fd())
	var window TextWindow
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&window)))
	ws.Rows, ws.Cols, ws.Xpixel, ws.Ypixel = window.Rows, window.Cols, window.Xpixel, window.Ypixel
	if errno != 0 {
		ws.Cols = defaultColumnWidth
	}
}
