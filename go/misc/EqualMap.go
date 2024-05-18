package misc

import "reflect"

// EqualMap - Compare two maps and return boolean indicating equivalence
func EqualMap[KeyType comparable, ValueType any](lhs, rhs *map[KeyType]ValueType) bool {
	if lhs == nil || rhs == nil {
		if lhs == nil && rhs == nil {
			return true
		}
		return false
	}
	for key, lValue := range *lhs {
		if rValue, ok := (*rhs)[key]; !ok || !reflect.DeepEqual(lValue, rValue) {
			return false
		}
	}
	for key, rValue := range *rhs {
		if lValue, ok := (*lhs)[key]; !ok || !reflect.DeepEqual(lValue, rValue) {
			return false
		}
	}
	return true
}
