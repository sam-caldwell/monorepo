package sqldbtest

import (
	"sort"
	"strings"
	"testing"
)

// CompareTwoStringLists - Compare two string lists
func CompareTwoStringLists(t *testing.T, lhs []string, rhs []string) {
	sort.Strings(lhs)
	sort.Strings(rhs)
	for i := 0; i < len(lhs); i++ {
		if i >= len(rhs) {
			t.Fatalf("Fail: missing column.\n"+
				"%d %v", i, lhs[i])
		}
		if strings.ToLower(lhs[i]) != strings.ToLower(rhs[i]) {
			t.Fatalf("Fail: These two lists are not equal\n"+
				"lhs: %v\n"+
				"rhs: %v",
				strings.ToLower(lhs[i]),
				strings.ToLower(rhs[i]))
		}
	}
}
