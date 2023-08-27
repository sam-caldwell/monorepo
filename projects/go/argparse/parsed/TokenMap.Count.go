package parsed

// Count - Return the number of items in the Namespace
func (ns *Namespace) Count() int {
	return len(ns.data)
}
