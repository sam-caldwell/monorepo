package file

// HasExtension - return boolean result if a YAML extension exists
func HasExtension(filename, extension string) bool {
	return GetExtension(filename) == extension
}
