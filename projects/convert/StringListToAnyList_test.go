package convert

import "testing"

func TestStringListToAnyList(t *testing.T) {
	input := []string{
		"A",
		"B",
		"C",
		"D",
	}
	expected := []any{
		"A",
		"B",
		"C",
		"D",
	}
	actual := StringListToAnyList(&input)
	for i, a := range actual {
		if a.(string) != expected[i].(string) {
			t.Fatal(i)
		}
	}
}
