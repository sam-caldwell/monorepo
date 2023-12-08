package monorepo

import (
	"path/filepath"
	"strings"
)

// ClassName - Return the class name based on the monorepo manifest path
func (m *Manifest) ClassName(root string) string {
	parts := strings.Split(strings.TrimPrefix(m.FileName, root), string(filepath.Separator))
	return parts[1]
}
