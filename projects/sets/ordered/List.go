package ordered

// List - return a list of items
func (set *Set) List() (result []any) {
	set.lock.RLock()
	defer set.lock.RUnlock()

	if len(set.data) > 0 {
		result = make([]any, len(set.data))
		copy(result, set.data)
	}

	return result
}
