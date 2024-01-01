package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/list"
)

// SortManifestsByDependencies - Sort manifests by dependency order.
func (m *Monorepo) SortManifestsByDependencies(debug bool) {

	// create a simplified string list of projects
	var sortedList []string
	for _, manifest := range m.manifestList {
		projectName := m.projectShortName(manifest)
		for _, depName := range manifest.config.Header.Dependencies {
			for _, subDepName := range m.GetDependenciesOf(depName) {
				if (subDepName == projectName) || (depName == projectName) || list.Contains(sortedList, subDepName) {
					continue
				}
				ansi.Yellow().Printf("Adding project %s dep %s subDep %s", projectName, depName, subDepName).
					LF().Reset()
				sortedList = append(sortedList, subDepName)
			}
			if (depName == projectName) || list.Contains(sortedList, depName) {
				continue
			}
			ansi.Yellow().Printf("Adding project %s dep %s", projectName, depName).LF().Reset()
			sortedList = append(sortedList, depName)
		}
		if list.Contains(sortedList, projectName) {
			continue
		}
		ansi.Yellow().Printf("Adding project %s", projectName).LF().Reset()
		sortedList = append(sortedList, projectName)
	}
	if debug {
		for pos, row := range sortedList {
			ansi.Blue().Printf(" %d : %s", pos, row).LF().Reset()
		}
	}
}

//
//func itemInList(item string, list []string) bool {
//	for _, i := range list {
//		if item == i {
//			return true
//		}
//	}
//	return false
//}
