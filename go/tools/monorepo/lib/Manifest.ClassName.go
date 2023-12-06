package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"path/filepath"
	"strings"
)

// ClassName - Return the class name based on the monorepo manifest path
func (m *Manifest) ClassName() string {
	path := filepath.Dir(strings.TrimPrefix(m.FileName, directory.GetCurrent()))
	parts := strings.Split(path, string(filepath.Separator))
	parts = parts[1:]
	return parts[1]
}
