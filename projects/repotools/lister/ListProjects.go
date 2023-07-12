package repotools

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/fs"
	"github.com/sam-caldwell/go/v2/projects/fs/directory"
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	"github.com/sam-caldwell/go/v2/projects/repotools/filters"
	"os"
	"path/filepath"
)

// ListProjects - List the enabled or disabled projects in the repo
func ListProjects(filter filters.Filter) (list map[string]Project, err error) {

	rootDirectory, err := findRepoRoot()
	if err != nil {
		return list, err
	}

	// Verify the path exists.
	if !directory.Existsp(&rootDirectory) {
		return list, fmt.Errorf(fs.ErrDoesNotExist, fs.Directory, rootDirectory)
	}
	list = make(map[string]Project)

	var projectTypes []string
	if filter.HasFilter(filters.ShowCommands) || filter.HasFilter(filters.ShowAny) {
		projectTypes = append(projectTypes, Command)
	}
	if filter.HasFilter(filters.ShowProjects) || filter.HasFilter(filters.ShowAny) {
		projectTypes = append(projectTypes, Package)
	}

	for _, projectType := range projectTypes {
		ansi.Green().Printf("Walking %s\n", projectType).Reset()
		// Walk the directory structure and find all the projects.
		err = filepath.Walk(projectType, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			pathParts := directory.SplitToList(path)
			manifestFile := filepath.Base(path)
			if (pathParts[0] != projectType) || (manifestFile != ManifestFile) || (len(pathParts) < 2) {
				return nil // Filter off the irrelevant
			}

			manifest, err := LoadProjectManifest(path)
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
			case Command:
				if !file.Exists(filepath.Join(thisPath, "main.go")) {
					return nil
				}
				list[fmt.Sprintf("%s::%s/%s", Command, pathParts[1], pathParts[2])] = Project{
					ProjectType: CommandProject,
					Project:     pathParts[1],
					Program:     pathParts[2],
					Path:        thisPath,
				}
			case Package:
				projectName := fmt.Sprintf("%s::%s", Package, pathParts[1])
				list[projectName] = Project{
					ProjectType: PackageProject,
					Project:     pathParts[1],
					Program:     pathParts[2],
					Path:        thisPath,
				}
			}
			return nil
		})
	}
	return list, err
}
