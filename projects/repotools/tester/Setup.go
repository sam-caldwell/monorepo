package repotester

import (
	"github.com/sam-caldwell/go/v2/projects/repotools"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
	repocli "github.com/sam-caldwell/go/v2/projects/repotools/ui"
	"io/fs"
	"path/filepath"
)

const (
	manifestFile = "MANIFEST.yaml"
)

func Setup(
	notice repocli.NoticeMessagePrintFunc,
	pass repocli.PassMessagePrintFunc,
	skip repocli.SkipMessagePrintFunc,
	fail repocli.FailMessagePrintFunc) func(projectType string) (err error) {

	// Run - Run all tests for the repo or specific project
	return func(projectType string) (err error) {

		var rootDirectory string //This is the root of the monorepo
		if rootDirectory, err = repotools.FindRepoRoot(); err != nil {
			return err
		}

		err = filepath.Walk(rootDirectory, func(path string, info fs.FileInfo, thisError error) error {
			if thisError != nil {
				err = thisError
				return err
			}
			if filepath.Base(path) != manifestFile {
				return nil
			}
			var manifest projectmanifest.Manifest
			err = manifest.LoadFile(path)
			if err != nil {
				return err
			}
			if !manifest.IsTestEnabled() {
				return nil
			}

			//ToDo: Run test for the project

			return nil
		})
		return err
	}
}
