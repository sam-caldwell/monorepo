package parsed

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
)

// GetType - Lookup and return the token's type
func (ns *Namespace) GetType(n *string) (t types.ArgTypes, e error) {
	if o := ns.Lookup(n); o != nil {
		return o.GetType(), nil
	}
	return t, fmt.Errorf(errNotFound) //Not found.
}
