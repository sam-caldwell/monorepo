package descriptormap

import "github.com/sam-caldwell/go/v2/projects/argparse/descriptormap/descriptor"

// List - return the descriptor map
func (m *Map) List() map[string]descriptor.Descriptor {
	if m.data == nil {
		return make(map[string]descriptor.Descriptor)
	}
	return m.data
}
