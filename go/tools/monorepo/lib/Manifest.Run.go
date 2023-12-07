package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"path/filepath"
	"runtime"
)

// Run - Evaluate a monorepo command and execute the appropriate manifest definition.
func (m *Manifest) Run(command string, root *string, debug bool) (err error) {
	manifestDir := filepath.Dir(m.FileName)
	className := m.ClassName()
	projectName := m.ProjectName()
	opsys := runtime.GOOS
	arch := runtime.GOARCH
	switch command {
	case "build":
		if m.config.Build.Enabled {
			return m.config.Build.Execute(root, &manifestDir, &className, &projectName, &opsys, &arch, debug)
		} else {
			if debug {
				ansi.Magenta().Printf("  Disabled: %s::%s", className, projectName).LF().Reset()
			}
		}
	case "clean":
		if m.config.Clean.Enabled {
			return m.config.Clean.Execute(root, &manifestDir, &className, &projectName, &opsys, &arch, debug)
		} else {
			if debug {
				ansi.Magenta().Printf("  Disabled: %s::%s", className, projectName).LF().Reset()
			}
		}
	case "test":
		if m.config.Test.Enabled {
			if debug {
				ansi.Magenta().Printf("  Enabled: %s::%s", className, projectName).LF().Reset()
			}
			return m.config.Test.Execute(root, &manifestDir, &className, &projectName, &opsys, &arch, debug)
		} else {
			if debug {
				ansi.Magenta().Printf("  Disabled: %s::%s", className, projectName).LF().Reset()
			}
		}
	default:
		return fmt.Errorf("unsupported command %s", command)

	}
	return nil
}
