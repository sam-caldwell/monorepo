package parsed

import "fmt"

// GetName - Return a data value associated with descriptor n.
func (ns *Namespace) GetName(n *string) (string, error) {
	if o := ns.Lookup(n); o != nil {
		return *n, nil
	}
	return "", fmt.Errorf(errNotFound) //Not found.
}
