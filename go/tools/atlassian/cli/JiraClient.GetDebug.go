package cli

// GetDebug - Return the debug flag
func (client *JiraClient[T]) GetDebug() (flag bool) {
    return client.debug
}
