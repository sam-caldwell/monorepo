package repotools

/*
 * projects/repotool/VerifyRepoIntegrity.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines a function called VerifyRepoIntegrity()
 * which will verify that the given 'repoRoot' directory contains
 * the expected minimum directories and files for this monorepo
 * and is therefore in minimum healthy state.
 */

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/fs/directory"
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	"path/filepath"
)

func VerifyRepoIntegrity(repoRoot string) error {
	expectedDirectories := []string{
		"cmd",
		"cmd/tools",
		"projects",
		"projects/__system__",
	}
	expectedFiles := []string{
		"LICENSE.txt",
		"SECURITY.md",
		"README.md",
		"go.mod",
		"go.sum",
		".gitignore",
	}

	for _, d := range expectedDirectories {
		thisDirectory := filepath.Join(repoRoot, d)
		if !directory.Exists(thisDirectory) {
			return fmt.Errorf("repository root does not contain expected directory (%s):%s",
				repoRoot, thisDirectory)
		}
	}
	for _, f := range expectedFiles {
		thisFile := filepath.Join(repoRoot, f)
		if !file.Exists(thisFile) {
			return fmt.Errorf("repository root does not contain expected file (%s):%s",
				repoRoot, thisFile)
		}
	}
	return nil
}
