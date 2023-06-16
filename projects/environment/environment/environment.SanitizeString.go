package environment

// SanitizeString - return the string env var (sanitized by regex pattern)
func SanitizeString(name, pattern string) (string, error) {
	return SanitizeStringp(&name, &pattern)
}
