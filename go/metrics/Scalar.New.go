package metrics

// New - initialize the Scalar metric object
func (o *Scalar[ValueType, CounterType]) New(size uint8) *Scalar[ValueType, CounterType] {
	o.lock.Lock()
	defer o.lock.Unlock()

	o.movingWindow = size
	o.value = make([]ValueType, size)
	//Note: we assume o.count is zeroed

	return o
}
