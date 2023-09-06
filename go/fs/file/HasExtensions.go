package file

// HasExtensions - return boolean if file has any of a list of extensions
func HasExtensions(filename string, extensions []string) (result bool) {
	for _, thisExtension := range extensions {
		result = result || HasExtension(filename, thisExtension)
	}
	return result
}
