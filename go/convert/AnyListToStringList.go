package convert

import "fmt"

// AnyListToStringList - Convert an []any list to []string
//
//	(c) 2023 Sam Caldwell.  MIT License
func AnyListToStringList(list *[]any) []string {
	result := make([]string, len(*list))
	for i, v := range *list {
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
}
