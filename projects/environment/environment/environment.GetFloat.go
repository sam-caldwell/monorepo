package environment

// GetFloat - Return the floating-point value of a given environment variable
func GetFloat(name string) (r float64, e error) {
	return GetFloatp(&name)
}
