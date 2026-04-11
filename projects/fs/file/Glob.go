package file

import (
	"path/filepath"
)

// Glob - a wrapper for the filepath.Glob() function to avoid a lot of work if golang changes their function.
func Glob(path string, pattern string) (fileList []string, err error) {
	return filepath.Glob(filepath.Join(path, pattern))
}
