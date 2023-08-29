package keyvalue

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/misc/words"
	"reflect"
)

/*
 * keyvalue.FindValue()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Return boolean result and the name of the key where a given value
 * is found.
 *
 * Warning: This is a linear search.
 */

// FindValue - Return the key where a given value exists in the key-value store and boolean true if found.
func (kv *KeyValue) FindValue(target any) (key any, found bool) {
	if kv.data != nil {
		for key, value := range kv.data {
			if reflect.DeepEqual(value, target) {
				return key, true
			}
		}
	}
	return words.EmptyString, found
}
