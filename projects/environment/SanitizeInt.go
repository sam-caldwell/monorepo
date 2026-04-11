package environment

// SanitizeInt - return sanitized integer env var
func SanitizeInt(name string, min, max int) (int, error) {
	return SanitizeIntp(&name, min, max)
}
