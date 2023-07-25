package projectmanifest

/*
 * projects/repotool/manifest/ManifestPlatforms.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the ManifestPlatforms struct and maps
 * each appropriate element to its YAML file field.
 */

// ManifestPlatforms - Define the in-memory state of a YAML manifest
type ManifestPlatforms struct {
	Opsys   string `yaml:"opsys"`
	CpuArch string `yaml:"arch"`
}
