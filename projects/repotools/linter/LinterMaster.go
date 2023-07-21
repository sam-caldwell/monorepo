package repolinter

import (
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/repotools"
	"os"
	"path/filepath"
)

// LinterMaster - Walk through a directory and its contents and lint all objects.
func LinterMaster(quiet bool,
	pass func(name string) error,
	skip func(quiet bool, name, msg string) error,
	fail func(name string, err error) error) (err error) {

	var rootDirectory string //This is the root of the monorepo
	if rootDirectory, err = repotools.FindRepoRoot(); err != nil {
		return err
	}
	var linters = SetupLinter()
	var hasFail = false
	return filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {
		if hasFail || err != nil {
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
			if err.Error() == noLinter || err.Error() == errors.NotImplemented {
				return skip(quiet, path, err.Error())
			} else {
				hasFail = true
				return fail(path, err)
			}
		}
		return pass(path)
	})
}
