package argparse

import (
	"github.com/sam-caldwell/go/v2/projects/argparse/descriptormap/descriptor"
)

// Get - append the given Argument object to the list of Arguments.
func (arg *Arguments) Get(name string) *descriptor.Descriptor {
	return arg.descriptors.Get(&name)
}
