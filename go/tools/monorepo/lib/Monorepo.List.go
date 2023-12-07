package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"path/filepath"
	"strings"
)

// List - List all projects by project class
func (m *Monorepo) List() (err error) {

	objects := make(map[string][]string)

	for path, _ := range m.manifestList {
		path = filepath.Dir(strings.TrimPrefix(path, filepath.Join(m.Root, "monorepo")))
		var parts []string
		for _, part := range strings.Split(path, string(filepath.Separator)) {
			parts = append(
				parts,
				strings.TrimSpace(
					strings.TrimPrefix(
						strings.TrimSuffix(
							strings.TrimSpace(part), words.NewLine),
						words.NewLine)))
		}
		class := parts[1]
		var project string
		project = strings.Join(parts[2:], "/")
		objects[class] = append(objects[class], project)
	}
	if len(objects) == 0 {
		ansi.Red().Println("Empty list").Reset()
	}
	for class, projectList := range objects {
		ansi.Yellow().Printf("\n - class: '%s'\n", class)
		if len(projectList) == 0 {
			ansi.Yellow().Printf("    Empty Project List")
		}
		for _, project := range projectList {
			ansi.Yellow().Printf("    - %s\n", project)
		}
	}
	ansi.LF().Reset()
	return err
}
