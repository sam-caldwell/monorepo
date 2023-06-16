package misc

// StringListToAnyList - Convert a list of strings to an []any list
func StringListToAnyList(list *[]string) []any {
	result := make([]any, len(*list))
	for i, v := range *list {
		result[i] = v
	}
	return result
}
