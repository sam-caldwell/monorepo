package metrics

import (
	"sync"
)

// NumericMetric - A generic metric for all things big and numeric.
type NumericMetric[ValueType AnyNumber, CounterType AnyNumber] struct {

	// Lock - some safety for the concurrent stuff.
	lock sync.Mutex

	// value - The FIFO metric value bucket array for storing a moving window of metrics.
	value []ValueType

	// count - a total count of all instances of this metric for all time
	count CounterType

	// movingWindow - A moving window for calculating min/max and moving averages
	// This is used to determine the value bucket array size during initialization.
	movingWindow uint8 //The moving window for min/max
}
