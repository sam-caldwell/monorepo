package metrics

import (
	"sync"
)

// MetricStore - A structure for organizing a collection of metrics in a thread safe way that minimizes GC
type MetricStore struct {
	lock sync.Mutex
	err  []error
	data map[string]any
}
