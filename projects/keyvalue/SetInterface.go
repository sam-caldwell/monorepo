package keyvalue

/*
 * projects/keyvalue/SetInterface.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the keyvalue.SetInterface() method which will allow us to store an arbitrary data type
 * as the value of a string-keyed key-value store.
 *
 * See README.md
 */

// SetInterface - Allow an arbitrary struct to be stored...or anything else really.
func (kv *KeyValue) SetInterface(key string, value any) {
	kv.Initialize(0, preserve)
	kv.data[key] = value
}
