package metrics

// Set - Set the state we will track
func (metric *StateMetric[T]) Set(input T) {

	metric.lock.Lock()
	defer metric.lock.Unlock()

	metric.value = input
}
