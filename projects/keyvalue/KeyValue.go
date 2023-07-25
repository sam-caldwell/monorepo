package keyvalue

/*
 * keyvalue
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Create a key-value struct and methods to make working with KV data easier.
 */

type Map map[string]any

type KeyValue struct {
	data Map
}
