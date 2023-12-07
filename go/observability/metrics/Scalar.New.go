package metrics

// New - initialize the Scalar metric object
func (o *Scalar[ValueType]) New(size uint8) *Scalar[ValueType] {
	o.lock.Lock()
	defer o.lock.Unlock()

	o.movingWindow = size
	o.value = make([]ValueType, size)
	//Note: we assume o.count is zeroed

	return o
}
