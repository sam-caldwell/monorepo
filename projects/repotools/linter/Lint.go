package repolinter

import (
	"github.com/sam-caldwell/go/v2/projects/repotools"
	"os"
	"path/filepath"
)

func LinterMaster(pass func(name string), fail func(name string, err error)) (err error) {
	var rootDirectory string //This is the root of the monorepo
	if rootDirectory, err = repotools.FindRepoRoot(); err != nil {
		return err
	}
	var linters = SetupLinter()

	return filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			//ToDo: validate the project directory names against a manifest.
			//ToDo: make sure directory names follow convention
			return nil
		}
		if err = linters.Run(path); err != nil {
			fail(path, err)
			return err
		}
		pass(path)
		return nil
	})
}
