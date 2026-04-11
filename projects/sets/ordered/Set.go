package ordered

/*
 * projects/sets/ordered/Set.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */

import "sync"

// Set - Create map of any to its type
type Set[T comparable] struct {
	lock sync.RWMutex
	data []T
}
