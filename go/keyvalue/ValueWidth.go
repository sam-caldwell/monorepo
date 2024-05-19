package keyvalue

import (
	"fmt"
)

// ValueWidth - Return the maximum width of all values in the current KeyValue struct
//
//	(c) 2023 Sam Caldwell.  MIT License
func (kv *KeyValue[KeyType, ValueType]) ValueWidth() (width int) {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	if kv.data != nil {
		for _, value := range kv.data {
			if sz := len(fmt.Sprintf("%v", value)); sz > width {
				width = sz
			}
		}
	}
	return width
}
