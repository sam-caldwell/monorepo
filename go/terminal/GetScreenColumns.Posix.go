//go:build !windows
// +build !windows

package terminal

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	defaultColumnWidth = 80
)

type winsize struct {
	Rows   uint16
	Cols   uint16
	Xpixel uint16
	Ypixel uint16
}

// GetScreenColumns - Return Screen width of a terminal/console (tty/pty) in columns (Posix version)
func GetScreenColumns() int {
	var ws winsize

	fd := int(os.Stdin.Fd())

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&ws)))
	if errno != 0 {
		return defaultColumnWidth
	}
	return int(ws.Cols)
}
