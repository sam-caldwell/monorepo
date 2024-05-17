package keyvalue

/*
 * keyvalue.SetBool
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */

// Set - For a given key and value, store the same in our key-value store
func (kv *KeyValue[KeyType, ValueType]) Set(key KeyType, value ValueType) {
	kv.Initialize(0, preserve)
	kv.lock.Lock()
	defer kv.lock.Unlock()
	kv.data[key] = value
}
