package descriptormap

import "github.com/sam-caldwell/go/v2/projects/argparse/argparse/descriptormap/descriptor"

// GetByLong - perform a linear search of the descriptor map for a matching Long arg
func (m *Map) GetByLong(token *string) (*string, *descriptor.Descriptor) {
	for _, thisDesc := range m.data {
		if thisDesc.GetLong() == *token {
			return token, &thisDesc
		}
	}
	return nil, nil
}
