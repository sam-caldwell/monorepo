package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
)

func (m *Manifest) Run(command string, debug bool) (err error) {
	projectName := m.ProjectName()
	switch command {
	case "build":
		if m.config.Build.Enabled {
			return m.config.Build.Execute(projectName, debug)
		} else {
			if debug {
				ansi.Magenta().Printf("  Disabled: %s", projectName).LF()
			}
		}
	case "clean":
		if m.config.Clean.Enabled {
			return m.config.Clean.Execute(projectName, debug)
		} else {
			if debug {
				ansi.Magenta().Printf("  Disabled: %s", projectName).LF()
			}
		}
	case "test":
		if m.config.Test.Enabled {
			return m.config.Test.Execute(projectName, debug)
		} else {
			if debug {
				ansi.Magenta().Printf("  Disabled: %s", projectName).LF()
			}
		}
	default:
		return fmt.Errorf("unsupported command %s", command)

	}
	return nil
}
