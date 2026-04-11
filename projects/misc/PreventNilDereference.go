package misc

func PreventNilDereference[T any](n *T) T {
	if n == nil {
		var defaultValue T
		return defaultValue
	}
	return *n
}
