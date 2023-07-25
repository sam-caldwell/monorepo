package repotools

import (
	"os"
	"path/filepath"
)

// Clean - Clean any build/test state
func Clean() (err error) {
	var rootDirectory string //This is the root of the monorepo
	if rootDirectory, err = FindRepoRoot(); err != nil {
		return err
	}
	var outputDirectory = filepath.Join(rootDirectory, "build")
	return os.RemoveAll(outputDirectory)
}
