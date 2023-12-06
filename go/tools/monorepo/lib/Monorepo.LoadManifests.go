package monorepo

import (
	"os"
	"path/filepath"
	"strings"
)

func (m *Monorepo) LoadManifests() error {

	m.manifestList = make(map[string]Manifest)

	err := filepath.Walk(m.Root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(info.Name(), manifestYamlFile) {
			var manifest Manifest
			if err := manifest.Load(path); err != nil {
				return err
			}
			m.manifestList[path] = manifest
		}
		return nil
	})
	if err != nil {
		return err
	}
	return err
}

func (m *Monorepo) ManifestCount() int {
	return len(m.manifestList)
}
