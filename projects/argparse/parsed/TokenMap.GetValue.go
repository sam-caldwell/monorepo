package parsed

import (
	"fmt"
)

// GetValue - Lookup and return the token's value
func (ns *Namespace) GetValue(n *string) (any, error) {
	if o := ns.Lookup(n); o != nil {
		return o.GetValue(), nil
	}
	return nil, fmt.Errorf(errNotFound) //Not found.
}
