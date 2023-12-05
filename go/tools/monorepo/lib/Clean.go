package monorepo

import "github.com/sam-caldwell/monorepo/go/ansi"

func Clean(projectCriteria *string) error {
	projectList, err := GetProjectList(projectCriteria)
	for _, project := range projectList {
		ansi.Yellow().Println(project)
	}
	return err
}
