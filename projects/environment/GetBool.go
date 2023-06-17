package environment

// GetBool - Return the boolean value of an environment variable.
func GetBool(name string) (bool, error) {
	return GetBoolp(&name)
}
