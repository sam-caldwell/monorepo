package projectmanifest

import (
	"gopkg.in/yaml.v3"
	"os"
)

func (manifest *Manifest) LoadFile(fileName string) (err error) {
	var raw []byte
	if raw, err = os.ReadFile(fileName); err != nil {
		manifest.err = err
		return err
	}
	return yaml.Unmarshal(raw, manifest)
}
