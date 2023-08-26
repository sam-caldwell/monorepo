package convert

import (
	"testing"
)

func TestAnyListToStringList(t *testing.T) {
	input := []any{
		"A",
		1,
		-13.01,
		"true",
	}
	expected := []string{
		"A",
		"1",
		"-13.01",
		"true",
	}
	actual := AnyListToStringList(&input)
	for i, a := range actual {
		if a != expected[i] {
			t.Fatal(i)
		}
	}
}
