package keyvalue

import (
	"github.com/sam-caldwell/monorepo/go/keyvalue/pair"
)

// ToOrderedPairList - Convert KeyValue to OrderedPair
func (kv *KeyValue[KeyType, ValueType]) ToOrderedPairList(sorted bool) (list pair.OrderedPair[KeyType, ValueType]) {
	kv.lock.Lock()
	defer kv.lock.Unlock()
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
