package keyvalue

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestKeyWidthHappyPath(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

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

	for i := 1; i <= 5; i++ {
		t.Run(fmt.Sprintf("Iteration %d", i), func(t *testing.T) {
			kv := KeyValue{data: Map{}}

			rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })

			for index, key := range keys {
				kv.data[key] = index
			}

			expectedWidth := len("this_key_is_the_longest")
			width := kv.KeyWidth()

			if width != expectedWidth {
				t.Errorf("Expected width %d, but got %d", expectedWidth, width)
			}
		})
	}
}

func TestKeyWidthWithEmptyMap(t *testing.T) {
	kv := KeyValue{
		data: Map{},
	}

	width := kv.KeyWidth()

	if width != 0 {
		t.Errorf("Expected width 0, but got %d", width)
	}
}

func TestKeyWidthWithNilMap(t *testing.T) {
	kv := KeyValue{}

	width := kv.KeyWidth()

	if width != 0 {
		t.Errorf("Expected width 0, but got %d", width)
	}
}
