package LogTarget

import "os"

// File - Send log output of the given format to file
//
//	(c) 2023 Sam Caldwell.  MIT License
type File struct {
	config ConfigureFile
	file   *os.File
}
