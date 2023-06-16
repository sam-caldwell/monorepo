package simple

// List - return a list of items
func (set *Set) List() (result []any) {
	if set.data != nil {
		for item := range set.data {
			result = append(result, item)
		}
	}
	return result
}
