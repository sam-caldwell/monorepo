package environment

// RequireInt - Return the int value of a given environment variable (return error if not set)
func RequireInt(name string) (r int, e error) {
	return RequireIntp(&name)
}
