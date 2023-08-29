package file

// Exists - Return boolean value if file exists
func Exists(fileName string) bool {
	return Existsp(&fileName)
}
