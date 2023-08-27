package repotester

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/go/repotools"
	"github.com/sam-caldwell/go/v2/projects/go/repotools/manifest"
	repocli2 "github.com/sam-caldwell/go/v2/projects/go/repotools/ui"
	"io/fs"
	"path/filepath"
	"strings"
)

const (
	manifestFile = "MANIFEST.yaml"
)

func Setup(
	notice repocli2.NoticeMessagePrintFunc,
	pass repocli2.PassMessagePrintFunc,
	skip repocli2.SkipMessagePrintFunc,
	fail repocli2.FailMessagePrintFunc) func(projectType string) (err error) {

	testCommands := map[string]TestRunner{
		"c":      TestC,
		"cpp":    TestCpp,
		"go":     TestGolang,
		"nodejs": TestNodeJs,
	}

	// Run Tests - multi-language support
	runTest := func(projectPath string, manifest *projectmanifest.Manifest) (localErr error) {

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
			return fmt.Errorf("FindRepoRoot(): %v", err)
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
				if err = isDependencyEnabled(&rootDirectory, &manifest); err != nil {
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
