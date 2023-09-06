package pair

import (
	"sort"
)

func (o *OrderedPair) SortByKey() {
	sort.Slice(o.data, func(i int, j int) bool {
		return o.data[i].Key < o.data[j].Key
	})
}
