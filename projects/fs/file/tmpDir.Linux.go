//go:build linux || darwin
// +build linux darwin

package file

const (
	// tmpDir - Temporary Directory
	//
	//  This constant is defined for Linux, macOS.  Windows must be supported separately.
	tmpDir = "/tmp"
)
