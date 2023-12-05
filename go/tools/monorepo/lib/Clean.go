package monorepo

import "github.com/sam-caldwell/monorepo/go/ansi"

func Clean(class, projects *string) (err error) {
	projectList, err := GetProjectList(class, projects)
	if err != nil {
		return err
	}
	ansi.Yellow().
		Println(*class).
		Printf("%v\n", projectList).
		Reset()
	return err
}
