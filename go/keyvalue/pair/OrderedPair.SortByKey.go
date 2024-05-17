package pair

import (
	"bytes"
	"sort"
)

// SortByKey - Sort OrderedPair by key.
func (o *OrderedPair[KeyType, ValueType]) SortByKey() {
	o.lock.Lock()
	defer o.lock.Unlock()
	sort.Slice(o.data, func(i int, j int) bool {
		var lhs []byte
		var rhs []byte
		var err error
		if lhs, err = KeyToBytes(o.data[i].Key); err != nil {
			panic(err)
		}
		if rhs, err = KeyToBytes(o.data[j].Key); err != nil {
			panic(err)
		}
		return bytes.Compare(lhs, rhs) < 0
	})
}
