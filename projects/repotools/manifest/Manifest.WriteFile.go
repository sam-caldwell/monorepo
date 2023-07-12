package projectmanifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

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
