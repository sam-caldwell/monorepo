package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
	"path/filepath"
	"testing"
)

func TestMonorepo_LoadManifests(t *testing.T) {
	var repo Monorepo
	var original string
	var expectedCount int

	original = directory.GetCurrent()

	defer func() {
		_ = os.Chdir(original)
	}()

	t.Run("Get to repo root", func(t *testing.T) {
		repo.Root = original
		for {
			if filepath.Base(filepath.Dir(repo.Root)) == "git" {
				break
			}
			repo.Root = filepath.Dir(repo.Root)
		}
		repo.Root = filepath.Join(repo.Root)
		_ = os.Chdir(repo.Root)
	})
	t.Run("verify we are in the repo root", func(t *testing.T) {
		if repo.Root == original {
			t.Fatalf("repoRoot wrong: %s", repo.Root)
		}
		t.Logf("repoRoot: %s", repo.Root)
	})
	t.Run("load the monorepo", func(t *testing.T) {
		var err error
		repo.LoadManifests()
		if expectedCount, err = file.CountManifestFilesWalk(repo.Root, manifestYamlFile); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("verify counts", func(t *testing.T) {
		if expectedCount <= 1 {
			t.Fatal("ExpectedCount should be at least one")
		}
		if actualCount := len(repo.manifestList); actualCount != expectedCount {
			t.Fatalf("loaded manifest count mismatch\n"+
				"Actual: %d\n"+
				"Expect: %d", actualCount, expectedCount)
		}
	})
	t.Run("Verify paths exist", func(t *testing.T) {
		for n, manifest := range repo.manifestList {
			if !file.Exists(manifest.FileName) {
				t.Fatalf("missing file (%d): %s", n, manifest.FileName)
			}
		}
	})
}
