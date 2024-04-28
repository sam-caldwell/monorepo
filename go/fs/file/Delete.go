package file

import "os"

// Delete - Delete the given file and return any error.
//
//	This is just a point of abstraction in case the golang
//	community changes things, so it doesn't require a lot of
//	effort.
func Delete(fileName string) error {
	return os.Remove(fileName)
}
