package file

import (
	"path/filepath"
	"strings"
)

// GetExtension - return the given file's extension
func GetExtension(name string) string {
	return strings.ToLower(filepath.Ext(name))
}
