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
				sortedList = append(sortedList, subDepName)
			}
			if (depName == projectName) || list.Contains(sortedList, depName) {
				continue
			}
			sortedList = append(sortedList, depName)
		}
		if list.Contains(sortedList, projectName) {
			continue
		}
		sortedList = append(sortedList, projectName)
	}
	if debug {
		for pos, row := range sortedList {
			ansi.Blue().Printf(" %d : %s", pos, row).LF().Reset()
		}
	}
	var manifestList []Manifest
	for _, project := range sortedList {
		for _, original := range m.manifestList {
			if m.projectShortName(original) == project {
				manifestList = append(manifestList, original)
			}
		}
	}
	m.manifestList = manifestList
}
