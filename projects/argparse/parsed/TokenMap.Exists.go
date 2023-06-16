package parsed

// Exists - return object if found or nil of not found
func (ns *Namespace) Exists(n *string) bool {
	return ns.Lookup(n) != nil
}
