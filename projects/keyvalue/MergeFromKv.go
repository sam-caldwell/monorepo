package keyvalue

/*
 * keyvalue.MergeFromKv
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */

// MergeFromKv - Given the current struct and another KeyValue struct, merge the two into the current.
//
//	     This will lock both structs for the duration of the merge operation to avoid
//	     integrity issues.
//
//			(c) 2023 Sam Caldwell.  MIT License
func (kv *KeyValue[KeyType, ValueType]) MergeFromKv(source *KeyValue[KeyType, ValueType]) {
	kv.lock.Lock()
	source.lock.Lock()
	defer kv.lock.Unlock()
	defer source.lock.Unlock()
	for key, value := range source.data {
		kv.data[key] = value
	}
}
