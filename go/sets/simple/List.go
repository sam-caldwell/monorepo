package simple

// List - Return a list of items
func (set *Set[T]) List() (result []any) {

	if set.data != nil {

		for item := range set.data {

			result = append(result, item)

		}

	}

	return result

}
