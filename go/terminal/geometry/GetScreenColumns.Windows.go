//go:build windows
// +build windows

package geometry

const (
	defaultColumnWidth = 80
)

//ToDo: We need to make this consistent with the posix version

// ColumnSize - Return Screen width of a terminal/console in columns (Windows Version)
func GetScreenColumns() int {
	handle := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	var info syscall.ConsoleScreenBufferInfo
	err := syscall.GetConsoleScreenBufferInfo(handle, &info)
	if err != nil {
		return defaultColumnWidth
	}
	return info.Window.Right - info.Window.Left
}
