package keyvalue

/*
 * KeyValue.SetFloat64
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * For a given key and value, store the same in our key-value store
 */

// SetFloat64 - For a given key and value, store the same in our key-value store
func (kv *KeyValue) SetFloat64(key string, value float64) {
	kv.Initialize(0, preserve)
	kv.data[key] = value
}
