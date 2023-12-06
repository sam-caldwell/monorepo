package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
)

func (m *Manifest) Run(command string, debug bool) (err error) {
	switch command {
	case "build":
		if m.config.Build.Enabled {
			return m.config.Build.Execute(debug)
		} else {
			if debug {
				ansi.Magenta().Printf("  Disabled: %s", m.ProjectName()).LF()
			}
		}
	case "clean":
		if m.config.Clean.Enabled {
			return m.config.Clean.Execute(debug)
		} else {
			if debug {
				ansi.Magenta().Printf("  Disabled: %s", m.ProjectName()).LF()
			}
		}
	case "test":
		if m.config.Test.Enabled {
			return m.config.Test.Execute(debug)
		} else {
			if debug {
				ansi.Magenta().Printf("  Disabled: %s", m.ProjectName()).LF()
			}
		}
	default:
		return fmt.Errorf("unsupported command %s", command)

	}
	return nil
}
