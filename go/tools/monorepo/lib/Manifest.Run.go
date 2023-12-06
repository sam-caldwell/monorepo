package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
)

func (m *Manifest) Run(command string) (err error) {
	switch command {
	case "build":
		if m.config.Build.Enabled {
			return m.config.Build.Execute()
		} else {
			ansi.Cyan().Printf("  Disabled: %s", m.ProjectName()).LF()
		}
	case "clean":
		if m.config.Clean.Enabled {
			return m.config.Clean.Execute()
		} else {
			ansi.Cyan().Printf("  Disabled: %s", m.ProjectName()).LF()
		}
	case "test":
		if m.config.Test.Enabled {
			return m.config.Test.Execute()
		} else {
			ansi.Cyan().Printf("  Disabled: %s", m.ProjectName()).LF()
		}
	default:
		return fmt.Errorf("unsupported command %s", command)

	}
	//ansi.Green().Printf("Success: %s", m.ProjectName()).LF().Reset()
	return nil
}
