package keyvalue

/*
 * KeyValue.SetBool
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * For a given key and value, store the same in our key-value store
 */

// SetBool - For a given key and value, store the same in our key-value store
func (kv *KeyValue) SetBool(key string, value bool) {
	kv.Initialize(0, preserve)
	kv.data[key] = value
}
