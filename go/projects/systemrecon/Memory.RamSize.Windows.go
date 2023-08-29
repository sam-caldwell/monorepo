//go:build windows
// +build windows

package systemrecon

/*
 * SystemInfo ()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * SystemInfo() for Windows
 *
 * 	Return the amount of RAM in the system (in KB)
 */

// RamSize - Return the ram size in KB
func RamSize() (value int, err error) {
	//ToDo: return RamSize in KB
	return value, err
}
