package environment

// GetString - Return the int value of a given environment variable
func GetString(name string) (r string, e error) {
	return GetStringp(&name)
}
