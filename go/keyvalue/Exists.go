package keyvalue

// Exists - Return true if key exists, false if it does not.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (kv *KeyValue[KeyType, ValueType]) Exists(key KeyType) bool {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	_, found := kv.data[key]
	return found
}
