package environment

// GetInt - Return the int value of a given environment variable
func GetInt(name string) (r int, e error) {
	return GetIntp(&name)
}
