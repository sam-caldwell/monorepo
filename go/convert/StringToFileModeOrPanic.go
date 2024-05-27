package convert

import (
	"os"
	"strconv"
)

// StringToFileModeOrPanic - Convert a string permission (e.g. 0600) to FileMode numeric value
//
//	(c) 2023 Sam Caldwell.  MIT License
func StringToFileModeOrPanic(value string) os.FileMode {
	number, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return os.FileMode(number)
}
