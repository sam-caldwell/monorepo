package repolinter

import (
	"github.com/sam-caldwell/go/v2/projects/repotools"
	"os"
	"path/filepath"
)

func LinterMaster(pass func(name string) error, skip func(name, msg string) error,
	fail func(name string, err error) error) (err error) {

	var rootDirectory string //This is the root of the monorepo
	if rootDirectory, err = repotools.FindRepoRoot(); err != nil {
		return err
	}
	var linters = SetupLinter()

	return filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if skipIgnoredDirectories(path) {
			return nil
		}
		if info.IsDir() {
			//ToDo: validate the project directory names against a manifest.
			//ToDo: make sure directory names follow convention
			return nil
		}
		if err = linters.Run(path); err != nil {
			if err.Error() == noLinter {
				return skip(path, noLinter)
			} else {
				return fail(path, err)
			}
		}
		return pass(path)
	})
}
