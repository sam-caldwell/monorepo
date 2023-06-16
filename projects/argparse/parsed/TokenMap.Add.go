package parsed

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/parsed/token"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
)

// Add - Add a new token to the Namespace
func (ns *Namespace) Add(name *string, typ types.ArgTypes, value *any) error {
	if ns.data == nil {
		ns.data = make(map[string]token.Token)
	}
	if thisArg, ok := ns.data[*name]; ok {
		if thisArg.GetType() != typ {
			return fmt.Errorf(errTypeCannotChange, *name)
		}
		if err := thisArg.SetValue(*value); err != nil {
			return err
		}
		ns.data[*name] = thisArg
	} else {
		// Create a new Token
		var element token.Token

		if err := element.Set(typ, *value); err != nil {
			return err
		}
		// Store the new record in the map
		ns.data[*name] = element
	}
	return nil
}
