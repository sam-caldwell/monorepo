package monorepo

import (
	"os"
	"path/filepath"
	"strings"
)

func (m *Monorepo) Clean() (err error) {
	printHeader("Cleaning...")

	err = filepath.Walk(m.Root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(info.Name(), manifestYamlFile) {
			var manifest Manifest
			if err := manifest.Load(path); err != nil {
				return err
			}
			if err := manifest.Run("clean", m.Debug); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	printFooter()
	return err
}
