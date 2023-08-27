package keyvalue

/*
 * keyvalue.SetString
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * For a given key and value, store the same in our key-value store
 */

// SetString - For a given key and value, store the same in our key-value store
func (kv *KeyValue) SetString(key string, value string) {
	kv.Initialize(0, preserve)
	kv.data[key] = value
}
