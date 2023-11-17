package metrics

import (
	"sync"
)

/*
 * Theory of Use
 *	The NumericMetric is intended to store any scalar metric.
 *
 *  (For the new coder, we have all been there.  Let me explain.  A scalar metric is a measurable
 *   value (such as distance or temperature).  We just give it a fancy name, so we seem smarter
 *   and can confuse people.)
 */

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
