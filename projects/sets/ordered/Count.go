package ordered

// Count - return a count of the records in a set
func (set *Set) Count() int {
	set.lock.RLock()
	defer set.lock.RUnlock()
	return len(set.data)
}
