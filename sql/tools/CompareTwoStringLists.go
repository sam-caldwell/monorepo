package database

import (
	"reflect"
	"sort"
	"testing"
)

// CompareTwoStringLists - Compare two string lists
func CompareTwoStringLists(t *testing.T, lhs []string, rhs []string) {
	sort.Strings(lhs)
	sort.Strings(rhs)
	if !reflect.DeepEqual(lhs, rhs) {
		t.Fatalf("These two lists are not equal\nlhs: %v\nrhs: %v", lhs, rhs)
	}
}
