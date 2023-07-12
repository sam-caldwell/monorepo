package repotools

/*
 * projects/repotool/FindRepoRoot.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines LoadFile() which find the root
 * of the repository from anywhere in the repository.
 */

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/fs/directory"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
	"github.com/sam-caldwell/go/v2/projects/wrappers/os"
	"path/filepath"
)

// FindRepoRoot - From a given directory in the repo walk back to the parent directories until we find .git/
func FindRepoRoot() (rootDirectory string, err error) {
	const gitDirectory = ".git"
	var homeDir string

	if homeDir, err = systemrecon.GetCurrentHomeDir(); err != nil {
		return rootDirectory, err
	}

	if rootDirectory, err = os.Getwd(); err != nil {
		return rootDirectory, err
	}

	for rootDirectory != homeDir {

		if directory.Exists(filepath.Join(rootDirectory, gitDirectory)) {
			return rootDirectory, nil //Found it!
		}

		rootDirectory = filepath.Dir(rootDirectory)

	}
	// Not found...we failed.  Go home and cry.  Something's wrong.
	return rootDirectory, fmt.Errorf(errors.NotFound)
}
