package keyvalue

/*
 * keyvalue.RenameKey
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * 	Given a set of key-value data, rename an existing key to a new name.
 */

// RenameKey - Return the maximum width of all values in the current KeyValue struct
func (kv *KeyValue) RenameKey(currKey, newKey string) bool {
	if kv.data != nil {
		if _, ok := kv.data[currKey]; ok {
			kv.data[newKey] = kv.data[currKey]
			delete(kv.data, currKey)
			return true
		}
	}
	return false
}
