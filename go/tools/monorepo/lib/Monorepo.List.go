package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"strings"
)

// List - List all projects by project class
func (m *Monorepo) List() (err error) {

	ansi.Yellow().LF().
		Printf(" ┌─────────────┐").LF().
		Printf(" │monorepo root│").LF().
		Printf(" └─────╥───────┘").LF()
	if len(m.manifestList) == 0 {
		ansi.Red().Println("Empty list").Reset().LF()
	}
	prevClass := ""
	for _, item := range m.manifestList {
		class := item.ClassName(m.Root)
		if strings.TrimSpace(class) == "" {
			continue
		}
		project := item.ProjectName(m.Root)
		if strings.TrimSpace(project) == "" {
			continue
		}

		if class != prevClass {
			ansi.Yellow().
				Printf("       ║").LF().
				Printf("       ╠════╦═class: '%s'", class).LF().
				Printf("       ║    ║").LF()
		}
		ansi.Yellow().
			Printf("       ║    ╠══%s", project).LF()
		prevClass = class
	}
	ansi.LF().Reset()
	return err
}
