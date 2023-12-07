package metrics

import (
	"bytes"
	"github.com/sam-caldwell/monorepo/go/types/generic"
	"github.com/sam-caldwell/monorepo/go/types/hashes"
	"strings"
	"testing"
)

func TestStateMetric_Set(t *testing.T) {
	t.Run("StateMetric.Set([]byte) test", func(t *testing.T) {
		var o StateMetric[[]byte]
		expected := []byte("test input")
		o.Set(expected)
		o.value = expected
		if v := o.Get(); !bytes.Equal(v, expected) {
			t.Fatal("mismatch")
		}
	})
	t.Run("StateMetric.Set(string) test", func(t *testing.T) {
		var o StateMetric[string]
		expected := "test input"
		o.Set(expected)
		o.value = expected
		if v := o.Get(); v != expected {
			t.Fatal("mismatch")
		}
	})
	t.Run("StateMetric.Set(hashes.Sha1) test", func(t *testing.T) {
		var o StateMetric[hashes.Sha1]
		testInput := "test input"
		expected := "49883b34e5a0f48224dd6230f471e9dc1bdbeaf5"
		var hash hashes.Sha1
		hash.HashString(testInput)
		o.Set(hash)
		o.value.HashString(testInput)
		if actual := o.Get(); actual.String() != expected {
			t.Fatalf("mismatch\n"+
				"actual:  %v\n"+
				"expected:%v",
				actual.String(), expected)
		}
	})
	t.Run("StateMetric.Set(hashes.Sha256) test", func(t *testing.T) {
		var o StateMetric[hashes.Sha256]
		testInput := "test input"
		expected := "9dfe6f15d1ab73af898739394fd22fd72a03db01834582f24bb2e1c66c7aaeae"
		var hash hashes.Sha256
		hash.HashString(testInput)
		o.Set(hash)
		o.value.HashString(testInput)
		if actual := o.Get(); actual.String() != expected {
			t.Fatalf("mismatch\n"+
				"actual:  %v\n"+
				"expected:%v",
				actual.String(), expected)
		}
	})
	t.Run("StateMetric.Set(hashes.Sha512) test", func(t *testing.T) {
		var o StateMetric[hashes.Sha512]
		testInput := "test input"
		expected := "40aa1b203c9d8ee150b21c3c7cda8261492e5420c5f2b9f7380700e094c303b48e62" +
			"f319c1da0e32eb40d113c5f1749cc61aeb499167890ab82f2cc9bb706971"
		var hash hashes.Sha512
		hash.HashString(testInput)
		o.Set(hash)
		o.value.HashString(testInput)
		if actual := o.Get(); actual.String() != expected {
			t.Fatalf("mismatch\n"+
				"actual:  %v\n"+
				"expected:%v",
				actual.String(), expected)
		}
	})
	t.Run("StateMetric.Set(generic.Block1KB) test", func(t *testing.T) {
		var o StateMetric[generic.Block1KB]
		testInput := "test input"
		expected := []byte(testInput + strings.Repeat(string([]byte{0x00}), len(o.value)-len(testInput)))
		var block generic.Block1KB
		block.FromSlice(expected)
		o.Set(block)
		if actual := o.Get(); !bytes.Equal(actual.ToSlice(), expected) {
			t.Fatalf("mismatch\n"+
				"actual:  %v\n"+
				"expected:%v",
				actual.ToSlice(), expected)
		}
	})
	t.Run("StateMetric.Set(generic.Block2KB) test", func(t *testing.T) {
		var o StateMetric[generic.Block2KB]
		testInput := "test input"
		expected := []byte(testInput + strings.Repeat(string([]byte{0x00}), len(o.value)-len(testInput)))
		var block generic.Block2KB
		block.FromSlice(expected)
		o.Set(block)
		if actual := o.Get(); !bytes.Equal(actual.ToSlice(), expected) {
			t.Fatalf("mismatch\n"+
				"actual:  %v\n"+
				"expected:%v",
				actual.ToSlice(), expected)
		}
	})
	t.Run("StateMetric.Set(generic.Block4KB) test", func(t *testing.T) {
		var o StateMetric[generic.Block4KB]
		testInput := "test input"
		expected := []byte(testInput + strings.Repeat(string([]byte{0x00}), len(o.value)-len(testInput)))
		var block generic.Block4KB
		block.FromSlice(expected)
		o.Set(block)
		if actual := o.Get(); !bytes.Equal(actual.ToSlice(), expected) {
			t.Fatalf("mismatch\n"+
				"actual:  %v\n"+
				"expected:%v",
				actual.ToSlice(), expected)
		}
	})
	t.Run("StateMetric.Set(generic.Block8KB) test", func(t *testing.T) {
		var o StateMetric[generic.Block8KB]
		testInput := "test input"
		expected := []byte(testInput + strings.Repeat(string([]byte{0x00}), len(o.value)-len(testInput)))
		var block generic.Block8KB
		block.FromSlice(expected)
		o.Set(block)
		if actual := o.Get(); !bytes.Equal(actual.ToSlice(), expected) {
			t.Fatalf("mismatch\n"+
				"actual:  %v\n"+
				"expected:%v",
				actual.ToSlice(), expected)
		}
	})
	t.Run("StateMetric.Set(generic.Block16KB) test", func(t *testing.T) {
		var o StateMetric[generic.Block16KB]
		testInput := "test input"
		expected := []byte(testInput + strings.Repeat(string([]byte{0x00}), len(o.value)-len(testInput)))
		var block generic.Block16KB
		block.FromSlice(expected)
		o.Set(block)
		if actual := o.Get(); !bytes.Equal(actual.ToSlice(), expected) {
			t.Fatalf("mismatch\n"+
				"actual:  %v\n"+
				"expected:%v",
				actual.ToSlice(), expected)
		}
	})
}
