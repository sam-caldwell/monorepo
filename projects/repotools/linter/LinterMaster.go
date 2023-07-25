package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/repotools"
	repocli "github.com/sam-caldwell/go/v2/projects/repotools/ui"
	"os"
	"path/filepath"
)

// LinterMaster - Walk through a directory and its contents and lint all objects.
func LinterMaster(
	useColor bool,
	notice repocli.NoticeMessagePrintFunc,
	pass repocli.PassMessagePrintFunc,
	skip repocli.SkipMessagePrintFunc,
	fail repocli.FailMessagePrintFunc) (err error) {

	var rootDirectory string //This is the root of the monorepo
	if rootDirectory, err = repotools.FindRepoRoot(); err != nil {
		return err
	}
	var linters = SetupLinter(useColor)

	reportResults := func(err error, linter, file string) {
		if err != nil {
			if err.Error() == noLinter || err.Error() == errors.NotImplemented {
				skip(file, linter, err.Error())
			} else {
				fail(file, linter, err)
				os.Exit(exit.GeneralError)
			}
		}
		pass(file, linter)
	}

	for extension, linter := range linters.Table {
		if !linter.enabled {
			notice("linter %s is disabled", linter.name)
			continue //skip disabled linters
		}
		notice("linter %s starting...", linter.name)
		if linter.directoryLevel { //run these from root directory
			source := filepath.Join(rootDirectory, "/...")
			notice("Source: %s", source)
			err = linters.Run(linter.runner, &source)
			reportResults(err, linter.name, source)
			if err != nil {
				return err
			}
		} else { //glob files with the given extension and lint individually.
			source := filepath.Join(rootDirectory, "**", "**", fmt.Sprintf("*.%s", extension))
			notice("Source: %s", source)
			files, err := filepath.Glob(source)
			if err != nil {
				return err
			}
			if len(files) == 0 {
				skip(source, linter.name, "no files found for "+extension)
			}
			notice("Found %d files", len(files))
			for _, file := range files {
				if skipIgnoredDirectories(file) {
					notice("Skipping %s (ignored file/directory)", file)
					continue //Skip
				}
				err = linters.Run(linter.runner, &file)
				reportResults(err, linter.name, file)
				if err != nil {
					return err
				}
			}
		}
		notice("linter %s completed...", linter.name)
	}
	return err
}
