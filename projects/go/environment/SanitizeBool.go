package environment

// SanitizeBool - Return the sanitized boolean value of a given environment variable name.
func SanitizeBool(name string) (r bool, e error) {
	return RequireBoolp(&name)
}
