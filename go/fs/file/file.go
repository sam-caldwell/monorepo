package file

import (
	"os"
	"sync"
)

// File - This is a wrapper around file-access logic.
//
//	     The goal is to create an abstraction layer so changes to golang will not require
//	     significant effort to update all projects.
//
//			(c) 2023 Sam Caldwell.  MIT License
type File struct {
	lock   sync.Mutex
	handle *os.File
}
