package pair

import (
	"bytes"
	"sort"
)

// SortByValue - sort ordered pair by value
func (o *OrderedPair[KeyType, ValueType]) SortByValue() {
	o.lock.Lock()
	defer o.lock.Unlock()
	sort.Slice(o.data, func(i, j int) bool {
		var lhs []byte
		var rhs []byte
		var err error

		if lhs, err = ValueToBytes(o.data[i].Value); err != nil {
			panic(err)
		}
		if rhs, err = ValueToBytes(o.data[j].Value); err != nil {
			panic(err)
		}
		return bytes.Compare(lhs, rhs) < 0
	})
}
