package file

const (
	jsonExtension = ".json"
)

// HasJsonExtension - return boolean result if JSON extension exists (case-insensitive)
func HasJsonExtension(filename string) bool {
	return GetExtensionp(&filename) == jsonExtension
}
