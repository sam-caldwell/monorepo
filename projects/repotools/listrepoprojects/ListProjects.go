package listrepoprojects

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
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/fs"
	"github.com/sam-caldwell/go/v2/projects/fs/directory"
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	"github.com/sam-caldwell/go/v2/projects/repotools"
	"github.com/sam-caldwell/go/v2/projects/repotools/filters"
	repoLinter "github.com/sam-caldwell/go/v2/projects/repotools/linter"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
	"os"
	"path/filepath"
)

// ListProjects - List the enabled or disabled projects in the repo
func ListProjects(filter filters.Filter) (list map[string]projectmanifest.Manifest, err error) {
	var rootDirectory string //This is the root of the monorepo

	/*
	 * find the root of the repository.  This is where .git/ exists, but more importantly
	 * this is the root from which either commands (cmd/<project_name>/<program_name>) and
	 * reusable code packages (projects/<project_name>) can be found.
	 */
	if rootDirectory, err = repotools.FindRepoRoot(); err != nil {
		return list, err
	}
	if !directory.Existsp(&rootDirectory) {
		return list, fmt.Errorf(fs.ErrDoesNotExist, fs.Directory, rootDirectory)
	}

	list = make(map[string]projectmanifest.Manifest)

	var projectType string
	if filter.Commands {
		projectType = "cmd"
	} else {
		projectType = "projects"
	}

	err = filepath.Walk(projectType, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		pathParts := directory.SplitToList(path)
		manifestFile := filepath.Base(path)
		if (pathParts[0] != projectType) || (manifestFile != repotools.ManifestFile) || (len(pathParts) < 2) {
			return nil // Filter off the irrelevant
		}

		manifest, err := repoLinter.LoadProjectManifest(path)
		if err != nil {
			return err
		}

		if filter.ApplyFilter(projectType, manifest) {
			ansi.Yellow().Printf("filtered: '%s'", path).LF().Reset()
			return nil //skip filtered project
		}
		ansi.Green().Printf("allowed:  '%s'", path).LF().Reset()

		thisPath := filepath.Base(path)
		switch projectType {
		case repotools.Command:
			if !file.Exists(filepath.Join(thisPath, "main.go")) {
				return nil
			}
			list[fmt.Sprintf("%s::%s/%s", repotools.Command, pathParts[1], pathParts[2])] = repotools.Project{
				ProjectType: repotools.CommandProject,
				Project:     pathParts[1],
				Program:     pathParts[2],
				Path:        thisPath,
			}
		case repotools.Package:
			projectName := fmt.Sprintf("%s::%s", repotools.Package, pathParts[1])
			list[projectName] = repotools.Project{
				ProjectType: repotools.PackageProject,
				Project:     pathParts[1],
				Program:     pathParts[2],
				Path:        thisPath,
			}
		}
		return nil
	})
	return list, err
}
