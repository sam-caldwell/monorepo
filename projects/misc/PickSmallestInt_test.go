package misc

import "testing"

func TestPickSmallestInt(t *testing.T) {
	t.Run("Picks smallest integer from multiple numbers", func(t *testing.T) {
		result := PickSmallestInt(5, 2, 9, 3)
		expected := 2
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("All numbers are equal, returns last number", func(t *testing.T) {
		result := PickSmallestInt(8, 8, 8, 8)
		expected := 8
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("Empty input, returns 0", func(t *testing.T) {
		result := PickSmallestInt()
		expected := 0
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})
}
