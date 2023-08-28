package descriptormap

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/sets/simple"
)

// FindDuplicates - find any duplicate arguments and throw and error
func (m *Map) FindDuplicates() error {
	var argSet simple.Set
	var shortSet simple.Set
	var longSet simple.Set

	for name, argument := range m.data {
		if err := shortSet.Add(argument.GetShort()); err != nil {
			return fmt.Errorf(errArgumentCannotBeRedefined, name)
		}
		if err := longSet.Add(argument.GetLong()); err != nil {
			return fmt.Errorf(errArgumentCannotBeRedefined, name)
		}
		if err := argSet.Add(name); err != nil {
			return fmt.Errorf(errArgumentCannotBeRedefined, name)
		}
	}
	return nil
}
