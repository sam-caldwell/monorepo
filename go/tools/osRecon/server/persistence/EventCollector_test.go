package persistence

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/types"
	"testing"
)

func TestEventCollectorStruct(t *testing.T) {
	var e EventCollector

	t.Run("EventCollector has a sequence (id)", func(t *testing.T) {

		var s types.Sequence

		if e.id != s {
			t.Fatal("initial sequence mismatch")
		}

	})

	t.Run("EventCollector has a directory.path (path)", func(t *testing.T) {

		var p directory.Path

		if e.path != p {
			t.Fatal("initial path mismatch")
		}
	})

}
