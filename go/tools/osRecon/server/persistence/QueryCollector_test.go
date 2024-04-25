package persistence

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/types"
	"testing"
)

func TestQueryCollector_Struct(t *testing.T) {
	var s types.Sequence
	var p directory.Path

	t.Run("Test initial state for the QueryCollector", func(t *testing.T) {
		var q QueryCollector

		if q.count != s {
			t.Fatal("initial state mismatch (sequence)")
		}

		if q.path.String() != p.String() {
			t.Fatal("initial state mismatch (path)")
		}

		if q.queue != nil {
			t.Fatal("initial state mismatch (queue)")
		}

	})

}
