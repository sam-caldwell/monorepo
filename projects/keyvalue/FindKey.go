package keyvalue

// FindKey -Return boolean result and value for a given key in the key-value set
//
//	(c) 2023 Sam Caldwell.  MIT License
func (kv *KeyValue[KeyType, ValueType]) FindKey(key KeyType) (value ValueType, found bool) {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	value, found = kv.data[key]
	return value, found
}
