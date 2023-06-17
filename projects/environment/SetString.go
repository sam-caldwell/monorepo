package environment

// SetString - SetString Environment variable value (pointer)
func SetString(name string, value string) error {
	var anyValue any = value
	return SetAnyp(&name, &anyValue)
}
