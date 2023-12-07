package monorepo

import (
	"os"
	"path/filepath"
	"strings"
)

// ListManifests - Identify the manifest file paths.
func (m *Monorepo) ListManifests() (results []string, err error) {
	err = filepath.Walk(m.Root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(info.Name(), manifestYamlFile) {
			results = append(results, path)
		}
		return nil
	})
	if err != nil {
		return results, err
	}
	return results, err
}
