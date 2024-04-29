package persistence

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"testing"
)

func TestQueryCollector_List(t *testing.T) {
	t.Skip("debugging")
	const (
		tmpDir   = "/tmp"
		clientId = "77143d5a-b69f-47a7-8b64-318fa1e82cc3"
		queryDir = tmpDir +
			"/queries/" +
			clientId
	)
	var (
		query QueryCollector
	)

	t.Cleanup(func() {
		if err := directory.Delete(queryDir, directory.RecursiveDelete); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create a query directory", func(t *testing.T) {
		if err := directory.Create(queryDir, 0600); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Initialize QueryCollector", func(t *testing.T) {
		p := directory.Path(queryDir)
		if err := query.Init(&p); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("count queries in directory, expect 0", func(t *testing.T) {
		if count := query.Count(); count != 0 {
			t.Fatalf("count expected 0 but got %d", count)
		}
	})

	t.Run("Create query file", func(t *testing.T) {

	})

	t.Run("Verify the query file list", func(t *testing.T) {

	})
	t.Run("Read the first query file and verify the query", func(t *testing.T) {

	})
}
