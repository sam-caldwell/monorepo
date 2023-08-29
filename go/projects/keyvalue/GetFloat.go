package keyvalue

/*
 * keyvalue.GetFloat
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Return the value corresponding to a given key as float or typecheck error
 */

// GetFloat - Return the value corresponding to a given key as float or typecheck error
func (kv *KeyValue) GetFloat(key string) (value float64, err error) {
	return kv.GetFloat64(key)
}
