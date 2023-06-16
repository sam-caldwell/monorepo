package descriptormap

import "github.com/sam-caldwell/go-monorepo/v2/projects/argparse/descriptormap/descriptor"

// GetByShort - perform a linear search of the descriptor map for a matching short arg
func (m *Map) GetByShort(token *string) (*string, *descriptor.Descriptor) {
	for _, thisDesc := range m.data {
		if thisDesc.GetShort() == *token {
			return token, &thisDesc
		}
	}
	return nil, nil
}
