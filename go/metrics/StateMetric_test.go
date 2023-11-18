package metrics

import (
	"github.com/sam-caldwell/monorepo/go/types/generic"
	"github.com/sam-caldwell/monorepo/go/types/hashes"
	"testing"
)

func TestStateMetric(t *testing.T) {
	t.Run("[]byte test", func(t *testing.T) {
		var o StateMetric[[]byte]
		expected := []byte("test input")
		o.value = expected
	})
	t.Run("string test", func(t *testing.T) {
		var o StateMetric[string]
		expected := "test input"
		o.value = expected
	})
	t.Run("hashes.Sha1 test", func(t *testing.T) {
		var o StateMetric[hashes.Sha1]
		o.value = hashes.Sha1{}
		o.value.HashString("test input")
	})
	t.Run("hashes.Sha256 test", func(t *testing.T) {
		var o StateMetric[hashes.Sha256]
		o.value = hashes.Sha256{}
		o.value.HashString("test input")
	})
	t.Run("hashes.Sha512 test", func(t *testing.T) {
		var o StateMetric[hashes.Sha512]
		o.value = hashes.Sha512{}
		o.value.HashString("test input")
	})
	t.Run("generic.Block1KB test", func(t *testing.T) {
		var o StateMetric[generic.Block1KB]
		o.value.FromSlice([]byte("test input"))
	})
	t.Run("generic.Block2KB test", func(t *testing.T) {
		var o StateMetric[generic.Block2KB]
		o.value.FromSlice([]byte("test input"))
	})
	t.Run("generic.Block4KB test", func(t *testing.T) {
		var o StateMetric[generic.Block4KB]
		o.value.FromSlice([]byte("test input"))
	})
	t.Run("generic.Block8KB test", func(t *testing.T) {
		var o StateMetric[generic.Block8KB]
		o.value.FromSlice([]byte("test input"))
	})
	t.Run("generic.Block16KB test", func(t *testing.T) {
		var o StateMetric[generic.Block16KB]
		o.value.FromSlice([]byte("test input"))
	})
}
