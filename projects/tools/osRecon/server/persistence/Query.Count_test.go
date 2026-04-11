package persistence

import "testing"

func TestQueryCount(t *testing.T) {
	var query QueryCollector
	if query.count != 0 {
		t.Fatal("expected 0")
	}
	if c := query.Count(); c != 0 {
		t.Fatal("expected 0")
	}
	if query.count = 1337; query.Count() != 1337 {
		t.Fatal("non-zero value expected")
	}
}
