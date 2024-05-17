package keyvalue

/*
 * keyvalue.MergeFromKv
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */

// MergeFromKv - Merge a given source KeyValue into the current struct
func (kv *KeyValue[KeyType, ValueType]) MergeFromKv(source KeyValue[KeyType, ValueType]) {
	kv.lock.Lock()
	source.lock.Lock()
	defer kv.lock.Unlock()
	defer source.lock.Unlock()
	for key, value := range source.data {
		kv.data[key] = value
	}
}
