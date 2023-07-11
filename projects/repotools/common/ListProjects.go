package repotools

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/fs"
	"github.com/sam-caldwell/go/v2/projects/fs/directory"
	"github.com/sam-caldwell/go/v2/projects/repotools/common/projectFilter"
	"github.com/sam-caldwell/go/v2/projects/wrappers/os"
	"path/filepath"
)

// ListProjects - List the enabled or disabled projects in the repo
func ListProjects(filter projectFilter.Filter) (list map[string]Project, err error) {

	rootDirectory, err := findRepoRoot()
	if err != nil {
		return list, err
	}

	// Verify the path exists.
	if !directory.Existsp(&rootDirectory) {
		return list, fmt.Errorf(fs.ErrDoesNotExist, fs.Directory, rootDirectory)
	}

	list = make(map[string]Project)

	for _, codeDirectory := range []string{Command, Package} {
		// Walk the directory structure and find all the projects.
		err = filepath.Walk(codeDirectory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			pathParts := directory.SplitToList(path)
			if len(pathParts) < 2 {
				return nil
			}

			switch pathParts[0] {
			case Command:
				if filter.HasFilter(projectFilter.Command) {
					if filepath.Base(path) == "main.go" {
						if projectFilter.FilterThisPath(filter, filepath.Dir(path)) {
							//We are filtering this path...nothing else to see here.  Move along people.
							return nil
						}
						projectName := fmt.Sprintf("%s::%s/%s", Command, pathParts[1], pathParts[2])
						list[projectName] = Project{
							projectType: CommandProject,
							project:     pathParts[1],
							program:     pathParts[2],
							pathToMain:  path,
						}
					}
				}
			case Package:
				if filter.HasFilter(projectFilter.Package) {
					if projectFilter.FilterThisPath(filter, filepath.Dir(path)) {
						//We are filtering this path...nothing else to see here.  Move along people.
						return nil
					}
					projectName := fmt.Sprintf("%s::%s", Package, pathParts[1])
					list[projectName] = Project{
						projectType: PackageProject,
						project:     pathParts[1],
						program:     pathParts[2],
						pathToMain:  path,
					}
				}
			}
			return nil
		})
	}
	return list, err
}
