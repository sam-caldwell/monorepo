package repotools

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/fs"
	"github.com/sam-caldwell/go/v2/projects/fs/directory"
	"github.com/sam-caldwell/go/v2/projects/repotools/common/projectFilter"
	"os"
	"path/filepath"
)

type projectType byte
const(
	CommandProject projectType = iota
	PackageProject
)

const (
	Command = "cmd"
	Package = "projects"
)

type Project struct {
	projectType projectType
	project    string
	program    string
	pathToMain string
}


// ListProjects - List the enabled or disabled projects in the repo
func ListProjects(filter projectFilter.Filter) (list []Project, err error) {

	rootDirectory, err := findRepoRoot()
	if err != nil {
		return list, err
	}

	// Verify the path exists.
	if !directory.Existsp(&rootDirectory) {
		return list, fmt.Errorf(fs.ErrDoesNotExist, fs.Directory, rootDirectory)
	}


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
						list = append(list, Project{
							project:    pathParts[1],
							program:    pathParts[2],
							pathToMain: path,
						})
					}
					return nil
				}
			case Package:
				if filter.HasFilter(projectFilter.Package) {
					if projectFilter.FilterThisPath(filepath.Dir(path)) {
						//We are filtering this path...nothing else to see here.  Move along people.
						return nil
					}
					list = append(list, Project{
						project:    pathParts[1],
						program:    pathParts[2],
						pathToMain: path,
					})
				}
				return nil
			}
		}
	})
}