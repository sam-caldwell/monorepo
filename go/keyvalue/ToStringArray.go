package keyvalue

import (
	"fmt"
	"sort"
)

// ToStringArray - Return a list of strings where each row represents a key-value line
//
//		Given an internal key-value data set, this method will return a string
//		table of key-value data where each row is formatted using rowFormat.
//
//	    If pretty is true, we will pad the key column with spaces for the maximum
//	    key width, creating pretty columns that make my OCD happy.  This does require
//	    an additional pass over the data, so it's an option.
//
//	  (c) 2023 Sam Caldwell.  MIT License
func (kv *KeyValue[KeyType, ValueType]) ToStringArray(columnDelimiter string, pretty bool) (output []string) {
	const rowFormat = "%*v%v%v"
	var keyWidth = 0

	if kv.data == nil {
		return output
	}
	if pretty {
		keyWidth = kv.KeyWidth()
	}

	kv.lock.Lock()
	// Format each key-value pair with uniform columns
	for key, value := range kv.data {
		line := fmt.Sprintf(rowFormat, keyWidth, key, columnDelimiter, value)
		output = append(output, line)
	}
	kv.lock.Unlock()
	sort.Strings(output)
	return output
}
