package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"path/filepath"
	"strings"
)

func (m *Monorepo) List() (err error) {

	objects := make(map[string][]string)

	for _, path := range m.manifestList {
		path = filepath.Dir(strings.TrimPrefix(path, filepath.Join(m.Root, "monorepo")))
		parts := strings.Split(path, string(filepath.Separator))
		class := parts[1]
		project := strings.Join(parts[2:], "/")
		objects[class] = append(objects[class], project)
	}
	if len(objects) == 0 {
		ansi.Red().Println("Empty list").Reset()
	}
	for class, projectList := range objects {
		ansi.Yellow().Printf("\n class: '%s'\n", class)
		if len(projectList) == 0 {
			ansi.Yellow().Printf("    Empty Project List")
		}
		for _, project := range projectList {
			ansi.Yellow().Printf("    %s\n", project)
		}
	}
	ansi.LF().Reset()
	return err
}
