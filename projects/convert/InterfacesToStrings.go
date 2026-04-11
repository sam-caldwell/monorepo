package convert

// InterfacesToStrings - convert []interface to []string
//
//	(c) 2023 Sam Caldwell.  MIT License
func InterfacesToStrings(input []interface{}) []string {
	result := make([]string, len(input))
	for i, v := range input {
		result[i] = v.(string)
	}
	return result
}
