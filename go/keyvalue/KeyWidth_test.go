package keyvalue

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestKeyValue_KeyWidth(t *testing.T) {
	keys := []string{
		"a",
		"ab",
		"abc",
		"abcd",
		"abcde",
		"short",
		"medium_size",
		"really_long_key",
		"even_longer_key",
		"this_key_is_the_longest",
	}
	t.Run("test KeyWidth() with happy path", func(t *testing.T) {
		for i := 1; i <= 5; i++ {
			t.Run(fmt.Sprintf("Iteration %d", i), func(t *testing.T) {
				var kv KeyValue[string, int]

				kv.data = make(map[string]int)

				rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })

				for index, key := range keys {
					kv.data[key] = index
				}

				expectedWidth := len("this_key_is_the_longest")
				width := kv.KeyWidth()

				if width != expectedWidth {
					t.Fatalf("Expected width %d, but got %d", expectedWidth, width)
				}
			})
		}
	})
	t.Run("test KeyWidth() with empty/nil map", func(t *testing.T) {
		var kv KeyValue[string, int]
		if width := kv.KeyWidth(); width != 0 {
			t.Fatalf("Expected width 0, but got %d", width)
		}
	})
}
