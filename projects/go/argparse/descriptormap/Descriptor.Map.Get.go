package descriptormap

import (
	"github.com/sam-caldwell/go/v2/projects/go/argparse/descriptormap/descriptor"
)

// Get - Given an argument name, return its descriptor
func (m *Map) Get(name *string) *descriptor.Descriptor {

	if value, found := m.data[*name]; found {

		return &value

	}

	return nil

}
