package parsed

import (
	"github.com/sam-caldwell/argparse/v2/argparse/parsed/token"
)

// Lookup - return object if found or nil of not found
func (ns *Namespace) Lookup(n *string) *token.Token {
	if (ns.data != nil) && (n != nil) {
		if thisArg, found := (*ns).data[*n]; found {
			return &thisArg
		}
	}
	return nil
}
