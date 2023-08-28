package file

import (
	os2 "github.com/sam-caldwell/monorepo/v2/projects/go/wrappers/os"
)

// Existsp - Return boolean value if file exists
func Existsp(name *string) bool {
	fileInfo, err := os2.Stat(*name)
	return !os2.IsNotExist(err) && (fileInfo != nil) && !fileInfo.IsDir()
}
