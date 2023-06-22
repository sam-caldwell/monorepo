//go:build !windows
// +build !windows

package errors

// errString converts error code to string.
func errString(errno int) string {
	return ""
}
