package environment

// SetIntp - Set Environment variable value (pointer)
func SetIntp(name *string, value int) error {
	var anyValue any = value
	return SetAnyp(name, &anyValue)
}
