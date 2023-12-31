package monorepo

import (
	"fmt"
	"sort"
)

func (m *Monorepo) SortByClass(debug bool) {
	if len(m.manifestList) < 2 {
		return
	}
	sort.Slice(m.manifestList, func(i, j int) bool {
		a := m.manifestList[i]
		b := m.manifestList[j]
		lhs := fmt.Sprintf("%v", a.ClassName(m.Root))
		rhs := fmt.Sprintf("%v", b.ClassName(m.Root))
		return lhs < rhs
	})
}
