package keyvalue

import (
	"reflect"
)

// FindValue - Return the key and boolean true where a given value exists
//
//	 Note: this is a linear search.  If you're using a larger data set, you may want to
//	 find a better means of searching values.
//
//		(c) 2023 Sam Caldwell.  MIT License
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
