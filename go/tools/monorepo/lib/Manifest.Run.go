package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/list"
	"path/filepath"
	"runtime"
)

// Run - Evaluate a monorepo command and execute the appropriate manifest definition.
func (m *Manifest) Run(command string, root *string, debug bool) (err error) {
	// The config stage will point to our actual actions to be performed
	var configStage Stage
	switch command {
	case "build":
		configStage = m.config.Build
	case "clean":
		configStage = m.config.Clean
	case "test":
		configStage = m.config.Test
	default:
		return fmt.Errorf("unsupported command %s", command)
	}

	manifestDir := filepath.Dir(m.FileName)
	className := m.ClassName(*root)
	projectName := m.ProjectName(*root)

	// If opsys contains, any we will build for our current opsys only
	if len(m.config.Header.Opsys) == 0 || list.Contains(m.config.Header.Opsys, "any") {
		m.config.Header.Opsys = []string{runtime.GOOS}
	}
	// If arch contains, any we will build for our current architecture only
	if len(m.config.Header.Arch) == 0 || list.Contains(m.config.Header.Arch, "any") {
		m.config.Header.Arch = []string{runtime.GOARCH}
	}

	//ansi.
	//	Magenta().
	//	Printf("Cleaning %s:%-40s (%s/%s) %s",
	//		className,
	//		projectName,
	//		m.config.Header.Opsys,
	//		m.config.Header.Arch,
	//		manifestDir).
	//	LF().
	//	Reset()

	for _, opsys := range m.config.Header.Opsys {
		for _, arch := range m.config.Header.Arch {
			ansi.
				Dim().
				Magenta().
				Printf("%s:%s/%s (%s/%s)", command, className, projectName, opsys, arch).
				Bold().
				LF()
			if configStage.Enabled {
				ansi.
					Cyan().
					Printf("   └─").
					Green().
					Bold().
					Printf("Enabled:  ").
					Dim().
					White().
					Printf("%s/%s", className, projectName).
					Bold().
					LF().
					Reset()
				err = configStage.Execute(root, &manifestDir, &className, &projectName, &opsys, &arch, debug)
			} else {
				ansi.
					Cyan().
					Printf("   └─").
					White().Dim().
					Printf("Disabled: ").
					Magenta().
					Printf("%s/%s", className, projectName).
					Bold().
					LF().
					Reset()
			}
		}
	}
	return err
}
