package descriptormap

import (
	"github.com/sam-caldwell/go/v2/projects/argparse/argparse/descriptormap/descriptor"
	"github.com/sam-caldwell/go/v2/projects/argparse/argparse/types"
	"github.com/sam-caldwell/go/v2/projects/argparse/argparse/valid"
	"strings"
)

// Add - Add the new descriptor to the descriptor map
func (m *Map) Add(pos *counters.ConditionalCounter, name, short, long string, typ types.ArgTypes,
	required bool, dValue any, help string) error {

	// Make sure our map is initialized
	if m.data == nil {
		m.data = make(map[string]descriptor.Descriptor)
	}

	//Sanitize our argument name.
	if err := valid.IsNameArg(&name); err != nil {
		return err
	}
	argName := strings.ToLower(name)

	//Create a new descriptor
	var argDesc descriptor.Descriptor
	if err := argDesc.Add(pos, short, long, typ, required, dValue, help); err != nil {
		return err
	}

	//Add the descriptor to our map.
	m.data[argName] = argDesc
	return nil
}
