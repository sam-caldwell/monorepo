package environment

// SetAny - Set an environment variable of any type
func SetAny(name string, value any) error {
	return SetAnyp(&name, &value)
}
