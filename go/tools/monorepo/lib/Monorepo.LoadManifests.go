package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"os"
	"path/filepath"
	"strings"
)

func (m *Monorepo) LoadManifests() {
	ansi.Cyan().
		Print("Discovering project manifests").LF().Reset()

	err := filepath.Walk(m.Root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(info.Name(), manifestYamlFile) {
			var manifest Manifest
			manifest.FileName = path
			if err := manifest.Load(&path); err != nil {
				return err
			}
			m.manifestList = append(m.manifestList, manifest)
		}
		return nil
	})
	if err != nil {
		ansi.Red().
			Printf("Error discovering manifests\n%s", err).
			Reset().Fatal(exit.GeneralError)
	}
	ansi.Green().
		Printf("Discovered %d manifests (loaded)\n", len(m.manifestList)).
		LF().Reset()
}

func (m *Monorepo) ManifestCount() int {
	return len(m.manifestList)
}
