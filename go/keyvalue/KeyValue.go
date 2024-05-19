package keyvalue

import "sync"

// KeyValue - a generic key-value struct.
//
//	    This struct has methods which should make it easier to work with
//	    any generic key-value use-case, including parsing key-value files
//	    command shell inputs, etc.
//
//		(c) 2023 Sam Caldwell.  MIT License
type KeyValue[KeyType comparable, ValueType any] struct {
	lock sync.RWMutex
	data map[KeyType]ValueType
}
