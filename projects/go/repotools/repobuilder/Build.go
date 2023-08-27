package repobuilder

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/exit/errors"
	repotools2 "github.com/sam-caldwell/go/v2/projects/go/repotools"
	"github.com/sam-caldwell/go/v2/projects/go/repotools/manifest"
	repocli2 "github.com/sam-caldwell/go/v2/projects/go/repotools/ui"
	"io/fs"
	"path/filepath"
	"strings"
)

// Build - Run all builds for the repo or for a specific project path
func Build(notice repocli2.NoticeMessagePrintFunc,
	pass repocli2.PassMessagePrintFunc,
	skip repocli2.SkipMessagePrintFunc,
	fail repocli2.FailMessagePrintFunc) (err error) {

	const (
		projectType  = "cmd"
		manifestFile = "MANIFEST.yaml"
	)

	var rootDirectory string //This is the root of the monorepo
	if rootDirectory, err = repotools2.FindRepoRoot(); err != nil {
		return err
	}
	var projectRoot = filepath.Join(rootDirectory, projectType)
	var outputDirectory = filepath.Join(rootDirectory, "build")

	// Run Builds - multi-language support
	runBuild := func(projectPath string, manifest *projectmanifest.Manifest) (localErr error) {
		testCommands := map[string]BuildRunner{
			"c":      BuildC,
			"cpp":    BuildCpp,
			"go":     BuildGolang,
			"nodejs": BuildNodeJs,
		}
		thisLanguage := strings.ToLower(strings.TrimSpace(manifest.GetLanguage()))
		var command = func() BuildRunner {
			if value, ok := testCommands[thisLanguage]; ok {
				return value
			}
			return nil
		}()
		if command == nil {
			return fmt.Errorf(errors.UnsupportedLanguage+errors.Details, thisLanguage)
		}
		return command(&manifest.Name, &outputDirectory, notice, pass, fail, &projectPath)
	}
	err = repotools2.Clean()
	if err != nil {
		return fmt.Errorf("builder failed to clean: %v", err)
	}
	err = filepath.Walk(projectRoot, func(path string, info fs.FileInfo, thisError error) error {
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
		projectDirectory := filepath.Dir(path)
		if manifest.IsTestEnabled() {
			if err = isDependencyEnabled(&rootDirectory, &manifest); err != nil {
				fail(manifest.Name, projectDirectory, err)
				return err
			}
			err = runBuild(projectDirectory, &manifest)
		} else {
			skip(manifest.Name, path, "build not enabled")
		}
		return err
	})
	return err
}
