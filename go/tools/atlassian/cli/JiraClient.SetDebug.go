package cli

// SetDebug - Store the debug flag
func (client *JiraClient[T]) SetDebug(flag bool) {
    client.debug = flag
}
