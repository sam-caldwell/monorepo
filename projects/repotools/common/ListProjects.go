package repotools

// ListProjects - List the enabled or disabled projects in the repo
func ListProjects(directory string, enabled bool) ([]string, error) {
	if enabled {
		return []string{"test", "me", "now"}, nil
	} else {
		return []string{"test", "disabled", "projects"}, nil
	}
}
