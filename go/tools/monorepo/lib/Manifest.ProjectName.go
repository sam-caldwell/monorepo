package monorepo

import (
	"path/filepath"
)

// ProjectName - Return the project name based on the monorepo path
func (m *Manifest) ProjectName() string {
	projectName := filepath.Base(filepath.Dir(m.FileName))

	//path := filepath.Dir(strings.TrimPrefix(m.FileName, directory.GetCurrent()))
	//parts := strings.Split(path, string(filepath.Separator))
	//projectName := strings.Join(parts[2:], "/")
	//if projectName == "" {
	//	ansi.Red().
	//		Printf("Internal error (empty project name) ").
	//		Printf("File: %s", m.FileName).
	//		LF().
	//		Reset().
	//		Fatal(exit.GeneralError)
	//}
	return projectName
}
