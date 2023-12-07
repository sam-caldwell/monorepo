package metrics

// Error - If error state(s) occur, return the top-most (last)
func (store *MetricStore) Error() (err error) {
	//
	// If empty return nil
	//
	if (err == nil) || (len(store.err) == 0) {
		return nil
	}
	//
	// Get the last element in the list
	//
	err = store.err[len(store.err)-1]
	//
	// Remove the last element
	//
	store.err = store.err[:len(store.err)-1]
	//
	// Return the result
	//
	return err
}
