package monorepo

import (
	"path/filepath"
	"strings"
)

// ClassName - Return the class name based on the monorepo manifest path
func (m *Manifest) ClassName(root string) string {
	parts := strings.Split(strings.TrimPrefix(m.FileName, root), string(filepath.Separator))
	if parts[1] == "" {
		panic("empty className should not be possible")
	}
	return parts[1]
}
