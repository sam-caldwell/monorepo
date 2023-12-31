package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"sort"
)

// SortByClassAndProject - Sort all manifests by class and project names.
func (m *Monorepo) SortByClassAndProject(debug bool) {
	if len(m.manifestList) < 2 {
		return
	}
	ansi.Blue().Println("Sorting manifests alphabetically").Reset()
	sort.Slice(m.manifestList, func(i, j int) bool {
		a := m.manifestList[i]
		b := m.manifestList[j]
		lhs := fmt.Sprintf("%v:%v", a.ClassName(m.Root), a.ProjectName(m.Root))
		rhs := fmt.Sprintf("%v:%v", b.ClassName(m.Root), b.ProjectName(m.Root))
		return lhs < rhs
	})
}
