package list

import (
	"fmt"
)

// String - Dump the contents of our SmartList to a list of strings.
func (list *SmartList) String() []string {
	result := make([]string, len(*list))
	for i, v := range *list {
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
}
