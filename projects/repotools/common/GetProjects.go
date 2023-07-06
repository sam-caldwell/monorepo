package repotools

// GetProjects - List the enabled or disabled projects in the repo
func GetProjects(directory string, enabled bool) ([]string, error) {
	if enabled {
		return []string{"test", "me", "now"}, nil
	} else {
		return []string{"test", "disabled", "projects"}, nil
	}
}
