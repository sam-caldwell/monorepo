package repotools

func GetProjects(directory string, enabled bool) ([]string, error) {
	if enabled {
		return []string{"test", "me", "now"}, nil
	} else {
		return []string{"test", "disabled", "projects"}, nil
	}
}

func Setup(noop bool) error {
	return nil
}
