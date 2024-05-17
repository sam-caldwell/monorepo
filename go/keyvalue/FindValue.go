package keyvalue

import (
	"reflect"
)

/*
 * keyvalue.FindValue()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Return boolean result and the name of the key where a given value
 * is found.
 *
 * Warning: This is a linear search.
 */

// FindValue - Return the key where a given value exists in the key-value store and boolean true if found.
func (kv *KeyValue[KeyType, ValueType]) FindValue(target any) (key KeyType, found bool) {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	var empty KeyType
	if kv.data != nil {
		for key, value := range kv.data {
			if reflect.DeepEqual(value, target) {
				return key, true
			}
		}
	}
	return empty, found
}
