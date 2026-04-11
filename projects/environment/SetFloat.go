package environment

// SetFloat - Set Environment variable value
func SetFloat(name string, value float64) error {
	var anyValue any = value
	return SetAnyp(&name, &anyValue)
}
