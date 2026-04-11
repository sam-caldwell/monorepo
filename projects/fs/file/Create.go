package file

import "os"

// Create - given a filename, create a file and write the given bytes thereto.
// This is a wrapper because it seems like go keeps changing stuff and making
// me do a lot of work to keep up with their opinionated breaking changes.
func Create(fileName string, content []byte, perms os.FileMode) (err error) {
	return os.WriteFile(fileName, content, perms)
}
