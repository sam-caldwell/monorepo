package environment

// SanitizeFloat - get the float env var and bounds check the value
func SanitizeFloat(name string, min float64, max float64) (float64, error) {
	return SanitizeFloatp(&name, min, max)
}
