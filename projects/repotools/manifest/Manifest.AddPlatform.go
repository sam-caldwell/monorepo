package projectmanifest

/*
 * projects/repotool/manifest/AddPlatform.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file contains defines AddPlatform() which will set an operating system-
 * cpu architecture pair as a supported platform for the project defined in the
 * given manifest.
 */

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"strings"
)

// AddPlatform - Add an opsys/cpu architecture pair as a supported platform
func (manifest *Manifest) AddPlatform(operatingSystem string, cpuArchitecture string) *Manifest {
	opsys := strings.TrimSpace(strings.ToLower(operatingSystem))
	switch opsys {
	case "darwin", "linux", "windows", "freebsd":
		break
	default:
		manifest.err = fmt.Errorf(errors.UnsupportedOpsys)
		return manifest
	}
	arch := strings.TrimSpace(strings.ToLower(cpuArchitecture))
	switch arch {
	case "amd64", "arm64":
		break
	default:
		manifest.err = fmt.Errorf(errors.UnsupportedCpuArchitecture)
		return manifest
	}

	manifest.SupportedPlatforms = append(manifest.SupportedPlatforms, ManifestPlatforms{
		Opsys:   opsys,
		CpuArch: arch,
	})

	return manifest

}
