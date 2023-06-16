package parsed

// Get - Return a data value associated with descriptor n.
func (ns *Namespace) Get(n string) any {
	thisArg := ns.Lookup(&n)
	if thisArg == nil {
		return nil
	}
	return thisArg.GetValue()
}
