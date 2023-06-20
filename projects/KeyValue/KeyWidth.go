package keyvalue

/*
 * KeyValue.KeyWidth
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * 	Given a set of key-value data, return the maximum string length of all keys
 */

// KeyWidth - Return the maximum width of all keys in the current KeyValue struct
func (kv *KeyValue) KeyWidth() (width int) {
	if kv.data != nil {
		for key := range kv.data {
			if len(key) > width {
				width = len(key)
			}
		}
	}
	return width
}
