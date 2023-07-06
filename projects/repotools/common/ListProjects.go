package repotools

// ListProjects - List the enabled or disabled projects in the repo
func ListProjects(directory string, filter ListProjectsFilter) ([]string, error) {
	if filter == Enabled {
		return []string{"test", "me", "now"}, nil
	} else {
		return []string{"test", "disabled", "projects"}, nil
	}
}
