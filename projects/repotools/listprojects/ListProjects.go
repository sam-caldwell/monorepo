package listprojects

/*
 * projects/repotools/listprojects/ListProjects.go
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
	"github.com/sam-caldwell/go/v2/projects/fs/directory"
	keyvalue "github.com/sam-caldwell/go/v2/projects/keyvalue"
	"github.com/sam-caldwell/go/v2/projects/repotools"
	"github.com/sam-caldwell/go/v2/projects/repotools/filters"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
	"os"
	"path/filepath"
)

// ListProjects - List the enabled or disabled projects in the repo
func ListProjects(filter filters.Filter) (list keyvalue.KeyValue, err error) {

	list.Initialize(0, true)

	var rootDirectory string //This is the root of the monorepo
	if rootDirectory, err = repotools.FindRepoRoot(); err != nil {
		return list, err
	}

	projectDirectory, minNumberOfParts := determineProjectDirectory(rootDirectory, filter.Commands)

	err = filepath.Walk(projectDirectory, func(path string, info os.FileInfo, err error) error {
		fileName := filepath.Base(path)
		if (err != nil) || info.IsDir() || (fileName != projectmanifest.ManifestYaml) {
			// bail on irrelevant stuff.
			return err
		}

		pathParts := directory.SplitToList(path)
		if len(pathParts) < minNumberOfParts {
			/* ...at minimum we should have either
			 *     cmd/<project>/<program>/MANIFEST.yaml (four elements)
			 * ...or...
			 *     projects/<project>/MANIFEST.yaml (three elements)
			 */
			return fmt.Errorf("a path with a %s file should have at least %d parts (%s)",
				projectmanifest.ManifestYaml, minNumberOfParts, path)
		}
		/*
		 * Load the project manifest and apply the filters.
		 */
		var manifest projectmanifest.Manifest
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
