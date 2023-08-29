package environment

// SetBool - Set Environment variable value
func SetBool(name string, value bool) error {
	var anyValue any = value
	return SetAnyp(&name, &anyValue)
}
