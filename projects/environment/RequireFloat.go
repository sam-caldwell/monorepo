package environment

// RequireFloat - Return the Float value of an environment variable (returning error if empty).
func RequireFloat(name string) (float64, error) {
	return RequireFloatp(&name)
}
