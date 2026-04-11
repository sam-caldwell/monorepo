package simple

import (
	"fmt"
	"sort"
)

// ListString - return teh set as a list of strings
func (set *Set[T]) ListString(sortResult bool) (result []string) {
	if set.data != nil {
		for item := range set.data {
			result = append(result, fmt.Sprintf("%v", item))
		}
	}
	if sortResult {
		sort.Strings(result)
	}
	return result
}
