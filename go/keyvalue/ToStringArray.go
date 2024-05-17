package keyvalue

/*
 * keyvalue.ToStringArray
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Given an internal key-value data set, this method will return a string
 * table of key-value data where each row is formatted using rowFormat.
 *
 * if pretty is true, we will pad the key column with spaces for the maximum
 * key width, creating pretty columns that make my OCD happy.  This does require
 * an additional pass over the data, so it's an option.
 *
 */

import (
	"fmt"
	"sort"
)

// ToStringArray - Return a list of strings where each row represents a key-value line
func (kv *KeyValue[KeyType, ValueType]) ToStringArray(columnDelimiter string, pretty bool) (output []string) {
	const rowFormat = "%*s%s%s"
	defer kv.lock.Unlock()
	if kv.data != nil {
		keyWidth := 0
		if pretty {
			keyWidth = kv.KeyWidth()
		}
		// Format each key-value pair with uniform columns
		kv.lock.Lock()
		for key, value := range kv.data {
			line := fmt.Sprintf(rowFormat, keyWidth, key, columnDelimiter, value)
			output = append(output, line)
		}
		kv.lock.Unlock()
	}
	sort.Strings(output)
	return output
}
