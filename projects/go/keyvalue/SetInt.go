package keyvalue

/*
 * keyvalue.SetInt
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * For a given key and value, store the same in our key-value store
 */

// SetInt - For a given key and value, store the same in our key-value store
func (kv *KeyValue) SetInt(key string, value int) {
	kv.Initialize(0, preserve)
	kv.data[key] = value
}
