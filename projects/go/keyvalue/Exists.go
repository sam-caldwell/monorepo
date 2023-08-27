package keyvalue

// Exists - Return true if key exists, false if it does not.
func (kv *KeyValue) Exists(key string) bool {
	_, found := kv.data[key]
	return found
}
