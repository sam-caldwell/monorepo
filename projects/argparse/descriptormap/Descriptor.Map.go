package descriptormap

import "github.com/sam-caldwell/argparse/v2/argparse/descriptormap/descriptor"

// Map - map of descriptors (name:descriptor)
type Map struct {
	data map[string]descriptor.Descriptor
}
