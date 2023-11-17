package metrics

// NewNumericMetric initializes and returns a pointer to a new instance of NumericMetric.
func NewNumericMetric[ValueType BigNumber, CounterType AnyNumber](sz uint8) *NumericMetric[ValueType, CounterType] {
	return &NumericMetric[ValueType, CounterType]{
		value: make([]ValueType, sz),
		//count:        0, ...this is implied
		movingWindow: sz,
	}
}
