package projectmanifest

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"strings"
)

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
		OpsysName: opsys,
		CpuArch:   arch,
	})

	return manifest

}
