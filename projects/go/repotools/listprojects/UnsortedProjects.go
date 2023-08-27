package listprojects

/*
 * projects/repotools/listprojects/UnsortedProjects.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the ListProjects() function which will
 * apply our filters and manifests to list a set of projects
 * as a map of file paths and project attributes.
 *
 * See README.md
 */

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/fs/directory"
	"github.com/sam-caldwell/go/v2/projects/go/keyvalue"
	"github.com/sam-caldwell/go/v2/projects/go/repotools"
	"github.com/sam-caldwell/go/v2/projects/go/repotools/filters"
	projectmanifest2 "github.com/sam-caldwell/go/v2/projects/go/repotools/manifest"
	"os"
	"path/filepath"
)

// UnsortedProjects - List the enabled or disabled projects in the repo
func UnsortedProjects(filter filters.Filter) (list keyvalue.KeyValue, err error) {
	var rootDirectory string //This is the root of the monorepo
	if rootDirectory, err = repotools.FindRepoRoot(); err != nil {
		return list, err
	}
	projectDirectory, minNumberOfParts := determineProjectDirectory(rootDirectory, filter.Commands)
	list.Initialize(0, true)

	err = filepath.Walk(projectDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil //Bail on the directories...we want projects.  We bail here to bail fast.
		}
		fileName := filepath.Base(path)
		if fileName != projectmanifest2.ManifestYaml {
			return nil // Filter off the irrelevant
		}

		pathParts := directory.SplitToList(path)
		if len(pathParts) < minNumberOfParts {
			return fmt.Errorf("a path with a %s file should have at least %d parts (%s)",
				projectmanifest2.ManifestYaml, minNumberOfParts, path)
		}

		var manifest projectmanifest2.Manifest
		if err = manifest.LoadFile(path); err != nil {
			return err
		}
		if !filter.Apply(manifest) {
			return nil //skip filtered project
		}
		list.SetInterface(filepath.Dir(path), manifest)
		return nil
	})
	return list, err
}
