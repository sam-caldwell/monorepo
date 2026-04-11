package environment

// SetBoolp - Set Environment variable value (pointer)
func SetBoolp(name *string, value bool) error {
	var anyValue any = value
	return SetAnyp(name, &anyValue)
}
