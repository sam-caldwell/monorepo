package pair

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"testing"
)

func TestKeyToBytes(t *testing.T) {
	t.Run("test with invalid (struct) type", func(t *testing.T) {
		type invalidType struct{}
		var lhs invalidType
		if _, err := KeyToBytes(lhs); err == nil {
			t.Fatalf("expected error for invalid type but got %v", err)
		} else {
			if err.Error() != errors.UnsupportedType {
				t.Fatalf("expected UnsupportedType error but got %v", err)
			}
		}
	})
	t.Run("test with boolean type", func(t *testing.T) {})
	t.Run("test with signed integer type", func(t *testing.T) {})
	t.Run("test with unsigned integer type", func(t *testing.T) {})
	t.Run("test with float32 type", func(t *testing.T) {})
	t.Run("test with float64 type", func(t *testing.T) {})
	t.Run("test with string type", func(t *testing.T) {})
	t.Run("test with rune type", func(t *testing.T) {})
	t.Run("test with struct pointer type", func(t *testing.T) {})
	t.Run("test with boolean pointer type", func(t *testing.T) {})
	t.Run("test with signed integer pointer type", func(t *testing.T) {})
	t.Run("test with unsigned integer pointer type", func(t *testing.T) {})
	t.Run("test with float32 pointer type", func(t *testing.T) {})
	t.Run("test with float64 pointer type", func(t *testing.T) {})
	t.Run("test with string pointer type", func(t *testing.T) {})
	t.Run("test with rune pointer type", func(t *testing.T) {})
}
