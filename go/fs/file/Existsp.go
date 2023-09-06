package file

import (
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
)

// Existsp - Return boolean value if file exists
func Existsp(name *string) bool {
	fileInfo, err := os.Stat(*name)
	return !os.IsNotExist(err) && (fileInfo != nil) && !fileInfo.IsDir()
}
