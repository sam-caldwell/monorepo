package repotools

import (
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func ListDependencies(directory string) ([]string, error) {
	projects := make(map[string][]string)

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() != directory {
			manifestPath := filepath.Join(path, "MANIFEST.yaml")
			manifestData, err := os.ReadFile(manifestPath)
			if err != nil {
				return err
			}

			var manifest projectmanifest.Manifest
			err = yaml.Unmarshal(manifestData, &manifest)
			if err != nil {
				return err
			}

			projects[manifest.Name] = manifest.Dependencies
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	sortedProjects := make([]string, 0, len(projects))
	visited := make(map[string]bool)

	var visit func(string)
	visit = func(project string) {
		if visited[project] {
			return
		}

		visited[project] = true
		for _, dependency := range projects[project] {
			visit(dependency)
		}

		sortedProjects = append(sortedProjects, project)
	}

	for project := range projects {
		visit(project)
	}

	// Reverse the sorted projects
	for i, j := 0, len(sortedProjects)-1; i < j; i, j = i+1, j-1 {
		sortedProjects[i], sortedProjects[j] = sortedProjects[j], sortedProjects[i]
	}

	return sortedProjects, nil
}
