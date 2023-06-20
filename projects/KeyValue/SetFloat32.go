package keyvalue

/*
 * KeyValue.SetFloat32
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * For a given key and value, store the same in our key-value store
 */

// SetFloat32 - For a given key and value, store the same in our key-value store
func (kv *KeyValue) SetFloat32(key string, value float32) {
	kv.Initialize(0, preserve)
	kv.data[key] = value
}
