package projectmanifest

/*
 * projects/repotool/manifest/Manifest.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the Manifest struct and maps
 * each appropriate element to its YAML file field.
 */

// Manifest - Define the in-memory state of a YAML manifest
type Manifest struct {
	fileName           string
	Name               string              `yaml:"name"`
	Author             string              `yaml:"author"`
	Options            ManifestOptions     `yaml:"options"`
	SupportedPlatforms []ManifestPlatforms `yaml:"supportedPlatforms"`
	Dependencies       []string            `yaml:"dependencies"`
	err                error
}
