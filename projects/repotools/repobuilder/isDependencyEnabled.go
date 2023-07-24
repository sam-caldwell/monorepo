package repobuilder

import (
	"fmt"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
	"io/fs"
	"path/filepath"
)

func isDependencyEnabled(rootDirectory *string, manifest *projectmanifest.Manifest) (err error) {

	const projectType = "projects"

	for _, dependency := range manifest.Dependencies {

		searchPath := filepath.Join(*rootDirectory, projectType, dependency)

		err = filepath.Walk(searchPath, func(path string, info fs.FileInfo, err error) error {

			if err != nil {
				return err
			}

			var depManifest projectmanifest.Manifest

			if filepath.Base(path) != "MANIFEST.yaml" {
				return nil
			}

			err = depManifest.LoadFile(path)

			if err != nil {
				return fmt.Errorf("dependency manifest (%s) Load failed: %v", depManifest.Name, err)
			}

			if !depManifest.IsBuildEnabled() {
				return fmt.Errorf("dependency build not enabled.  Project '%s' dependency '%s'",
					manifest.Name,
					depManifest.Name)

			}

			return nil

		})

		if err != nil {
			return err
		}

	}

	return err

}
