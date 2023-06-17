package environment

// RequireString - Return the string value of a given environment variable (return error if not set)
func RequireString(name string) (r string, e error) {
	return RequireStringp(&name)
}
