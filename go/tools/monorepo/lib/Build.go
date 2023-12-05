package monorepo

import (
	"bufio"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"os"
	"strings"
)

func Build(projectCriteria *string) (err error) {
	projectList, err := GetProjectList(projectCriteria)
	for _, projectManifest := range projectList {
		ansi.Blue().Printf("Project Manifest: %s", projectManifest).LF().Reset()

		for _, project := range LoadManifest(&projectManifest) {
			ansi.Blue().Printf("project: %s", project).LF().Reset()
		}
	}
	return err
}

func LoadManifest(manifest *string) (projects map[string]string, err error) {
	projects = make(map[string]string)
	f, err := os.Open(*manifest)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "#") {
			continue
		}
		project, language := func() (p string, l string) {
			//
			//project:language #comment string is ignored
			//
			parts := strings.Split(text, "#")
			p = parts[0]
			parts = strings.Split(p, ":")
			p = parts[0]
			l = parts[1]
			return p, l
		}()

	}
}
