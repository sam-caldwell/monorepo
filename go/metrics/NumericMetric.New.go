package metrics

// New - initialize the numeric metric object
func (o *NumericMetric[ValueType, CounterType]) New(size uint8) *NumericMetric[ValueType, CounterType] {
	o.lock.Lock()
	defer o.lock.Unlock()

	o.movingWindow = size
	o.value = make([]ValueType, size)
	//Note: we assume o.count is zeroed

	return o
}
