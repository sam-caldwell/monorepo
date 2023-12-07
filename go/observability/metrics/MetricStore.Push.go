package metrics

/*
// Push - Push a metric value to the Metric Store.
func (store *MetricStore) Push(name *string, value any) *MetricStore {
	//
	// Lookup the metric object
	//
	metric, ok := store.data[*name]

	if ok {
		//
		// If the metric exists, push the new value to the metric object
		//
		if err := metric.Push(value); err != nil {
			//
			// retain any error returned from the metric Push operation.
			//
			store.err = append(store.err, err)
		}
	} else {
		//
		// If the metric does not exist, append the error to the store.errors array
		// for later analysis.
		//
		store.err = append(store.err, fmt.Errorf("unknown metric (%s)", *name))
	}
	//
	// Return the MetricStore object reference to allow chained operations.
	//
	return store
}
*/
