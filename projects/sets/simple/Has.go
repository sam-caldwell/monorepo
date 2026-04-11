package simple

// Has - return true if row exists
func (set *Set[T]) Has(item T) bool {
	return (set.data != nil) &&
		func() bool {
			_, ok := set.data[item]
			return ok
		}()
}
