package file

import (
	"os"
)

// Existsp - Return boolean if file exists
//
//	     The goal is to create an abstraction layer so changes to golang will not require
//	     significant effort to update all projects.
//
//			(c) 2023 Sam Caldwell.  MIT License
func (fp *File) Existsp(fileName *string) bool {
	if fileName == nil {
		return false
	}
	fileInfo, err := os.Stat(*fileName)
	return !os.IsNotExist(err) && (fileInfo != nil) && !fileInfo.IsDir()
}
