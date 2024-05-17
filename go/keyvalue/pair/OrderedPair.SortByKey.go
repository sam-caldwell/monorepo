package pair

import (
	"bytes"
	"github.com/sam-caldwell/monorepo/go/keyvalue"
	"sort"
)

// SortByKey - Sort OrderedPair by key.
func (o *OrderedPair[KeyType, ValueType]) SortByKey() {
	sort.Slice(o.data, func(i int, j int) bool {
		var lhs []byte
		var rhs []byte
		var err error
		if lhs, err = keyvalue.KeyToBytes(o.data[i].Key); err != nil {
			panic(err)
		}
		if rhs, err = keyvalue.KeyToBytes(o.data[j].Key); err != nil {
			panic(err)
		}
		return bytes.Compare(lhs, rhs) < 0
	})
}
