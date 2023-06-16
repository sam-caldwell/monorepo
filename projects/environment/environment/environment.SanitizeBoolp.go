package environment

// SanitizeBoolp - Return the sanitized boolean value of a given environment variable name.
func SanitizeBoolp(name *string) (r bool, e error) {
	return RequireBoolp(name)
}
