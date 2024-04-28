package file

import "os"

// Stat - Wrapper for the os.Stat() function to avoid a lot of work if golang changes things
func Stat(fileName string) (os.FileInfo, error) {
	return os.Stat(fileName)
}
