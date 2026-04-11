package monorepo

import (
	"path/filepath"
	"strings"
)

// ProjectName - Return the project name based on the monorepo path
func (m *Manifest) ProjectName(root string) string {
	parts := strings.Split(strings.TrimPrefix(m.FileName, root), string(filepath.Separator))
	projectName := strings.Join(parts[2:len(parts)-1], "/")
	if projectName == "" {
		panic("empty projectName should not be possible")
	}
	return projectName
}
