package projectmanifest

/*
 * projects/repotool/manifest/WriteFile.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines WriteFile() which will write the internal state
 * of a Manifest to the filename identified by the internal filename
 * property.
 */

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// WriteFile - Write the manifest file
func (manifest *Manifest) WriteFile() *Manifest {
	if manifest.err != nil {
		return manifest
	}
	yamlContent, err := yaml.Marshal(*manifest)
	if err != nil {
		manifest.err = fmt.Errorf("failed to marshal YAML: %v", err)
	}
	file, err := os.Create(manifest.fileName)
	if err != nil {
		manifest.err = fmt.Errorf("failed to create file: %v", err)
	}
	defer func() { _ = file.Close() }()

	_, err = file.Write(yamlContent)
	if err != nil {
		manifest.err = fmt.Errorf("failed to write YAML content to file: %v", err)
	}
	return manifest
}
