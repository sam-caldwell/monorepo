package convert

// InterfacesToStrings - convert []interface to []string
func InterfacesToStrings(input []interface{}) []string {
	result := make([]string, len(input))
	for i, v := range input {
		result[i] = v.(string)
	}
	return result
}
