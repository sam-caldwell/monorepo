package keyvalue

import "sort"

// ToOrderedList - return the KeyValue data as a list of OrderedPairs (key,value)
func (kv *KeyValue) ToOrderedList(sorted bool) (result []OrderedPair) {
	if kv.data != nil {
		for key, value := range kv.data {
			result = append(result, OrderedPair{
				Key:   key,
				Value: value,
			})
		}
		if sorted {
			sort.Slice(result, func(i, j int) bool {
				return result[i].Key < result[j].Key
			})
		}
	}
	return result
}
