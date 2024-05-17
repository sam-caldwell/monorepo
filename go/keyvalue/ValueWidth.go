package keyvalue

import "fmt"

/*
 * keyvalue.ValueWidth
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * 	Given a set of key-value data, return the maximum string length of all values
 */

// ValueWidth - Return the maximum width of all values in the current KeyValue struct
func (kv *KeyValue[KeyType, ValueType]) ValueWidth() (width int) {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	if kv.data != nil {
		for _, value := range kv.data {
			if sz := len(fmt.Sprintf("%v", value)); sz > width {
				return sz
			}
		}
	}
	return width
}
