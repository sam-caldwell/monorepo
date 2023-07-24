package repotester

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/repotools"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
	repocli "github.com/sam-caldwell/go/v2/projects/repotools/ui"
	"io/fs"
	"path/filepath"
	"strings"
)

const (
	manifestFile = "MANIFEST.yaml"
)

type TestRunner func(
	name *string,
	notice repocli.NoticeMessagePrintFunc,
	pass repocli.PassMessagePrintFunc,
	fail repocli.FailMessagePrintFunc,
	projectPath *string) error

func Setup(
	notice repocli.NoticeMessagePrintFunc,
	pass repocli.PassMessagePrintFunc,
	skip repocli.SkipMessagePrintFunc,
	fail repocli.FailMessagePrintFunc) func(projectType string) (err error) {

	// Run Tests - multi-language support
	runTest := func(projectPath string, manifest *projectmanifest.Manifest) (localErr error) {
		testCommands := map[string]TestRunner{
			"go": TestGolang,
		}
		thisLanguage := strings.ToLower(strings.TrimSpace(manifest.GetLanguage()))
		var command = func() TestRunner {
			if value, ok := testCommands[thisLanguage]; ok {
				return value
			}
			return nil
		}()
		if command == nil {
			return fmt.Errorf(errors.UnsupportedLanguage+errors.Details, thisLanguage)
		}
		return command(&manifest.Name, notice, pass, fail, &projectPath)
	}

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
			d := filepath.Dir(path)
			if manifest.IsTestEnabled() {
				if err = isDependencyEnabled(&projectDirectory, &manifest); err != nil {
					fail(manifest.Name, d, err)
					return err
				}
				err = runTest(d, &manifest)
			} else {
				skip(manifest.Name, path, "test not enabled")
			}
			return err
		})
		return err
	}
}
