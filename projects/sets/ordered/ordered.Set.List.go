package ordered

// List - return a list of items
func (set *Set) List() (result []any) {
	if set.data != nil {
		result = append(result, set.data...)
	}
	return result
}
