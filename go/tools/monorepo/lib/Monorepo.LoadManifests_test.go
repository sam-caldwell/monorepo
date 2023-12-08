package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
	"path/filepath"
	"strings"
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
	t.Run("Verify filenames exist", func(t *testing.T) {
		for n, manifest := range repo.manifestList {
			if strings.TrimSpace(manifest.FileName) == "" {
				t.Fatalf("unexpected empty filename at position %d", n)
			}
			if !file.Exists(manifest.FileName) {
				t.Fatalf("missing file (%d): %s", n, manifest.FileName)
			}
		}
	})
	t.Run("verify each manifest has a proper config", func(t *testing.T) {
		for n, manifest := range repo.manifestList {
			if manifest.config.Build.Enabled && !manifest.config.Build.Enabled {
				t.Fatalf("Build stage are bad %d on %s", n, manifest.FileName)
			}
			if manifest.config.Clean.Enabled && !manifest.config.Clean.Enabled {
				t.Fatalf("Clean stage are bad %d on %s", n, manifest.FileName)
			}
			if manifest.config.Test.Enabled && !manifest.config.Test.Enabled {
				t.Fatalf("Test stage are bad %d on %s", n, manifest.FileName)
			}
			if manifest.config.Header.Author == "" {
				t.Fatalf("Author should not be empty.  See %s", manifest.FileName)
			}
			if manifest.config.Header.Email == "" {
				t.Fatalf("Email should not be empty.  See %s", manifest.FileName)
			}
			if manifest.config.Header.Description == "" {
				t.Fatalf("Description should not be empty.  See %s", manifest.FileName)
			}
			if len(manifest.config.Header.Opsys) == 0 {
				t.Fatalf("Opsys should not be empty.  See %s", manifest.FileName)
			}
			if len(manifest.config.Header.Arch) == 0 {
				t.Fatalf("Arch should not be empty.  See %s", manifest.FileName)
			}

		}
	})

}
