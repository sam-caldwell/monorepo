package projectmanifest

/*
 * projects/repotool/manifest/LoadFile.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines LoadFile() which load
 * a project manifest given a filename
 */

import (
	"gopkg.in/yaml.v3"
	"os"
)

// LoadFile - load the manifest file (YAML)
func (manifest *Manifest) LoadFile(fileName string) (err error) {
	var raw []byte
	if raw, err = os.ReadFile(fileName); err != nil {
		manifest.err = err
		return err
	}
	return yaml.Unmarshal(raw, manifest)
}
