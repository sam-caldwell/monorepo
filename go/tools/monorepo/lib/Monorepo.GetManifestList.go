package monorepo

import (
	"os"
	"path/filepath"
	"strings"
)

func (m *Monorepo) GetManifestList() error {
	err := filepath.Walk(m.Root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(info.Name(), manifestYamlFile) {
			m.manifestList = append(m.manifestList, path)
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
