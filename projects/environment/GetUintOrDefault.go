package environment

// GetUintOrDefault - Return the uint value of a given environment variable or default (if error)
func GetUintOrDefault(name string, defaultValue uint) (r uint) {
	value, err := GetIntp(&name)
	if err != nil {
		return defaultValue
	}
	return uint(value)
}
