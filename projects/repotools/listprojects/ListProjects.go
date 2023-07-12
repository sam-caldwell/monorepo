package listprojects

import (
	"github.com/sam-caldwell/go/v2/projects/repotools/filters"
)

// ListProjects - List the enabled or disabled projects in the repo
func ListProjects(filter filters.Filter) (list map[string]any, err error) {
	//
	//rootDirectory, err := repotools.findRepoRoot()
	//if err != nil {
	//	return list, err
	//}
	//
	//// Verify the path exists.
	//if !directory.Existsp(&rootDirectory) {
	//	return list, fmt.Errorf(fs.ErrDoesNotExist, fs.Directory, rootDirectory)
	//}
	//list = make(map[string]repotools.Project)
	//
	//var projectTypes []string
	//if filter.HasFilter(filters.ShowCommands) || filter.HasFilter(filters.ShowAny) {
	//	projectTypes = append(projectTypes, repotools.Command)
	//}
	//if filter.HasFilter(filters.ShowProjects) || filter.HasFilter(filters.ShowAny) {
	//	projectTypes = append(projectTypes, repotools.Package)
	//}
	//
	//for _, projectType := range projectTypes {
	//	ansi.Green().Printf("Walking %s\n", projectType).Reset()
	//	// Walk the directory structure and find all the projects.
	//	err = filepath.Walk(projectType, func(path string, info os.FileInfo, err error) error {
	//		if err != nil {
	//			return err
	//		}
	//		pathParts := directory.SplitToList(path)
	//		manifestFile := filepath.Base(path)
	//		if (pathParts[0] != projectType) || (manifestFile != repotools.ManifestFile) || (len(pathParts) < 2) {
	//			return nil // Filter off the irrelevant
	//		}
	//
	//		manifest, err := repoLinter.LoadProjectManifest(path)
	//		if err != nil {
	//			return err
	//		}
	//
	//		if filter.ApplyFilter(projectType, manifest) {
	//			ansi.Yellow().Printf("filtered: '%s'", path).LF().Reset()
	//			return nil //skip filtered project
	//		}
	//		ansi.Green().Printf("allowed:  '%s'", path).LF().Reset()
	//
	//		thisPath := filepath.Base(path)
	//		switch projectType {
	//		case repotools.Command:
	//			if !file.Exists(filepath.Join(thisPath, "main.go")) {
	//				return nil
	//			}
	//			list[fmt.Sprintf("%s::%s/%s", repotools.Command, pathParts[1], pathParts[2])] = repotools.Project{
	//				ProjectType: repotools.CommandProject,
	//				Project:     pathParts[1],
	//				Program:     pathParts[2],
	//				Path:        thisPath,
	//			}
	//		case repotools.Package:
	//			projectName := fmt.Sprintf("%s::%s", repotools.Package, pathParts[1])
	//			list[projectName] = repotools.Project{
	//				ProjectType: repotools.PackageProject,
	//				Project:     pathParts[1],
	//				Program:     pathParts[2],
	//				Path:        thisPath,
	//			}
	//		}
	//		return nil
	//	})
	//}
	return list, err
}
