package misc

import "fmt"

// AnyListToStringList - Convert an []any list to []string
func AnyListToStringList(list *[]any) []string {
	result := make([]string, len(*list))
	for i, v := range *list {
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
}
