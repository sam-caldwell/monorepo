package convert

// StringListToAnyList - Convert a list of strings to an []any list
//
//	(c) 2023 Sam Caldwell.  MIT License
func StringListToAnyList(list *[]string) []any {
	result := make([]any, len(*list))
	for i, v := range *list {
		result[i] = v
	}
	return result
}
