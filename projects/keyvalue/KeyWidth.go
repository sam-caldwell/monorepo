package keyvalue

import (
	"fmt"
)

// KeyWidth - Return the maximum width of all keys in the current KeyValue struct
//
//	(c) 2023 Sam Caldwell.  MIT License
func (kv *KeyValue[KeyType, ValueType]) KeyWidth() (width int) {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	if kv.data != nil {
		for key := range kv.data {
			if sz := len(fmt.Sprintf("%v", key)); sz > width {
				width = sz
			}
		}
	}
	return width
}
