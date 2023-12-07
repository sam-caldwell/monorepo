package metrics

import (
	"sync"
)

// MetricStore - A structure for organizing a collection of metrics in a thread safe way that minimizes GC
type MetricStore struct {
	lock sync.Mutex

	// A set of accumulated errors between checks
	// This allows error checking to be deferred when methods are chained.
	err []error

	// This map stores our metric collector objects by name
	data map[string]any
}
