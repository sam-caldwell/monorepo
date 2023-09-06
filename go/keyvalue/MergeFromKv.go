package keyvalue

/*
 * keyvalue.MergeFromKv
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Merge a given source keyvalue into the current struct.
 * Note this will overwrite existing keys
 */

// MergeFromKv - Merge a given source KeyValue into the current struct
func (kv *KeyValue) MergeFromKv(source KeyValue) {
	for key, value := range source.data {
		kv.data[key] = value
	}
}
