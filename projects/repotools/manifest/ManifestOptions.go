package projectmanifest

/*
 * projects/repotool/manifest/ManifestOptions.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the ManifestOptions struct and maps
 * each appropriate element to its YAML file field.
 */

// ManifestOptions - Define the in-memory state of a YAML manifest
type ManifestOptions struct {
	BuildEnabled bool   `yaml:"buildEnabled"`
	LintEnabled  bool   `yaml:"lintEnabled"`
	PackEnabled  bool   `yaml:"packEnabled"`
	ScanEnabled  bool   `yaml:"scanEnabled"`
	SignEnabled  bool   `yaml:"signEnabled"`
	Language     string `yaml:"language"`
}
