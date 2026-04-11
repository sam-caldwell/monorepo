package environment

// GetIntOrDefault - Return the int value of a given environment variable or default (if error)
func GetIntOrDefault(name string, defaultValue int) (r int) {
	r, err := GetIntp(&name)
	if err != nil {
		return defaultValue
	}
	return r
}
