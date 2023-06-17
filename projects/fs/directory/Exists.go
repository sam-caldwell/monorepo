package directory

// Exists - Return boolean value if directory exists
func Exists(name string) bool {
	return Existsp(&name)
}
