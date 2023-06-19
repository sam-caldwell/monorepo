package simple

import "fmt"

// ListString - return teh set as a list of strings
func (set *Set) ListString() (result []string) {
	if set.data != nil {
		for item := range set.data {
			result = append(result, fmt.Sprintf("%v", item))
		}
	}
	return result
}
