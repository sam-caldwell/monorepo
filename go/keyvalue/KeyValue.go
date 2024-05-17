package keyvalue

import "sync"

/*
 * keyvalue
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Create a key-value struct and methods to make working with KV data easier.
 */

type KeyValue[KeyType comparable, ValueType any] struct {
	lock sync.RWMutex
	data map[KeyType]ValueType
}
