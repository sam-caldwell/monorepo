package environment

// SetInt - Set Environment variable value
func SetInt(name string, value int) error {
	var anyValue any = value
	return SetAnyp(&name, &anyValue)
}
