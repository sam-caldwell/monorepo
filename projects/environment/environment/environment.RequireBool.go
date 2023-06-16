package environment

// RequireBool - Return the boolean value of an environment variable (returning error if empty).
func RequireBool(name string) (bool, error) {
	return RequireBoolp(&name)
}
