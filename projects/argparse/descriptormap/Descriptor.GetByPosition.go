package descriptormap

import (
	"github.com/sam-caldwell/go-monorepo/v2/projects/argparse/descriptormap/descriptor"
)

// GetByPosition - given a position, return the descriptor name and object
func (m *Map) GetByPosition(position int) (*string, *descriptor.Descriptor) {
	for n, a := range m.data {
		if a.GetPosition() == position {
			name := n
			argument := a
			return &name, &argument
		}
	}
	return nil, nil
}
