package ordered

// Add - add item to set if the item is the same type as the set
func (set *Set) Add(item any) (err error) {
	//Bail on nil inputs.  not worth the time.
	if item == nil {
		return err
	}

	set.lock.Lock()
	defer set.lock.Unlock()

	// Type check the item and add it if unique
	if err = set.TypeCheck(item); err == nil {
		if !set.seenBefore(item) {
			set.data = append(set.data, item)
		}
	}
	return err
}
