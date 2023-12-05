package monorepo

import "github.com/sam-caldwell/monorepo/go/ansi"

func Test(class, projects *string) (err error) {
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
