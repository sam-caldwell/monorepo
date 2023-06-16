package descriptormap

import "github.com/sam-caldwell/argparse/v2/argparse/descriptormap/descriptor"

// GetByName - search the descriptor map for a token with the given *token name
func (m *Map) GetByName(token *string) (*string, *descriptor.Descriptor) {
	return token, m.Get(token)
}
