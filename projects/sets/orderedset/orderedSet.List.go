package orderedset

// List - return a list of items
func (set *Set) List() (result []any) {
	if set.data != nil {
		for _, item := range set.data {
			result = append(result, item)
		}
	}
	return result
}
