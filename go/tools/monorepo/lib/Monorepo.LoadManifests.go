package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"os"
	"path/filepath"
	"strings"
)

func (m *Monorepo) LoadManifests() {
	m.manifestList = make(map[string]Manifest)

	ansi.Cyan().Println("Discovering project manifests")

	count := 0
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
			count++
		}
		return nil
	})
	if err != nil {
		ansi.Red().
			Printf("Error discovering manifests\n%s", err).
			Reset().
			Fatal(exit.GeneralError)
	}
	ansi.Green().
		Printf("Discovered %d manifests (loaded)\n", count).
		LF().
		Reset()
}

func (m *Monorepo) ManifestCount() int {
	return len(m.manifestList)
}
