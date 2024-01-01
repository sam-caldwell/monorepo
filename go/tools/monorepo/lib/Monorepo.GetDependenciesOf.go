package monorepo

// GetDependenciesOf - given a project identified by classname and projectname, return its list of dependencies.
func (m *Monorepo) GetDependenciesOf(name string) (result []string) {
	for _, manifest := range m.manifestList {
		if m.projectShortName(manifest) == name {
			return manifest.config.Header.Dependencies
		}
	}
	return []string{}
}
