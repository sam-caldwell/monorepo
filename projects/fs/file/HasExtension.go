package file

import "strings"

// HasExtension - return boolean result if a given extension exists on the provided filename.
//
// The extension must start with a period (.) character, and the extension is case insensitive
// for the comparison.
func HasExtension(filename, extension string) bool {
	extension = strings.ToLower(extension)
	return GetExtension(filename) == extension
}
