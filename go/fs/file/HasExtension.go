package file

import "strings"

// HasExtension - return boolean result if a YAML extension exists
func HasExtension(filename, extension string) bool {
	extension = strings.ToLower(extension)
	return GetExtension(filename) == extension
}
