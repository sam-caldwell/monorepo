//go:build linux
// +build linux

package systemrecon

/*
 * SystemInfo ()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * SystemInfo() for Linux
 *
 * 	Return the amount of RAM in the system (in KB)
 */

// RamSize - Return the ram size in KB
func RamSize() (value int, err error) {
	//var info string
	//if info, err = SystemInfo(); err != nil {
	//	return value, err
	//}
	//if info, ok := info[words.MemTotal]; !ok {
	//	return value, fmt.Errorf(errors.InternalError)
	//} else {
	//	value = info
	//}
	return value, err
}
