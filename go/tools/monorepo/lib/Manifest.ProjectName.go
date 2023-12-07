package monorepo

import (
	"path/filepath"
)

// ProjectName - Return the project name based on the monorepo path
func (m *Manifest) ProjectName() string {
	projectName := filepath.Base(filepath.Dir(m.FileName))
	return projectName
}
