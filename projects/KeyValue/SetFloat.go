package keyvalue

/*
 * KeyValue.SetFloat
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * For a given key and value, store the same in our key-value store
 */

// SetFloat - For a given key and value, store the same in our key-value store
func (kv *KeyValue) SetFloat(key string, value float64) {
	kv.Initialize(0, preserve)
	kv.data[key] = value
}
