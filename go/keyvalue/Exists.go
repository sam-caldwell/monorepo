package keyvalue

// Exists - Return true if key exists, false if it does not.
func (kv *KeyValue[KeyType, ValueType]) Exists(key KeyType) bool {
	_, found := kv.data[key]
	return found
}
