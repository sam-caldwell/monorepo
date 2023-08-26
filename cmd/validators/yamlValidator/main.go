package main

/*
 * YamlValidator
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This program will glob a path and its subdirectories and validate any
 * files therein which have the .yml or .yaml file extension.
 *
 * For testing: find ./ -name "*.yml" -o -name "*.yaml" -type f
 */
import (
	"fmt"
	ansi "github.com/sam-caldwell/go/v2/projects/go/ansi"
	"github.com/sam-caldwell/go/v2/projects/go/exit"
	"github.com/sam-caldwell/go/v2/projects/go/fs"
	"github.com/sam-caldwell/go/v2/projects/go/fs/directory"
	"github.com/sam-caldwell/go/v2/projects/go/fs/file"
	"github.com/sam-caldwell/go/v2/projects/go/misc/words"
	"os"
	"path/filepath"
)

const (
	errAccessDenied = "Cannot access '%s': %v"

	cmdUsage = "\n\nUsage:\nyamlValidator <path>"
)

func main() {
	exit.IfVersionRequested()
	// Check if the directory path is provided as an argument
	exit.OnCondition(
		len(os.Args) < 2,
		exit.MissingArg,
		fmt.Sprintf(fs.ErrNotFound, fs.Directory),
		cmdUsage)

	dirPath := os.Args[1]

	// Verify the path exists.
	exit.OnCondition(
		!directory.Existsp(&dirPath),
		exit.NotFound,
		fmt.Sprintf(fs.ErrDoesNotExist, fs.Directory, dirPath),
		cmdUsage)

	// Find and validate YAML files
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			ansi.Red().Printf(errAccessDenied, path, err).LF().Reset()
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if file.HasYamlExtension(path) {
			ansi.Blue().Printf("%v", path).Reset().Space().Print(":")
			if err = file.IsValidYaml(path); err != nil {
				ansi.Red().Printf("FAIL: %v", err).LF().Reset()
			} else {
				ansi.Green().Print("OK").LF().Reset()
			}
		}
		return nil
	})
	exit.OnError(err, exit.GeneralError, words.EmptyString)
}
