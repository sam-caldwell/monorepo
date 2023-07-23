package repotester

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/repotools"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
	repocli "github.com/sam-caldwell/go/v2/projects/repotools/ui"
	"github.com/sam-caldwell/go/v2/projects/runcommand"
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

		projectDirectory := filepath.Join(rootDirectory, projectType)

		err = filepath.Walk(projectDirectory, func(path string, info fs.FileInfo, thisError error) error {
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
			if manifest.IsTestEnabled() {
				d := filepath.Dir(path)
				command := fmt.Sprintf("go test -failfast -v -count=2 -vet=all %s/...", d)
				out, err := runcommand.ShellExecute(command)
				if err == nil {
					pass(manifest.Name, path)
				} else {
					err = fmt.Errorf("out:%s\nerr:%s", out, err)
					fail(manifest.Name, path, err)
				}
			} else {
				skip(manifest.Name, path, "test not enabled")
			}
			return err
		})
		return err
	}
}
