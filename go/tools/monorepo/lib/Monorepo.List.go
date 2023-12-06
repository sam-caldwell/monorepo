package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"os"
	"path/filepath"
	"strings"
)

func (m *Monorepo) List() (err error) {

	objects := make(map[string][]string)

	err = filepath.Walk(m.Root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(info.Name(), manifestYamlFile) {
			path := filepath.Dir(strings.TrimPrefix(path, m.Root))
			parts := strings.Split(path, string(filepath.Separator))
			parts = parts[1:]
			class := parts[1]
			if len(parts) < 3 {
				return nil
			}
			project := strings.Join(parts[2:], "/")
			objects[class] = append(objects[class], project)
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(objects) == 0 {
		ansi.Red().Println("Empty list").Reset()
	}
	for class, projectList := range objects {
		ansi.Yellow().Printf("\n class: '%s'\n", class)
		for _, project := range projectList {
			ansi.Yellow().Printf("    %s\n", project)
		}
	}
	ansi.LF().Reset()
	return err
}
