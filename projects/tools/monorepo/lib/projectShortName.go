package monorepo

import "fmt"

// projectShortName - Create class/project short name
func (m *Monorepo) projectShortName(longName Manifest) string {
	return fmt.Sprintf("%v/%v", longName.ClassName(m.Root), longName.ProjectName(m.Root))
}
