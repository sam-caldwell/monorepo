package keyvalue

import (
	"github.com/sam-caldwell/go/v2/projects/KeyValue/pair"
)

func (kv *KeyValue) ToOrderedPairList(sorted bool) (list pair.OrderedPair) {
	if kv.data == nil {
		return list
	}
	for key, value := range kv.data {
		list.Add(key, value)
	}
	if sorted {
		list.SortByKey()
	}
	return list
}
