package metrics

// Get - Get returns the internal state
func (metric *StateMetric[T]) Get() T {

	metric.lock.Lock()
	defer metric.lock.Unlock()

	return metric.value
}
