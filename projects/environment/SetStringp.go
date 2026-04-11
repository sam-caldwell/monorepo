package environment

// SetStringp - Set Environment variable value (pointer)
func SetStringp(name *string, value *string) error {
	var anyValue any = *value
	return SetAnyp(name, &anyValue)
}
