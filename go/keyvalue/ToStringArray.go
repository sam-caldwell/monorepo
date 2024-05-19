package keyvalue

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"sort"
	"strings"
)

// ToStringArray - Return a list of strings where each row represents a key-value line
//
//			Given an internal key-value data set, this method will return a string
//			table of key-value data where each row is formatted using rowFormat.
//
//	     Each column is delimited with the given delimiter string (one or more characters).
//
//	     Each row is delimited by a lineDelimiter string (one or more characters).
//
//		    If pretty is true, we will pad the key column with spaces for the maximum
//		    key width, creating pretty columns that make my OCD happy.  This does require
//		    an additional pass over the data, so it's an option.
//
//		  (c) 2023 Sam Caldwell.  MIT License
func (kv *KeyValue[KeyType, ValueType]) ToStringArray(columnDelimiter, lineDelimiter string, pretty bool) []string {
	const rowFormat = "%v%s%v%v%s"
	var (
		keyWidth = 0
		output   []string
	)

	if kv.data == nil {
		return output
	}
	if pretty {
		keyWidth = kv.KeyWidth()
	}

	kv.lock.Lock()
	// Format each key-value pair with uniform columns
	for key, value := range kv.data {
		sz := keyWidth - len(fmt.Sprintf("%v", key))
		if sz < 0 {
			sz = 0
		}
		spacer := strings.Repeat(words.Space, sz)
		line := fmt.Sprintf(rowFormat, key, spacer, columnDelimiter, value, lineDelimiter)
		output = append(output, line)
	}
	kv.lock.Unlock()
	sort.Strings(output)
	return output
}
