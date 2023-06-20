package keyvalue

/*
 * KeyValue.FindKey()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Return boolean result and value for a given key in the key-value set
 */

// FindKey -Return boolean result and value for a given key in the key-value set
func (kv *KeyValue) FindKey(key string) (value any, found bool) {
	value, found = kv.data[key]
	return value, found
}
