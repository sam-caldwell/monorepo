//go:build windows
// +build windows

package ole

import (
	"fmt"
	"golang.org/x/sys/windows"
	"unicode/utf16"
)

// errstr converts error code to string.
func errstr(errno int) string {
	// ask windows for the remaining errors
	var flags uint32 = windows.FORMAT_MESSAGE_FROM_SYSTEM | windows.FORMAT_MESSAGE_ARGUMENT_ARRAY | windows.FORMAT_MESSAGE_IGNORE_INSERTS
	b := make([]uint16, 300)
	n, err := windows.FormatMessage(flags, 0, uint32(errno), 0, b, nil)
	if err != nil {
		return fmt.Sprintf("error %d (FormatMessage failed with: %v)", errno, err)
	}
	// trim terminating \r and \n
	for ; n > 0 && (b[n-1] == '\n' || b[n-1] == '\r'); n-- {
	}
	return string(utf16.Decode(b[:n]))
}
