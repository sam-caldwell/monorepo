package persistence

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"testing"
)

func TestEventInit(t *testing.T) {
	const (
		badDir = "/reallyBadPathThatShouldNotExist"
		tmpDir = "/tmp"
	)

	var e EventCollector
	var dir directory.Path

	t.Run("initialize EventCollector (happy)", func(t *testing.T) {
		dir.Set(tmpDir)
		if err := e.Init(&dir); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("initialize EventCollector (sad)", func(t *testing.T) {
		dir.Set(badDir)
		if err := e.Init(&dir); err == nil {
			t.Fatal("expected error.  got none.")
		}
	})
}
