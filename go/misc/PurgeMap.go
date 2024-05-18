package misc

import "runtime"

// PurgeMap - reset the memory state of a given generic map and call gc explicitly.
func PurgeMap[KeyType comparable, ValueType any](mp *map[KeyType]ValueType) {
	m := make(map[KeyType]ValueType)
	mp = &m
	runtime.GC()
}
