package environment

// GetStringOrDefault - Return the string value of a given environment variable or default (if error)
func GetStringOrDefault(name string, defaultValue string) (r string) {
	r, err := GetStringp(&name)
	if err != nil {
		return defaultValue
	}
	return r
}
