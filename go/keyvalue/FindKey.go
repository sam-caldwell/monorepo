package keyvalue

/*
 * keyvalue.FindKey()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Return boolean result and value for a given key in the key-value set
 */

// FindKey -Return boolean result and value for a given key in the key-value set
func (kv *KeyValue[KeyType, ValueType]) FindKey(key KeyType) (value ValueType, found bool) {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	value, found = kv.data[key]
	return value, found
}
