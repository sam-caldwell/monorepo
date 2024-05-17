package keyvalue

/*
 * keyvalue.RenameKey
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */

// RenameKey - Return the maximum width of all values in the current KeyValue struct
func (kv *KeyValue[KeyType, ValueType]) RenameKey(currKey, newKey KeyType) bool {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	if kv.data != nil {
		if _, ok := kv.data[currKey]; ok {
			kv.data[newKey] = kv.data[currKey]
			delete(kv.data, currKey)
			return true
		}
	}
	return false
}
