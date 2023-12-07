package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
)

func (m *Manifest) Run(command string, debug bool) (err error) {
	className := m.ClassName()
	projectName := m.ProjectName()
	switch command {
	case "build":
		if m.config.Build.Enabled {
			return m.config.Build.Execute(className, projectName, debug)
		} else {
			if debug {
				ansi.Magenta().Printf("  Disabled: %s::%s", className, projectName).LF().Reset()
			}
		}
	case "clean":
		if m.config.Clean.Enabled {
			return m.config.Clean.Execute(className, projectName, debug)
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
			return m.config.Test.Execute(className, projectName, debug)
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
