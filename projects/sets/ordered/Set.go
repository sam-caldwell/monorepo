package ordered

import "sync"

// Set - Create map of any to its type
type Set struct {
	lock sync.RWMutex
	data []any
}
