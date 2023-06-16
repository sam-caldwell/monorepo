package environment

// SetFloatp - Set Environment variable value (pointer)
func SetFloatp(name *string, value float64) error {
	var anyValue any = value
	return SetAnyp(name, &anyValue)
}
