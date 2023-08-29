package directory

import (
	os2 "github.com/sam-caldwell/monorepo/go/projects/v2/wrappers/os"
)

// Existsp - Return boolean value if directory exists
func Existsp(name *string) bool {
	fileInfo, err := os2.Stat(*name)
	return !os2.IsNotExist(err) && (fileInfo != nil) && fileInfo.IsDir()
}
