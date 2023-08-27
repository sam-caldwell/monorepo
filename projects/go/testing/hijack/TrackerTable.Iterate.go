package hijack

// iterate - Iterates over the map and performs the specified operation for each key-value pair
func (table *TrackerTable) iterate(operation IteratorFunction) (err error) {
	/*
	 * Make sure we have an initialized patches map or bail like you're on the Titanic
	 * There's nothing to iterate over here because we have no data in the map.
	 */
	if table.patches == nil {
		return nil
	}
	/*
	 * iterate over every element of our patches map and for each element
	 * execute the IteratorFunction and return on first failure.
	 */
	for key, value := range table.patches {
		if err = operation(key, value); err != nil {
			return err
		}
	}
	return err
}
