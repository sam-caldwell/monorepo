package descriptormap

import (
	"github.com/sam-caldwell/go/v2/projects/go/argparse/descriptormap/descriptor"
)

// Map - map of descriptors (name:descriptor)
type Map struct {
	data map[string]descriptor.Descriptor
}
