package descriptormap

import (
	"fmt"
	"sort"
)

// ListPositionalArgs - return a formatted string list of positional arguments (name, type, help string)
func (m *Map) ListPositionalArgs(format string) (result []any) {
	var list []string
	for name, argument := range m.data {
		if (argument.GetShort() == "") || (argument.GetLong() == "") {
			typ := argument.GetType()
			list = append(list, fmt.Sprintf(format, name, typ.String(), argument.GetHelp()))
		}
	}
	sort.Strings(list)
	return misc.StringListToAnyList(&list)
}
