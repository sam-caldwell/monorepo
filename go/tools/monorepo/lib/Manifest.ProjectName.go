package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"path/filepath"
	"strings"
)

func (m *Manifest) ProjectName() string {
	path := filepath.Dir(strings.TrimPrefix(m.FileName, directory.GetCurrent()))
	parts := strings.Split(path, string(filepath.Separator))
	return strings.Join(parts[2:], "/")
}
