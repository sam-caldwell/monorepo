package listrepoprojects

/*
 * projects/repotools/listrepoprojects/ListProjects.go
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
	keyvalue "github.com/sam-caldwell/go/v2/projects/KeyValue"
	"github.com/sam-caldwell/go/v2/projects/fs/directory"
	"github.com/sam-caldwell/go/v2/projects/repotools"
	"github.com/sam-caldwell/go/v2/projects/repotools/filters"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
	"os"
	"path/filepath"
)

// ListProjects - List the enabled or disabled projects in the repo
func ListProjects(filter filters.Filter) (list keyvalue.KeyValue, err error) {
	/*
	* find the root of the repository.  This is where .git/ exists, but more importantly
	* this is the root from which either commands (cmd/<project_name>/<program_name>) and
	* reusable code packages (projects/<project_name>) can be found.
	 */
	var rootDirectory string //This is the root of the monorepo
	if rootDirectory, err = repotools.FindRepoRoot(); err != nil {
		return list, err
	}
	/*
	* Given our -commands os.Args which should have translated to a
	* filter.Commands state (bool), let's determine which type of project
	* and source directory we are going to list.
	 */
	projectDirectory, minNumberOfParts := determineProjectDirectory(rootDirectory, filter.Commands)
	/*
	* Initialize our map structure
	 */
	list.Initialize(0, true)
	/*
	* Start walking through every file and directory in <projectDirectory>
	 */
	err = filepath.Walk(projectDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil //Bail on the directories...we want projects.  We bail here to bail fast.
		}
		/*
		 * We need a filename because we are looking for MANIFEST.yaml
		 */
		fileName := filepath.Base(path)
		if fileName != projectmanifest.ManifestYaml {
			return nil // Filter off the irrelevant
		}
		/*
		 * Split our path into its parts by the path separator (/ in linux \ in windows)
		 * and verify we have at least the minimum number of parts.
		 */
		pathParts := directory.SplitToList(path)
		if len(pathParts) < minNumberOfParts {
			/*
			 * at minimum we should have either
			 *     cmd/<project>/<program>/MANIFEST.yaml (four elements)
			 * or
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

// determineProjectDirectory - if determine which types of projects we will list.
func determineProjectDirectory(rootDirectory string, showCommands bool) (projectType string, numberOfParts int) {
	if showCommands {
		projectType, numberOfParts = "cmd", 4
	} else {
		projectType, numberOfParts = "projects", 3
	}
	return filepath.Join(rootDirectory, projectType), numberOfParts
}
