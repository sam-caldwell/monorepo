package persistence

import (
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"testing"
)

func TestQueryCollector_Init(t *testing.T) {
	const (
		tmpDir = "/tmp"
	)

	var q QueryCollector

	t.Run("Initialize the QueryCollector", func(t *testing.T) {
		var (
			err error

			p *directory.Path
		)
		if p, err = directory.NewPath(tmpDir, true); err != nil {
			t.Fatal(err)
		}
		if err := q.Init(p); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("verify QueryCollector state", func(t *testing.T) {
		if q.count != 0 {
			t.Fatal("count mismatch")
		}
		if q.path.String() != tmpDir {
			t.Fatal("path mismatch")
		}
		if q.queue == nil {
			t.Fatal("queue is not initialized")
		}
		var err error
		var id uuid.UUID
		if id, err = uuid.NewRandom(); err != nil {
			t.Fatal("uuid generator failed")
		}
		q.queue[id] = []string{"test value"}
		if v, ok := q.queue[id]; ok {
			if v[0] != "test value" {
				t.Fatal("queue failed to store value")
			}
		}
	})

}
